package application

import (
	repositories "github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain/entities"
)

type GetTutoradoByIdUseCase struct {
	repo repositories.ITutorado
}

func NewGetTutoradoByIdUseCase(repo repositories.ITutorado) *GetTutoradoByIdUseCase {
	return &GetTutoradoByIdUseCase{repo: repo}
}

func (uc *GetTutoradoByIdUseCase) Run(id int32) (*entities.Tutorado, error) {
	return uc.repo.GetById(id)
}
