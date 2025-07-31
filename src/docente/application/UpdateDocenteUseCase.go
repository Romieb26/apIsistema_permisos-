package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
)

type UpdateDocenteUseCase struct {
	repo domain.IDocente
}

func NewUpdateDocenteUseCase(repo domain.IDocente) *UpdateDocenteUseCase {
	return &UpdateDocenteUseCase{repo: repo}
}

func (uc *UpdateDocenteUseCase) Run(docente *entities.Docente) (*entities.Docente, error) {
	err := uc.repo.Update(docente)
	if err != nil {
		return nil, err
	}
	return docente, nil
}
