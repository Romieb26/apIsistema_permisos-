package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/application"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
)

type SaveDocenteController struct {
	useCase *application.SaveDocenteUseCase
}

func NewSaveDocenteController(useCase *application.SaveDocenteUseCase) *SaveDocenteController {
	return &SaveDocenteController{useCase: useCase}
}

func (ctrl *SaveDocenteController) Run(c *gin.Context) {
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

	saved, err := ctrl.useCase.Run(docente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, saved)
}
