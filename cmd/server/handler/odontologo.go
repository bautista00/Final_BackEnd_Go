package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bautista00/Final_BackEnd_Go/internal/domain"
	"github.com/bautista00/Final_BackEnd_Go/internal/odontologo"
	"github.com/bautista00/Final_BackEnd_Go/pkg/web"

	"github.com/gin-gonic/gin"
)

type OdontologoHandler struct {
	service odontologo.Service
}

func NewOdontologoHandler(service odontologo.Service) *OdontologoHandler {
	return &OdontologoHandler{service: service}
}

func (Oh *OdontologoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		odontologos, err := Oh.service.GetAll()
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, fmt.Errorf("Error al obtener todos los odontólogos: %v", err))
			return
		}
		web.Success(c, http.StatusOK, odontologos)
	}
}

func (Oh *OdontologoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Id de odontologo no válido"))
			return
		}

		odontologo, err := Oh.service.GetByID(idInt)
		if err != nil {
			web.Failure(c, http.StatusNotFound, fmt.Errorf("Odontólogo no encontrado: %v", err))
			return
		}
		web.Success(c, http.StatusOK, odontologo)
	}
}


func (Oh *OdontologoHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newOdontologo domain.Odontologo
		if err := c.BindJSON(&newOdontologo); err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Datos de odontologo no válidos"))
			return
		}
		createdOdontologo, err := Oh.service.Create(newOdontologo)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, fmt.Errorf("Error al crear el odontólogo: %v", err))
			return
		}
		web.Success(c, http.StatusCreated, createdOdontologo)
	}
}

func (Oh *OdontologoHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Id de odontologo no válido"))
			return
		}

		var updatedOdontologo domain.Odontologo
		if err := c.BindJSON(&updatedOdontologo); err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Datos de odontologo no válidos"))
			return
		}

		updatedOdontologo, err = Oh.service.Update(id, updatedOdontologo)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, fmt.Errorf("Error al actualizar el odontólogo: %v", err))
			return
		}
		web.Success(c, http.StatusOK, updatedOdontologo)
	}
}

func (Oh *OdontologoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, fmt.Errorf("Id de odontologo no válido"))
			return
		}

		err = Oh.service.Delete(id)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, fmt.Errorf("Error al eliminar el odontólogo: %v", err))
			return
		}
		web.Success(c, http.StatusOK, "Odontólogo eliminado exitosamente")
	}
}
