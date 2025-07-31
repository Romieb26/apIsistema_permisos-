package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/application"
)

type GetAllTutoradosController struct {
	getAllUseCase *application.GetAllTutoradosUseCase
}

func NewGetAllTutoradosController(useCase *application.GetAllTutoradosUseCase) *GetAllTutoradosController {
	return &GetAllTutoradosController{
		getAllUseCase: useCase,
	}
}

func (ctrl *GetAllTutoradosController) Run(c *gin.Context) {
	tutorados, err := ctrl.getAllUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudieron obtener los tutorados", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tutorados)
}
