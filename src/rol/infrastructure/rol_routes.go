package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type RolRouter struct {
	engine *gin.Engine
}

func NewRolRouter(engine *gin.Engine) *RolRouter {
	return &RolRouter{engine: engine}
}

func (router *RolRouter) Run() {
	// Inicializar dependencias de los controladores
	createController, getByIdController, updateController, deleteController, getAllController := InitRolDependencies()

	// Grupo de rutas para "roles"
	rolGroup := router.engine.Group("/roles")
	{
		rolGroup.POST("/", createController.Run)     // Crear rol
		rolGroup.GET("/", getAllController.Run)      // Obtener todos los roles
		rolGroup.GET("/:id", getByIdController.Run)  // Obtener un rol por ID
		rolGroup.PUT("/:id", updateController.Run)   // Actualizar rol por ID
		rolGroup.DELETE("/:id", deleteController.Run) // Eliminar rol por ID
	}
}
