package domain

import (
	"encoding/json"
	"time"
	
)

type Turno struct {
	ID          int
	Paciente 	Paciente
	Odontologo  Odontologo
	FechaHora   time.Time
	Descripcion string
}

func (t *Turno) UnmarshalJSON(data []byte) error {
	var aux struct {
		Id           int    `json:"id" example:"1"`
		Paciente  Paciente    `json:"id_paciente" example:"1"`
		Odontologo Odontologo    `json:"id_odontologo" example:"1"`
		FechaHora        time.Time `json:"fecha" example:"2006-01-02 15:04:05"`
		Descripcion  string `json:"descripcion" example:"string"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.ID = aux.Id
	t.Paciente = aux.Paciente
	t.Odontologo = aux.Odontologo
	t.FechaHora = aux.FechaHora
	t.Descripcion = aux.Descripcion

	return nil
}