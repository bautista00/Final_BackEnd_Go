package turnos

import (
	"database/sql"

	"github.com/bautista00/Final_BackEnd_Go/internal/domain"

	"fmt"
)

type MySQLRepository struct {
	db *sql.DB
}

type TurnoDetail struct {
	Paciente    domain.Paciente   `json:"paciente"`
	Odontologo  domain.Odontologo `json:"odontologo"`
	FechaHora       string            `json:"fechaHora" example:"2006-01-02 15:04:05"`
	Descripcion string            `json:"descripcion" example:"string"`
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db}
}

type Repository interface {
	GetAll() ([]domain.Turno, error)
	GetByID(id int) (domain.Turno, error)
	GetByPacienteDNI(dni string) (TurnoDetail, error)
	Create(t domain.Turno) (domain.Turno, error)
	Update(id int, p domain.Turno) (domain.Turno, error)
	Delete(id int) error
}

func (r *MySQLRepository) GetAll() ([]domain.Turno, error) {
	rows, err := r.db.Query("SELECT * FROM turnos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var turnos []domain.Turno
	for rows.Next() {
		var t domain.Turno
		var fecha string
		err := rows.Scan(&t.Id, &t.PacienteId, &t.OdontologoId, &t.FechaHora, &t.Descripcion)
		if err != nil {
			return nil, err
		}

		t.FechaHora = fecha
		turnos = append(turnos, t)
	}
	return turnos, nil
}

func (r *MySQLRepository) GetByID(id int) (domain.Turno, error) {
	query := "SELECT * FROM turnos WHERE id = ?"
	row := r.db.QueryRow(query, id)
	var t domain.Turno

	err := row.Scan(&t.Id, &t.PacienteId, &t.OdontologoId, &t.FechaHora, &t.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Turno{}, fmt.Errorf("Turno with ID %d not found", id)
		}
		return domain.Turno{}, err
	}
	return t, nil
}

func (r *MySQLRepository) GetByPacienteDNI(dni string) (TurnoDetail, error) {
	query := `
		SELECT p.nombre, p.apellido, p.domicilio, p.dni, p.fecha_alta, o.nombre, o.apellido, o.matricula, t.fechaHora, t.descripcion
		FROM turnos t
		JOIN pacientes p ON t.paciente = p.id
		JOIN odontologos o ON t.odontologo = o.id
		WHERE p.dni = ?
		ORDER BY t.fechaHora DESC
		LIMIT 1
	`
	row := r.db.QueryRow(query, dni)
	var turno TurnoDetail

	err := row.Scan(&turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI,
		&turno.Paciente.FechaAlta, &turno.Odontologo.Nombre, &turno.Odontologo.Apellido, &turno.Odontologo.Matricula,
		&turno.FechaHora, &turno.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return TurnoDetail{}, fmt.Errorf("No turno found for the provided DNI")
		}
		return TurnoDetail{}, err
	}
	return turno, nil
}

func (r *MySQLRepository) Create(t domain.Turno) (domain.Turno, error) {
	query := "INSERT INTO turnos (paciente, odontologo, fechaHora, descripcion) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, t.PacienteId, t.OdontologoId, t.FechaHora, t.Descripcion)
	if err != nil {
		return domain.Turno{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.Turno{}, err
	}
	t.Id = int(id)
	return t, nil
}

func (r *MySQLRepository) Update(id int, t domain.Turno) (domain.Turno, error) {
	query := "UPDATE turnos SET paciente = ?, odontologo = ?, fechaHora = ?, descripcion = ? WHERE id = ?"
	_, err := r.db.Exec(query, t.PacienteId, t.OdontologoId, t.FechaHora, t.Descripcion, id)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}

func (r *MySQLRepository) Delete(id int) error {
	query := "DELETE FROM turnos WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Turno with ID %d not found", id)
	}
	return nil
}
