package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
	"github.com/bautista00/Final_BackEnd_Go/internal/turno"
	"github.com/bautista00/Final_BackEnd_Go/pkg/web"
    
	"github.com/gin-gonic/gin"
)

type TurnoHandler struct {
    s turno.Service
}

func NewTurnoHandler(s turno.Service) *TurnoHandler {
    return &TurnoHandler{
        s: s,
    }
}

func (Th *TurnoHandler) Create() gin.HandlerFunc {
    return func(c *gin.Context) {
        var nuevoTurno domain.Turno
        if err := c.BindJSON(&nuevoTurno); err != nil {
            web.Failure(c, http.StatusBadRequest, fmt.Errorf("Datos de turno no válidos: %s", err))
            return
        }
        creado, err := Th.s.Create(nuevoTurno, nuevoTurno.Paciente.ID, nuevoTurno.Odontologo.ID)
        if err != nil {
            web.Failure(c, http.StatusInternalServerError, err) 
            return
        }
        web.Success(c, http.StatusCreated, creado)
    }
}


func (Th *TurnoHandler) GetByID() gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id") 
        id, err := strconv.Atoi(idStr) 
        if err != nil {
            web.Failure(c, http.StatusBadRequest, fmt.Errorf("Id de turno no válido"))
            return
        }
        turno, err := Th.s.GetByID(id)
        if err != nil {
            web.Failure(c, http.StatusNotFound, fmt.Errorf("Turno no encontrado"))
            return
        }
        web.Success(c, http.StatusOK, turno)
    }
}

func (Th *TurnoHandler) Update() gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            web.Failure(c, http.StatusBadRequest, fmt.Errorf("Id de turno no válido"))
            return
        }

        var turnoActualizado domain.Turno
        if err := c.BindJSON(&turnoActualizado); err != nil {
            web.Failure(c, http.StatusBadRequest, fmt.Errorf("Datos de turno no válidos"))
            return
        }

        turnoActualizado.ID = id 
        turnoActualizado, err = Th.s.Update(id, turnoActualizado)
        if err != nil {
            web.Failure(c, http.StatusInternalServerError, err)
            return
        }

        web.Success(c, http.StatusOK, turnoActualizado)
    }
}


func (Th *TurnoHandler) Delete() gin.HandlerFunc {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            web.Failure(c, http.StatusBadRequest, fmt.Errorf("id de turno no válido"))
            return
        }
        err = Th.s.Delete(id)
        if err != nil {
            web.Failure(c, http.StatusNotFound, fmt.Errorf("Turno no encontrado"))
            return
        }
        web.Success(c, http.StatusOK, "Turno eliminado exitosamente")
    }
}

