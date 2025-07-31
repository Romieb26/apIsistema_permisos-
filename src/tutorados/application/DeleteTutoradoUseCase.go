package application

import "github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain"

type DeleteTutoradoUseCase struct {
	repo domain.ITutorado
}

func NewDeleteTutoradoUseCase(repo domain.ITutorado) *DeleteTutoradoUseCase {
	return &DeleteTutoradoUseCase{repo: repo}
}

func (uc *DeleteTutoradoUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}
