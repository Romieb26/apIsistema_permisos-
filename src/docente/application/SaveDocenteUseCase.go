package application

import (
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
)
type SaveDocenteUseCase struct {
	Repository domain.IDocente
}

func NewSaveDocenteUseCase(repo domain.IDocente) *SaveDocenteUseCase {
	return &SaveDocenteUseCase{
		Repository: repo,
	}
}

func (uc *SaveDocenteUseCase) Execute(docente *entities.Docente) error {
	// Verifica que el correo sea válido
	if err := docente.ValidarCorreo(); err != nil {
		return err
	}

	// Llama al repositorio si el correo es válido
	return uc.Repository.Save(docente)
}

