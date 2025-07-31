package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain/entities"
)

type UpdateUserUseCase struct {
	repo domain.IUser
}

func NewUpdateUserUseCase(repo domain.IUser) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo: repo}
}

func (uc *UpdateUserUseCase) Run(user *entities.User) (*entities.User, error) {
	err := uc.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
