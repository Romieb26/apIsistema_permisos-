package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain/entities"
)

type GetAllUsersUseCase struct {
	repo domain.IUser
}

func NewGetAllUsersUseCase(repo domain.IUser) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{repo: repo}
}

func (uc *GetAllUsersUseCase) Run() ([]entities.User, error) {
	return uc.repo.GetAll()
}
