package handler

import (
	"github.com/bautista00/Final_BackEnd_Go/internal/paciente"
	"github.com/bautista00/Final_BackEnd_Go/pkg/web"
	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

func (h *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		pacientes, err := h.s.GetAll()
		if err != nil {
			panic(err)
		}
		web.Success(c, 200, pacientes)
	}
}
