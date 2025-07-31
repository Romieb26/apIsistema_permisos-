package application

import (
	repositories "github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain/entities"
)

type SaveTutoradoUseCase struct {
	repo repositories.ITutorado
}

func NewSaveTutoradoUseCase(repo repositories.ITutorado) *SaveTutoradoUseCase {
	return &SaveTutoradoUseCase{repo: repo}
}

func (uc *SaveTutoradoUseCase) Run(tutorado *entities.Tutorado) (*entities.Tutorado, error) {
	err := uc.repo.Save(tutorado)
	if err != nil {
		return nil, err
	}
	return tutorado, nil
}
