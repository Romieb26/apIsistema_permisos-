package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Romieb26/ApIsistema_permisos/src/user/application"
)

type LoginUserController struct {
	loginUseCase *application.LoginUserUseCase
}

func NewLoginUserController(loginUseCase *application.LoginUserUseCase) *LoginUserController {
	return &LoginUserController{loginUseCase: loginUseCase}
}

func (ctrl *LoginUserController) Run(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	user, err := ctrl.loginUseCase.Run(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Credenciales inválidas",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
