package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/application"
)

type DeleteTutoradoController struct {
	deleteUseCase *application.DeleteTutoradoUseCase
}

func NewDeleteTutoradoController(useCase *application.DeleteTutoradoUseCase) *DeleteTutoradoController {
	return &DeleteTutoradoController{
		deleteUseCase: useCase,
	}
}

func (ctrl *DeleteTutoradoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(int32(id))
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo eliminar el tutorado", "error": errDelete.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Tutorado eliminado exitosamente"})
}
