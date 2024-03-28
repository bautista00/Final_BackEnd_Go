package domain

import "encoding/json"

type Turno struct {
	Id           int    `json:"id" example:"1"`
	PacienteId   int    `json:"paciente" example:"1"`
	OdontologoId int    `json:"odontologo" example:"1"`
	FechaHora        string `json:"fechaHora" example:"2006-01-02 15:04:05"`
	Descripcion  string `json:"descripcion" example:"string"`
}

func (t *Turno) UnmarshalJSON(data []byte) error {
	var aux struct {
		Id           int    `json:"id" example:"1"`
		PacienteId   int    `json:"paciente" example:"1"`
		OdontologoId int    `json:"odontologo" example:"1"`
		FechaHora       string `json:"fechaHora" example:"2006-01-02 15:04:05"`
		Descripcion  string `json:"descripcion" example:"string"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.Id = aux.Id
	t.PacienteId = aux.PacienteId
	t.OdontologoId = aux.OdontologoId
	t.FechaHora = aux.FechaHora
	t.Descripcion = aux.Descripcion

	return nil
}