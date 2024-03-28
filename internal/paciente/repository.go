package paciente

import (
	"database/sql"
	"fmt"

	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db}
}

type Repository interface {
	GetAll() ([]domain.Paciente, error)
	GetByID(id int) (domain.Paciente, error)
	Create(paciente domain.Paciente) (domain.Paciente, error)
	GetByDoc(doc string) (domain.Paciente, error)
	Update(id int, paciente domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

func (r *MySQLRepository) GetAll() ([]domain.Paciente, error) {
	rows, err := r.db.Query("SELECT id, nombre, apellido, domicilio, dni, fecha_alta FROM pacientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pacientes []domain.Paciente
	for rows.Next() {
		var p domain.Paciente
		err := rows.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Domicilio, &p.DNI, &p.FechaAlta)
		if err != nil {
			return nil, err
		}
		pacientes = append(pacientes, p)
	}
	return pacientes, nil
}

func (r *MySQLRepository) GetByID(id int) (domain.Paciente, error) {
	row := r.db.QueryRow("SELECT id, nombre, apellido, domicilio, dni, fecha_alta FROM pacientes WHERE id = ?", id)
	var p domain.Paciente
	err := row.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Domicilio, &p.DNI, &p.FechaAlta)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}


func (r *MySQLRepository) GetByDoc(doc string) (domain.Paciente, error) {
	query := "SELECT * FROM pacientes WHERE dni = ?"
	row := r.db.QueryRow(query, doc)
	var paciente domain.Paciente

	err := row.Scan(&paciente.ID, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaAlta)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Paciente{}, fmt.Errorf("Paciente with DOC %s not found", doc)
		}
		return domain.Paciente{}, err
	}
	return paciente, nil
}


func (r *MySQLRepository) Create(paciente domain.Paciente) (domain.Paciente, error) {
	result, err := r.db.Exec("INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_alta) VALUES (?, ?, ?, ?, ?)",
		paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.FechaAlta)
	if err != nil {
		return domain.Paciente{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.Paciente{}, err
	}
	paciente.ID = int(id)
	return paciente, nil
}

func (r *MySQLRepository) Update(id int, paciente domain.Paciente) (domain.Paciente, error) {
	_, err := r.db.Exec("UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_alta = ? WHERE id = ?",
		paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.FechaAlta, id)
	if err != nil {
		return domain.Paciente{}, err
	}
	paciente.ID = id
	return paciente, nil
}

func (r *MySQLRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM pacientes WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
