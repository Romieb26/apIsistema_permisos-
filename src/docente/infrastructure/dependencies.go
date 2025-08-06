package infrastructure

import (
	"github.com/Romieb26/ApIsistema_permisos/src/docente/application"
)

func InitDocenteDependencies() (
	*SaveDocenteController,
	*GetDocenteByIdController,
	*UpdateDocenteController,
	*DeleteDocenteController,
	*GetAllDocentesController,
) {
	repo := NewMySQLDocenteRepository()
	emailValidator := NewSimulatedEmailValidator()

	saveUseCase := application.NewSaveDocenteUseCase(repo, emailValidator)
	getByIdUseCase := application.NewGetDocenteByIdUseCase(repo)
	updateUseCase := application.NewUpdateDocenteUseCase(repo)
	deleteUseCase := application.NewDeleteDocenteUseCase(repo)
	getAllUseCase := application.NewGetAllDocentesUseCase(repo)

	saveController := NewSaveDocenteController(saveUseCase)
	getByIdController := NewGetDocenteByIdController(getByIdUseCase)
	updateController := NewUpdateDocenteController(updateUseCase)
	deleteController := NewDeleteDocenteController(deleteUseCase)
	getAllController := NewGetAllDocentesController(getAllUseCase)

	return saveController, getByIdController, updateController, deleteController, getAllController
}