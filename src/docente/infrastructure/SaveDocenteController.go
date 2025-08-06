package infrastructure

import (
	"net/http"
	"strings"

	"github.com/Romieb26/ApIsistema_permisos/src/docente/application"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
	"github.com/gin-gonic/gin"
)

type SaveDocenteController struct {
	useCase *application.SaveDocenteUseCase
}

func NewSaveDocenteController(useCase *application.SaveDocenteUseCase) *SaveDocenteController {
	return &SaveDocenteController{useCase: useCase}
}

func (ctrl *SaveDocenteController) Run(c *gin.Context) {
	var request struct {
		Name     string `json:"name" binding:"required"`
		LastName string `json:"lastname" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos de entrada inválidos",
			"details": err.Error(),
		})
		return
	}

	docente := entities.NewDocente(request.Name, request.LastName, request.Email)

	if err := ctrl.useCase.Execute(docente); err != nil {
		status := http.StatusBadRequest
		errorType := "validation_error"
		
		if strings.Contains(err.Error(), "servicio de validación") {
			errorType = "service_error"
		}

		c.JSON(status, gin.H{
			"error":      err.Error(),
			"error_type": errorType,
			"email":      docente.Email,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Docente registrado exitosamente",
		"data":    docente,
	})
}