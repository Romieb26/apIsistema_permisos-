package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type DocenteRouter struct {
	engine *gin.Engine
}

func NewDocenteRouter(engine *gin.Engine) *DocenteRouter {
	return &DocenteRouter{
		engine: engine,
	}
}

func (router *DocenteRouter) Run() {
	saveController, getByIdController, updateController, deleteController, getAllController := InitDocenteDependencies()

	docenteGroup := router.engine.Group("/docentes")
	{
		docenteGroup.POST("/", saveController.Run)
		docenteGroup.GET("/:id", getByIdController.Run)
		docenteGroup.PUT("/:id", updateController.Run)
		docenteGroup.DELETE("/:id", deleteController.Run)
		docenteGroup.GET("/", getAllController.Run)
	}
}
