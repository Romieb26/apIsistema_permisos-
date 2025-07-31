package application

import (
	repositories "github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain/entities"
)

type UpdateTutoradoUseCase struct {
	repo repositories.ITutorado
}

func NewUpdateTutoradoUseCase(repo repositories.ITutorado) *UpdateTutoradoUseCase {
	return &UpdateTutoradoUseCase{repo: repo}
}

func (uc *UpdateTutoradoUseCase) Run(tutorado *entities.Tutorado) (*entities.Tutorado, error) {
	err := uc.repo.Update(tutorado)
	if err != nil {
		return nil, err
	}
	return tutorado, nil
}
