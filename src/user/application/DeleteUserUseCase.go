package application

import "github.com/Romieb26/ApIsistema_permisos/src/user/domain"

type DeleteUserUseCase struct {
	repo domain.IUser
}

func NewDeleteUserUseCase(repo domain.IUser) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: repo}
}

func (uc *DeleteUserUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}
