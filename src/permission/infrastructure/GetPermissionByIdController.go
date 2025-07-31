package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/application"
)

type GetPermissionByIdController struct {
	useCase *application.GetPermissionByIdUseCase
}

func NewGetPermissionByIdController(useCase *application.GetPermissionByIdUseCase) *GetPermissionByIdController {
	return &GetPermissionByIdController{useCase: useCase}
}

func (ctrl *GetPermissionByIdController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	permission, err := ctrl.useCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, permission)
}
