package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
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

func (Ph *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		pacientes, err := Ph.s.GetAll()
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, fmt.Errorf("Error al obtener todos los pacientes: %v", err))
			return
		}
		web.Success(c, http.StatusOK, pacientes)
	}
}

func (Ph *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Id de paciente no válido"))
			return
		}

		paciente, err := Ph.s.GetByID(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, fmt.Errorf("Paciente no encontrado: %v", err))
			return
		}
		web.Success(c, http.StatusOK, paciente)
	}
}

func (Ph *pacienteHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newPaciente domain.Paciente
		if err := c.BindJSON(&newPaciente); err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Datos de paciente no válidos"))
			return
		}

		createdPaciente, err := Ph.s.Create(newPaciente)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, fmt.Errorf("Error al crear el paciente: %v", err))
			return
		}
		web.Success(c, http.StatusCreated, createdPaciente)
	}
}

func (Ph *pacienteHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Id de paciente no válido"))
			return
		}

		var updatedPaciente domain.Paciente
		if err := c.BindJSON(&updatedPaciente); err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Datos de paciente no válidos"))
			return
		}

		updatedPaciente, err = Ph.s.Update(id, updatedPaciente)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, fmt.Errorf("Error al actualizar el paciente: %v", err))
			return
		}
		web.Success(c, http.StatusOK, updatedPaciente)
	}
}

func (Ph *pacienteHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Id de paciente no válido"))
			return
		}

		err = Ph.s.Delete(id)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, fmt.Errorf("Error al eliminar el paciente: %v", err))
			return
		}
		web.Success(c, http.StatusOK, "Paciente eliminado exitosamente")
	}
}
