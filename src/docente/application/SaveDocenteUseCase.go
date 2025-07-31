package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
)

type SaveDocenteUseCase struct {
	repo domain.IDocente
}

func NewSaveDocenteUseCase(repo domain.IDocente) *SaveDocenteUseCase {
	return &SaveDocenteUseCase{repo: repo}
}

func (uc *SaveDocenteUseCase) Run(docente *entities.Docente) (*entities.Docente, error) {
	err := uc.repo.Save(docente)
	if err != nil {
		return nil, err
	}
	return docente, nil
}
