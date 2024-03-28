package odontologo

import (
	"database/sql"
	"fmt"

	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
)

type MySQLRepository struct {
    db *sql.DB
}

type Repository interface {
	GetAll() ([]domain.Odontologo, error)
	GetByID(id int) (domain.Odontologo, error)
    GetByMat(mat string) (domain.Odontologo, error)
	Create(dentist domain.Odontologo) (domain.Odontologo, error)
	Update(id int, dentist domain.Odontologo) (domain.Odontologo, error)
	Delete(id int) error
}


func NewMySQLRepository(db *sql.DB) *MySQLRepository {
    return &MySQLRepository{db}
}


func (r *MySQLRepository) GetAll() ([]domain.Odontologo, error) {
    rows, err := r.db.Query("SELECT id, nombre, apellido, matricula FROM odontologos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var odontologos []domain.Odontologo
    for rows.Next() {
        var o domain.Odontologo
        err := rows.Scan(&o.ID, &o.Nombre, &o.Apellido, &o.Matricula)
        if err != nil {
            return nil, err
        }
        odontologos = append(odontologos, o)
    }
    return odontologos, nil
}


func (r *MySQLRepository) GetByID(id int) (domain.Odontologo, error) {
    var o domain.Odontologo
    err := r.db.QueryRow("SELECT id, nombre, apellido, matricula FROM odontologos WHERE id = ?", id).
        Scan(&o.ID, &o.Nombre, &o.Apellido, &o.Matricula)
    if err != nil {
        return domain.Odontologo{}, err
    }
    return o, nil
}

func (r *MySQLRepository) GetByMat(mat string) (domain.Odontologo, error) {
	query := "SELECT * FROM odontologos WHERE matricula = ?"
	row := r.db.QueryRow(query, mat)
	var o domain.Odontologo

	err := row.Scan(&o.ID, &o.Nombre, &o.Apellido, &o.Matricula)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Odontologo{}, fmt.Errorf("Odontologo with matricula %s not found", mat)
		}
		return domain.Odontologo{}, err
	}
	return o, nil
}

func (r *MySQLRepository) Create(dentist domain.Odontologo) (domain.Odontologo, error) {
    result, err := r.db.Exec("INSERT INTO odontologos (nombre, apellido, matricula) VALUES (?, ?, ?)",
        dentist.Nombre, dentist.Apellido, dentist.Matricula)
    if err != nil {
        return domain.Odontologo{}, err
    }
    lastInsertID, err := result.LastInsertId()
    if err != nil {
        return domain.Odontologo{}, err
    }
    dentist.ID = int(lastInsertID)
    return dentist, nil
}


func (r *MySQLRepository) Update(id int, dentist domain.Odontologo) (domain.Odontologo, error) {
    _, err := r.db.Exec("UPDATE odontologos SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?",
        dentist.Nombre, dentist.Apellido, dentist.Matricula, id)
    if err != nil {
        return domain.Odontologo{}, err
    }
    dentist.ID = id
    return dentist, nil
}


func (r *MySQLRepository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM odontologos WHERE id = ?", id)
    return err
}
