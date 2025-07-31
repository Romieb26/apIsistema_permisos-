package application

import (
	repositories "github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain/entities"
)

type GetAllTutoradosUseCase struct {
	repo repositories.ITutorado
}

func NewGetAllTutoradosUseCase(repo repositories.ITutorado) *GetAllTutoradosUseCase {
	return &GetAllTutoradosUseCase{repo: repo}
}

func (uc *GetAllTutoradosUseCase) Run() ([]entities.Tutorado, error) {
	return uc.repo.GetAll()
}
