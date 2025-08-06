package domain

import "github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"

type IDocente interface {
	Save(docente *entities.Docente) error
	Update(docente *entities.Docente) error
	Delete(id int32) error
	GetById(id int32) (*entities.Docente, error)
	GetAll() ([]entities.Docente, error)
}
