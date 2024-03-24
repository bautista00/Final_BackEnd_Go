package main

import(
	"github.com/bautista00/Final_BackEnd_Go/cmd/server/handler"
	"github.com/bautista00/Final_BackEnd_Go/internal/odontologo"
	"github.com/bautista00/Final_BackEnd_Go/pkg/db"
	"github.com/gin-gonic/gin"
	
)

func main() {
	//if err := godotenv.Load("../.././.env"); err != nil {
	//	panic("Error loading .env file: " + err.Error())
	//}

	storage := db.StorageDB

	repo := odontologo.NewMySQLRepository(storage)
	service := odontologo.NewService(repo)
	odontologoHandler := handler.NewOdontologoHandler(service)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	odontologos := r.Group("/odontologos")
	{
		odontologos.GET("/all", odontologoHandler.GetAll())
	}
	r.Run(":8080")
}