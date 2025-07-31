package infrastructure

import (
	"github.com/Romieb26/ApIsistema_permisos/src/user/application"
)

func InitUserDependencies() (
	*SaveUserController,
	*GetUserByIdController,
	*UpdateUserController,
	*DeleteUserController,
	*GetAllUsersController,
	*LoginUserController,
) {
	// Inicializar repositorio
	repo := NewMySQLUserRepository()

	// Inicializar casos de uso
	saveUseCase := application.NewSaveUserUseCase(repo)
	getByIdUseCase := application.NewGetUserByIdUseCase(repo)
	updateUseCase := application.NewUpdateUserUseCase(repo)
	deleteUseCase := application.NewDeleteUserUseCase(repo)
	getAllUseCase := application.NewGetAllUsersUseCase(repo)
	loginUseCase := application.NewLoginUserUseCase(repo)

	// Inicializar controladores
	saveController := NewSaveUserController(saveUseCase)
	getByIdController := NewGetUserByIdController(getByIdUseCase)
	updateController := NewUpdateUserController(updateUseCase)
	deleteController := NewDeleteUserController(deleteUseCase)
	getAllController := NewGetAllUsersController(getAllUseCase)
	loginController := NewLoginUserController(loginUseCase)

	return saveController, getByIdController, updateController, deleteController, getAllController, loginController
}
