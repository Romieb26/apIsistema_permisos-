package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/application"
)

type GetDocenteByIdController struct {
	useCase *application.GetDocenteByIdUseCase
}

func NewGetDocenteByIdController(useCase *application.GetDocenteByIdUseCase) *GetDocenteByIdController {
	return &GetDocenteByIdController{useCase: useCase}
}

func (ctrl *GetDocenteByIdController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	docente, err := ctrl.useCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Docente no encontrado", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, docente)
}
