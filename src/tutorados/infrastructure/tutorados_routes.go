package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type TutoradosRouter struct {
	engine *gin.Engine
}

func NewTutoradosRouter(engine *gin.Engine) *TutoradosRouter {
	return &TutoradosRouter{
		engine: engine,
	}
}

func (router *TutoradosRouter) Run() {
	saveController, getByIdController, updateController, deleteController, getAllController := InitTutoradoDependencies()

	tutoradosGroup := router.engine.Group("/tutorados")
	{
		tutoradosGroup.POST("/", saveController.Run)
		tutoradosGroup.GET("/:id", getByIdController.Run)
		tutoradosGroup.PUT("/:id", updateController.Run)
		tutoradosGroup.DELETE("/:id", deleteController.Run)
		tutoradosGroup.GET("/", getAllController.Run)
	}
}
