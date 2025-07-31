package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/application"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
)

type UpdateDocenteController struct {
	useCase *application.UpdateDocenteUseCase
}

func NewUpdateDocenteController(useCase *application.UpdateDocenteUseCase) *UpdateDocenteController {
	return &UpdateDocenteController{useCase: useCase}
}

func (ctrl *UpdateDocenteController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	var request struct {
		Name     string `json:"name"`
		LastName string `json:"lastname"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos", "error": err.Error()})
		return
	}

	docente := entities.NewDocente(request.Name, request.LastName, request.Email)
	docente.ID = int32(id)

	updated, err := ctrl.useCase.Run(docente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al actualizar", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
