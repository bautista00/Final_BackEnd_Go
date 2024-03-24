package odontologo

import (
	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Odontologo, error)
	GetByID(id int) (domain.Odontologo, error)
	Create(odontologo domain.Odontologo) (domain.Odontologo, error)
	Update(id int, odontologo domain.Odontologo) (domain.Odontologo, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Odontologo, error) {
	odontologos, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return odontologos, nil
}

func (s *service) GetByID(id int) (domain.Odontologo, error) {
	odontologo, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (s *service) Create(odontologo domain.Odontologo) (domain.Odontologo, error) {
	createdOdontologo, err := s.r.Create(odontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return createdOdontologo, nil
}

func (s *service) Update(id int, odontologo domain.Odontologo) (domain.Odontologo, error) {
	updatedOdontologo, err := s.r.Update(id, odontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return updatedOdontologo, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

