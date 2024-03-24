package odontologo


import (
	"database/sql"
	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db}
}

type Repository interface {
	GetAll() ([]domain.Odontologo, error)
	GetByID(id int) (domain.Odontologo, error)
	Create(dentist domain.Odontologo) (domain.Odontologo, error)
	Update(id int, dentist domain.Odontologo) (domain.Odontologo, error)
	Delete(id int) error
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
