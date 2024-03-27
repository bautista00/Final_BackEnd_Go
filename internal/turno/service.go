package turno

import (
	
	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Turno, error)
	GetByID(id int) (domain.Turno, error)
	Create(turno domain.Turno, pacienteID int, odontologoID int) (domain.Turno, error)
	Update(id int, turno domain.Turno) (domain.Turno, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Turno, error) {
	turnos, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return turnos, nil
}

func (s *service) GetByID(id int) (domain.Turno, error) {
	turno, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (s *service) Create(turno domain.Turno, pacienteID int, odontologoID int) (domain.Turno, error) {
	createdTurno, err := s.r.Create(turno, pacienteID,odontologoID)
	if err != nil {
		return domain.Turno{}, err
	}
	return createdTurno, nil
}

func (s *service) Update(id int, turno domain.Turno) (domain.Turno, error) {
	updatedTurno, err := s.r.Update(id, turno)
	if err != nil {
		return domain.Turno{}, err
	}
	return updatedTurno, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
