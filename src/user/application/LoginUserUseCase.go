package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain/entities"
)

type LoginUserUseCase struct {
	repo domain.IUser
}

func NewLoginUserUseCase(repo domain.IUser) *LoginUserUseCase {
	return &LoginUserUseCase{repo: repo}
}

func (uc *LoginUserUseCase) Run(username, password string) (*entities.User, error) {
	return uc.repo.Login(username, password)
}
