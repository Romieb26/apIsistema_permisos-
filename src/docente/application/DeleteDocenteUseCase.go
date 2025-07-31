package application

import "github.com/Romieb26/ApIsistema_permisos/src/docente/domain"

type DeleteDocenteUseCase struct {
	repo domain.IDocente
}

func NewDeleteDocenteUseCase(repo domain.IDocente) *DeleteDocenteUseCase {
	return &DeleteDocenteUseCase{repo: repo}
}

func (uc *DeleteDocenteUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}
