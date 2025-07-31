package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/application"
)

type DeletePermissionController struct {
	useCase *application.DeletePermissionUseCase
}

func NewDeletePermissionController(useCase *application.DeletePermissionUseCase) *DeletePermissionController {
	return &DeletePermissionController{useCase: useCase}
}

func (ctrl *DeletePermissionController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.useCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permiso eliminado correctamente"})
}
