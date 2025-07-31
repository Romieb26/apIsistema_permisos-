package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
)

type GetDocenteByIdUseCase struct {
	repo domain.IDocente
}

func NewGetDocenteByIdUseCase(repo domain.IDocente) *GetDocenteByIdUseCase {
	return &GetDocenteByIdUseCase{repo: repo}
}

func (uc *GetDocenteByIdUseCase) Run(id int32) (*entities.Docente, error) {
	return uc.repo.GetById(id)
}
