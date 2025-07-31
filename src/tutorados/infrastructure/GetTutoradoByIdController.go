package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/application"
)

type GetTutoradoByIdController struct {
	getByIdUseCase *application.GetTutoradoByIdUseCase
}

func NewGetTutoradoByIdController(useCase *application.GetTutoradoByIdUseCase) *GetTutoradoByIdController {
	return &GetTutoradoByIdController{
		getByIdUseCase: useCase,
	}
}

func (ctrl *GetTutoradoByIdController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID inválido", "error": err.Error()})
		return
	}

	tutorado, err := ctrl.getByIdUseCase.Run(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo obtener el tutorado", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tutorado)
}
