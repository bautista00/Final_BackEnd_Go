package handler

import(
	"github.com/bautista00/Final_BackEnd_Go/internal/odontologo"
	"github.com/bautista00/Final_BackEnd_Go/pkg/web"
	"github.com/gin-gonic/gin"
)


type odontologoHandler struct {
	s odontologo.Service
}

func NewOdontologoHandler(s odontologo.Service) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

func (h *odontologoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		odontologos, err := h.s.GetAll()
		if err != nil {
			panic(err)
		}
		web.Success(c, 200, odontologos)
	}
}