package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain/entities"
)

type UpdatePermissionUseCase struct {
	repo domain.IPermission
}

func NewUpdatePermissionUseCase(repo domain.IPermission) *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{repo: repo}
}

func (uc *UpdatePermissionUseCase) Run(permission *entities.Permission) (*entities.Permission, error) {
	err := uc.repo.Update(permission)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
