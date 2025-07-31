package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain/entities"
)

type GetUserByIdUseCase struct {
	repo domain.IUser
}

func NewGetUserByIdUseCase(repo domain.IUser) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{repo: repo}
}

func (uc *GetUserByIdUseCase) Run(id int32) (*entities.User, error) {
	return uc.repo.GetById(id)
}
