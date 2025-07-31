package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/user/application"
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain/entities"
)

type UpdateUserController struct {
	updateUseCase *application.UpdateUserUseCase
}

func NewUpdateUserController(updateUseCase *application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{updateUseCase: updateUseCase}
}

func (ctrl *UpdateUserController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	var userRequest struct {
		Name     string `json:"name"`
		LastName string `json:"lastname"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
		IdRolFK  int32  `json:"id_rol_fk"`
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	user := entities.NewUser(
		userRequest.Name,
		userRequest.LastName,
		userRequest.Email,
		userRequest.Password,
		userRequest.Username,
		userRequest.IdRolFK,
	)
	user.ID = int32(id)

	updatedUser, err := ctrl.updateUseCase.Run(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo actualizar el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
