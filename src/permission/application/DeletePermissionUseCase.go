package application

import "github.com/Romieb26/ApIsistema_permisos/src/permission/domain"

type DeletePermissionUseCase struct {
	repo domain.IPermission
}

func NewDeletePermissionUseCase(repo domain.IPermission) *DeletePermissionUseCase {
	return &DeletePermissionUseCase{repo: repo}
}

func (uc *DeletePermissionUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}
