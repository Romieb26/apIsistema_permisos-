package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/application"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain/entities"
)

type UpdateTutoradoController struct {
	updateUseCase *application.UpdateTutoradoUseCase
}

func NewUpdateTutoradoController(useCase *application.UpdateTutoradoUseCase) *UpdateTutoradoController {
	return &UpdateTutoradoController{
		updateUseCase: useCase,
	}
}

func (ctrl *UpdateTutoradoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	var tutoradoRequest struct {
		IdUserFK  int32  `json:"id_user_fk"`
		Name      string `json:"name"`
		LastName  string `json:"lastname"`
		Matricula string `json:"matricula"`
		Email     string `json:"email"`
		Estatus   string `json:"estatus"`
	}

	if err := c.ShouldBindJSON(&tutoradoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos", "error": err.Error()})
		return
	}

	tutorado := entities.NewTutorado(
		tutoradoRequest.IdUserFK,
		tutoradoRequest.Name,
		tutoradoRequest.LastName,
		tutoradoRequest.Matricula,
		tutoradoRequest.Email,
		tutoradoRequest.Estatus,
	)
	tutorado.ID = int32(id)

	updatedTutorado, err := ctrl.updateUseCase.Run(tutorado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo actualizar el tutorado", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTutorado)
}
