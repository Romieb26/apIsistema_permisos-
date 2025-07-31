package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type PermissionRouter struct {
	engine *gin.Engine
}

func NewPermissionRouter(engine *gin.Engine) *PermissionRouter {
	return &PermissionRouter{
		engine: engine,
	}
}

func (router *PermissionRouter) Run() {
	// Inicializar dependencias
	saveController, getByIdController, updateController, deleteController, getAllController := InitPermissionDependencies()

	permissionGroup := router.engine.Group("/permissions")
	{
		permissionGroup.POST("/", saveController.Run)
		permissionGroup.GET("/:id", getByIdController.Run)
		permissionGroup.PUT("/:id", updateController.Run)
		permissionGroup.DELETE("/:id", deleteController.Run)
		permissionGroup.GET("/", getAllController.Run)
	}
}
