package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain/entities"
)

type SaveUserUseCase struct {
	repo domain.IUser
}

func NewSaveUserUseCase(repo domain.IUser) *SaveUserUseCase {
	return &SaveUserUseCase{repo: repo}
}

func (uc *SaveUserUseCase) Run(user *entities.User) (*entities.User, error) {
	err := uc.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
