package domain

import "github.com/Romieb26/ApIsistema_permisos/src/user/domain/entities"

type IUser interface {
	// Obtener todos los usuarios
	GetAll() ([]entities.User, error)
	
	// Obtener usuario por ID
	GetById(id int32) (*entities.User, error)
	
	// Crear un nuevo usuario
	Save(user *entities.User) error
	
	// Actualizar un usuario existente
	Update(user *entities.User) error
	
	// Eliminar un usuario por ID
	Delete(id int32) error
	
	// Buscar usuario por username y password (login)
	Login(username, password string) (*entities.User, error)
}
