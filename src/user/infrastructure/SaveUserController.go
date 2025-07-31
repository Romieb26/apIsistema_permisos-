package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/user/application"
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain/entities"
)

type SaveUserController struct {
	saveUseCase *application.SaveUserUseCase
}

func NewSaveUserController(saveUseCase *application.SaveUserUseCase) *SaveUserController {
	return &SaveUserController{saveUseCase: saveUseCase}
}

func (ctrl *SaveUserController) Run(c *gin.Context) {
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

	createdUser, err := ctrl.saveUseCase.Run(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear el usuario",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
