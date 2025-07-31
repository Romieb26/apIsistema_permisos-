package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
)

type GetAllDocentesUseCase struct {
	repo domain.IDocente
}

func NewGetAllDocentesUseCase(repo domain.IDocente) *GetAllDocentesUseCase {
	return &GetAllDocentesUseCase{repo: repo}
}

func (uc *GetAllDocentesUseCase) Run() ([]entities.Docente, error) {
	return uc.repo.GetAll()
}
