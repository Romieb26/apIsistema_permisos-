package application

import (
	repositories "github.com/Romieb26/ApIsistema_permisos/src/rol/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/rol/domain/entities"
)

type UpdateRolUseCase struct {
	repo repositories.IRol
}

func NewUpdateRolUseCase(repo repositories.IRol) *UpdateRolUseCase {
	return &UpdateRolUseCase{repo: repo}
}

func (uc *UpdateRolUseCase) Run(rol *entities.Rol) (*entities.Rol, error) {
	err := uc.repo.Update(rol)
	if err != nil {
		return nil, err
	}
	return rol, nil
}
