package paciente

import (
	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Paciente, error)
	GetByID(id int) (domain.Paciente, error)
	Create(paciente domain.Paciente) (domain.Paciente, error)
	Update(id int, paciente domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Paciente, error) {
	pacientes, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return pacientes, nil
}

func (s *service) GetByID(id int) (domain.Paciente, error) {
	paciente, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (s *service) Create(paciente domain.Paciente) (domain.Paciente, error) {
	createdPaciente, err := s.r.Create(paciente)
	if err != nil {
		return domain.Paciente{}, err
	}
	return createdPaciente, nil
}

func (s *service) Update(id int, paciente domain.Paciente) (domain.Paciente, error) {
	updatedPaciente, err := s.r.Update(id, paciente)
	if err != nil {
		return domain.Paciente{}, err
	}
	return updatedPaciente, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
