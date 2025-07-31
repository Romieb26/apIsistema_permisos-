package domain

import "github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain/entities"

type ITutorado interface {
	Save(tutorado *entities.Tutorado) error
	Update(tutorado *entities.Tutorado) error
	Delete(id int32) error
	GetById(id int32) (*entities.Tutorado, error)
	GetAll() ([]entities.Tutorado, error)
}
