package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain/entities"
)

type GetAllPermissionsUseCase struct {
	repo domain.IPermission
}

func NewGetAllPermissionsUseCase(repo domain.IPermission) *GetAllPermissionsUseCase {
	return &GetAllPermissionsUseCase{repo: repo}
}

func (uc *GetAllPermissionsUseCase) Run() ([]entities.Permission, error) {
	return uc.repo.GetAll()
}
