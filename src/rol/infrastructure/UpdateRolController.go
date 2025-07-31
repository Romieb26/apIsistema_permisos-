package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/rol/application"
	"github.com/Romieb26/ApIsistema_permisos/src/rol/domain/entities"
)

type UpdateRolController struct {
	updateUseCase *application.UpdateRolUseCase
}

func NewUpdateRolController(updateUseCase *application.UpdateRolUseCase) *UpdateRolController {
	return &UpdateRolController{updateUseCase: updateUseCase}
}

func (ctrl *UpdateRolController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	var rolRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&rolRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos", "error": err.Error()})
		return
	}

	rol := entities.NewRol(rolRequest.Title, rolRequest.Description)
	rol.ID = int32(id)

	updatedRol, err := ctrl.updateUseCase.Run(rol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo actualizar el rol", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRol)
}
