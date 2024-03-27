package turno

import (
	"database/sql"
	"time"

	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
	"github.com/bautista00/Final_BackEnd_Go/internal/odontologo"
	"github.com/bautista00/Final_BackEnd_Go/internal/paciente"
)

type MySQLRepository struct {
	db             *sql.DB
	pacienteRepo   paciente.Repository
	odontologoRepo odontologo.Repository
}

func NewMySQLRepository(db *sql.DB, pacienteRepo paciente.Repository, odontologoRepo odontologo.Repository) *MySQLRepository {
	return &MySQLRepository{
		db:             db,
		pacienteRepo:   pacienteRepo,
		odontologoRepo: odontologoRepo,
	}
}

type Repository interface {
	GetAll() ([]domain.Turno, error)
	GetByID(id int) (domain.Turno, error)
	Create(turno domain.Turno, pacienteID int, odontologoID int) (domain.Turno, error)
	Update(id int, turno domain.Turno) (domain.Turno, error)
	Delete(id int) error
}

func (r *MySQLRepository) GetAll() ([]domain.Turno, error) {
	rows, err := r.db.Query("SELECT id, paciente_dni, odontologo, fecha_hora, descripcion FROM turnos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var turnos []domain.Turno
	for rows.Next() {
		var t domain.Turno
		err := rows.Scan(&t.ID, &t.Paciente, &t.Odontologo, &t.FechaHora, &t.Descripcion)
		if err != nil {
			return nil, err
		}
		turnos = append(turnos, t)
	}
	return turnos, nil
}

func (r *MySQLRepository) Create(pacienteID int, odontologoID int, FechaHora time.Time, Descripcion string) (domain.Turno, domain.Paciente, domain.Odontologo, error) {
	paciente, err := r.pacienteRepo.GetByID(pacienteID)
	if err != nil {
		return domain.Turno{}, domain.Paciente{}, domain.Odontologo{}, err
	}

	odontologo, err := r.odontologoRepo.GetByID(odontologoID)
	if err != nil {
		return domain.Turno{}, domain.Paciente{}, domain.Odontologo{}, err
	}

	turno := domain.Turno{
		Paciente:    paciente,
		Odontologo:  odontologo,
		FechaHora:   FechaHora,
		Descripcion: Descripcion,
	}

	result, err := r.db.Exec("INSERT INTO turnos (paciente, odontologo, fecha_hora, descripcion) VALUES (?, ?, ?, ?)",
		paciente.DNI, odontologo.ID, FechaHora, Descripcion)
	if err != nil {
		return domain.Turno{}, domain.Paciente{}, domain.Odontologo{}, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return domain.Turno{}, domain.Paciente{}, domain.Odontologo{}, err
	}

	turno.ID = int(lastInsertID)
	return turno, paciente, odontologo, nil
}

func (r *MySQLRepository) GetByID(id int) (domain.Turno, error) {
	var turno domain.Turno
	err := r.db.QueryRow("SELECT id, paciente_dni, odontologo, fecha_hora, descripcion FROM turnos WHERE id = ?", id).
		Scan(&turno.ID, &turno.Paciente, &turno.Odontologo, &turno.FechaHora, &turno.Descripcion)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (r *MySQLRepository) Update(id int, turno domain.Turno) (domain.Turno, error) {
	_, err := r.db.Exec("UPDATE turnos SET paciente_dni = ?, odontologo = ?, fecha_hora = ?, descripcion = ? WHERE id = ?",
		turno.Paciente, turno.Odontologo, turno.FechaHora, turno.Descripcion, id)
	if err != nil {
		return domain.Turno{}, err
	}
	turno.ID = id
	return turno, nil
}

func (r *MySQLRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM turnos WHERE id = ?", id)
	return err
}
