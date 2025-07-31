package infrastructure

import (
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/application"
)

func InitTutoradoDependencies() (
	*SaveTutoradoController,
	*GetTutoradoByIdController,
	*UpdateTutoradoController,
	*DeleteTutoradoController,
	*GetAllTutoradosController,
) {
	repo := NewMySQLTutoradoRepository()

	saveUseCase := application.NewSaveTutoradoUseCase(repo)
	getByIdUseCase := application.NewGetTutoradoByIdUseCase(repo)
	updateUseCase := application.NewUpdateTutoradoUseCase(repo)
	deleteUseCase := application.NewDeleteTutoradoUseCase(repo)
	getAllUseCase := application.NewGetAllTutoradosUseCase(repo)

	saveController := NewSaveTutoradoController(saveUseCase)
	getByIdController := NewGetTutoradoByIdController(getByIdUseCase)
	updateController := NewUpdateTutoradoController(updateUseCase)
	deleteController := NewDeleteTutoradoController(deleteUseCase)
	getAllController := NewGetAllTutoradosController(getAllUseCase)

	return saveController, getByIdController, updateController, deleteController, getAllController
}
