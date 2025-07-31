package domain

import "github.com/Romieb26/ApIsistema_permisos/src/permission/domain/entities"

type IPermission interface {
	Save(permission *entities.Permission) error
	Update(permission *entities.Permission) error
	Delete(id int32) error
	GetById(id int32) (*entities.Permission, error)
	GetAll() ([]entities.Permission, error)
}
