package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/application"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain/entities"
)

type UpdatePermissionController struct {
	useCase *application.UpdatePermissionUseCase
}

func NewUpdatePermissionController(useCase *application.UpdatePermissionUseCase) *UpdatePermissionController {
	return &UpdatePermissionController{useCase: useCase}
}

func (ctrl *UpdatePermissionController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

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
	permission.ID = int32(id)

	result, err := ctrl.useCase.Run(permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
