package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/application"
)

type DeleteDocenteController struct {
	useCase *application.DeleteDocenteUseCase
}

func NewDeleteDocenteController(useCase *application.DeleteDocenteUseCase) *DeleteDocenteController {
	return &DeleteDocenteController{useCase: useCase}
}

func (ctrl *DeleteDocenteController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	if err := ctrl.useCase.Run(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al eliminar", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Docente eliminado correctamente"})
}
