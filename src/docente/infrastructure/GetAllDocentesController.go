package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/application"
)

type GetAllDocentesController struct {
	useCase *application.GetAllDocentesUseCase
}

func NewGetAllDocentesController(useCase *application.GetAllDocentesUseCase) *GetAllDocentesController {
	return &GetAllDocentesController{useCase: useCase}
}

func (ctrl *GetAllDocentesController) Run(c *gin.Context) {
	docentes, err := ctrl.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al obtener docentes", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, docentes)
}
