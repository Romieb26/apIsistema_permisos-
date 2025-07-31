package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/application"
)

type GetAllPermissionsController struct {
	useCase *application.GetAllPermissionsUseCase
}

func NewGetAllPermissionsController(useCase *application.GetAllPermissionsUseCase) *GetAllPermissionsController {
	return &GetAllPermissionsController{useCase: useCase}
}

func (ctrl *GetAllPermissionsController) Run(c *gin.Context) {
	permissions, err := ctrl.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, permissions)
}
