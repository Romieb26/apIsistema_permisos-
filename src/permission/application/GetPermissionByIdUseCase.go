package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain/entities"
)

type GetPermissionByIdUseCase struct {
	repo domain.IPermission
}

func NewGetPermissionByIdUseCase(repo domain.IPermission) *GetPermissionByIdUseCase {
	return &GetPermissionByIdUseCase{repo: repo}
}

func (uc *GetPermissionByIdUseCase) Run(id int32) (*entities.Permission, error) {
	return uc.repo.GetById(id)
}
