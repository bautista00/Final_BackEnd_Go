package main

import (
	"github.com/bautista00/Final_BackEnd_Go/cmd/server/handler"

	"github.com/bautista00/Final_BackEnd_Go/internal/odontologo"

	"github.com/bautista00/Final_BackEnd_Go/internal/paciente"

	turnos "github.com/bautista00/Final_BackEnd_Go/internal/turno"

	"github.com/bautista00/Final_BackEnd_Go/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {

	storage := db.StorageDB
	defer storage.Close()

	odontologoRepo := odontologo.NewMySQLRepository(storage)
	odontologoService := odontologo.NewService(odontologoRepo)
	odontologoHandler := handler.NewOdontologoHandler(odontologoService)

	pacienteRepo := paciente.NewMySQLRepository(storage)
	pacienteService := paciente.NewService(pacienteRepo)
	pacienteHandler := handler.NewPacienteHandler(pacienteService)

	turnoRepo := turnos.NewMySQLRepository(storage)
	turnoService := turnos.NewService(turnoRepo, pacienteRepo, odontologoRepo)
	turnoHandler := handler.NewTurnoHandler(turnoService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// _______ENPOINTS ODONTOLOGOS___________

	odontologos := r.Group("/odontologos")
	{
		odontologos.GET("/all", odontologoHandler.GetAll())
		odontologos.GET("/:id", odontologoHandler.GetByID())
		odontologos.POST("/create", odontologoHandler.Create())
		odontologos.PUT("/:id", odontologoHandler.Update())
		odontologos.DELETE("/:id", odontologoHandler.Delete())
	}

	// _______ENDPOINTS PACIENTES__________

	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("/all", pacienteHandler.GetAll())
		pacientes.GET("/:id", pacienteHandler.GetByID())
		pacientes.POST("/create", pacienteHandler.Create())
		pacientes.PUT("/:id", pacienteHandler.Update())
		pacientes.DELETE("/:id", pacienteHandler.Delete())
	}

	// _______ENDPOINTS TURNOS____________

	turnos := r.Group("/turnos")
	{
		turnos.GET("/:id", turnoHandler.GetByID())
		turnos.POST("/create/dniAndMat", turnoHandler.CreateByDniAndMatricula())
		turnos.POST("/create", turnoHandler.Create())
		turnos.PUT("/:id", turnoHandler.Update())
		turnos.DELETE("/:id", turnoHandler.Delete())
	}

	r.Run(":8080")
}
