package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain/entities"
)

type SavePermissionUseCase struct {
	repo domain.IPermission
}

func NewSavePermissionUseCase(repo domain.IPermission) *SavePermissionUseCase {
	return &SavePermissionUseCase{repo: repo}
}

func (uc *SavePermissionUseCase) Run(permission *entities.Permission) (*entities.Permission, error) {
	err := uc.repo.Save(permission)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
