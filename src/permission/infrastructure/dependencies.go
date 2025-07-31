package infrastructure

import (
	"github.com/Romieb26/ApIsistema_permisos/src/permission/application"
)

func InitPermissionDependencies() (
	*SavePermissionController,
	*GetPermissionByIdController,
	*UpdatePermissionController,
	*DeletePermissionController,
	*GetAllPermissionsController,
) {
	repo := NewMySQLPermissionRepository()

	saveUseCase := application.NewSavePermissionUseCase(repo)
	getByIdUseCase := application.NewGetPermissionByIdUseCase(repo)
	updateUseCase := application.NewUpdatePermissionUseCase(repo)
	deleteUseCase := application.NewDeletePermissionUseCase(repo)
	getAllUseCase := application.NewGetAllPermissionsUseCase(repo)

	saveController := NewSavePermissionController(saveUseCase)
	getByIdController := NewGetPermissionByIdController(getByIdUseCase)
	updateController := NewUpdatePermissionController(updateUseCase)
	deleteController := NewDeletePermissionController(deleteUseCase)
	getAllController := NewGetAllPermissionsController(getAllUseCase)

	return saveController, getByIdController, updateController, deleteController, getAllController
}
