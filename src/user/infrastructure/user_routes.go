package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	engine *gin.Engine
}

func NewUserRouter(engine *gin.Engine) *UserRouter {
	return &UserRouter{
		engine: engine,
	}
}

func (router *UserRouter) Run() {
	saveController, getByIdController, updateController, deleteController, getAllController, loginController := InitUserDependencies()

	userGroup := router.engine.Group("/users")
	{
		userGroup.POST("/", saveController.Run)          // Crear usuario
		userGroup.GET("/:id", getByIdController.Run)     // Obtener usuario por ID
		userGroup.PUT("/:id", updateController.Run)      // Actualizar usuario
		userGroup.DELETE("/:id", deleteController.Run)   // Eliminar usuario
		userGroup.GET("/", getAllController.Run)         // Obtener todos los usuarios
		userGroup.POST("/login", loginController.Run)    // Login usuario
	}
}
