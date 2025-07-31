package infrastructure

import (
	"github.com/Romieb26/ApIsistema_permisos/src/rol/application"
)

func InitRolDependencies() (
	*CreateRolController,
	*GetRolByIdController,
	*UpdateRolController,
	*DeleteRolController,
	*GetAllRolesController,
) {
	// Crear instancia del repositorio (MySQL en este caso)
	repo := NewMySQLRolRepository()

	// Inyectar casos de uso
	createUseCase := application.NewCreateRolUseCase(repo)
	getByIdUseCase := application.NewGetRolByIdUseCase(repo)
	updateUseCase := application.NewUpdateRolUseCase(repo)
	deleteUseCase := application.NewDeleteRolUseCase(repo)
	getAllUseCase := application.NewGetAllRolesUseCase(repo)

	// Inyectar controladores
	createController := NewCreateRolController(createUseCase)
	getByIdController := NewGetRolByIdController(getByIdUseCase)
	updateController := NewUpdateRolController(updateUseCase)
	deleteController := NewDeleteRolController(deleteUseCase)
	getAllController := NewGetAllRolesController(getAllUseCase)

	// Retornar los controladores
	return createController, getByIdController, updateController, deleteController, getAllController
}
