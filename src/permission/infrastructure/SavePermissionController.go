package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/application"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain/entities"
)

type SavePermissionController struct {
	useCase *application.SavePermissionUseCase
}

func NewSavePermissionController(useCase *application.SavePermissionUseCase) *SavePermissionController {
	return &SavePermissionController{useCase: useCase}
}

func (ctrl *SavePermissionController) Run(c *gin.Context) {
	var req struct {
		TutoradoID int32  `json:"id_tutorado_fk"`
		DocenteID  int32  `json:"id_docente_fk"`
		Evidencia  string `json:"evidencia"`
		Date       string `json:"date"`
		Motivo     string `json:"motivo"`
		Estatus    string `json:"estatus"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission := entities.NewPermission(
		req.TutoradoID,
		req.DocenteID,
		req.Evidencia,
		req.Date,
		req.Motivo,
		req.Estatus,
	)

	result, err := ctrl.useCase.Run(permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
