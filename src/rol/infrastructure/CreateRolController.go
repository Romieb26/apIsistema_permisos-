package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/rol/application"
	"github.com/Romieb26/ApIsistema_permisos/src/rol/domain/entities"
)

type CreateRolController struct {
	createUseCase *application.CreateRolUseCase
}

func NewCreateRolController(createUseCase *application.CreateRolUseCase) *CreateRolController {
	return &CreateRolController{createUseCase: createUseCase}
}

func (ctrl *CreateRolController) Run(c *gin.Context) {
	var rolRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&rolRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos", "error": err.Error()})
		return
	}

	rol := entities.NewRol(rolRequest.Title, rolRequest.Description)

	createdRol, err := ctrl.createUseCase.Run(rol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo crear el rol", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdRol)
}
