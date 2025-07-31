package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/application"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain/entities"
)

type SaveTutoradoController struct {
	saveUseCase *application.SaveTutoradoUseCase
}

func NewSaveTutoradoController(useCase *application.SaveTutoradoUseCase) *SaveTutoradoController {
	return &SaveTutoradoController{
		saveUseCase: useCase,
	}
}

func (ctrl *SaveTutoradoController) Run(c *gin.Context) {
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

	createdTutorado, err := ctrl.saveUseCase.Run(tutorado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo crear el tutorado", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTutorado)
}
