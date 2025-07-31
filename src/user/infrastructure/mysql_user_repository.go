package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Romieb26/ApIsistema_permisos/src/core"
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/user/domain/entities"
)

type MySQLUserRepository struct {
	conn *sql.DB
}

func NewMySQLUserRepository() domain.IUser {
	conn := core.GetDB()
	return &MySQLUserRepository{conn: conn}
}

func (r *MySQLUserRepository) GetAll() ([]entities.User, error) {
	query := `
		SELECT id_user, name, lastname, email, password, username, id_rol_fk
		FROM user
	`
	rows, err := r.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.Username,
			&user.IdRolFK,
		)
		if err != nil {
			log.Println("Error al escanear usuario:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return users, nil
}

func (r *MySQLUserRepository) GetById(id int32) (*entities.User, error) {
	query := `
		SELECT id_user, name, lastname, email, password, username, id_rol_fk
		FROM user
		WHERE id_user = ?
	`
	row := r.conn.QueryRow(query, id)

	var user entities.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.IdRolFK,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		log.Println("Error al buscar usuario por ID:", err)
		return nil, err
	}

	return &user, nil
}

func (r *MySQLUserRepository) Save(user *entities.User) error {
	query := `
		INSERT INTO user (name, lastname, email, password, username, id_rol_fk)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := r.conn.Exec(
		query,
		user.Name,
		user.LastName,
		user.Email,
		user.Password,
		user.Username,
		user.IdRolFK,
	)
	if err != nil {
		log.Println("Error al guardar usuario:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}

	user.ID = int32(id)
	return nil
}

func (r *MySQLUserRepository) Update(user *entities.User) error {
	query := `
		UPDATE user
		SET name = ?, lastname = ?, email = ?, password = ?, username = ?, id_rol_fk = ?
		WHERE id_user = ?
	`
	result, err := r.conn.Exec(
		query,
		user.Name,
		user.LastName,
		user.Email,
		user.Password,
		user.Username,
		user.IdRolFK,
		user.ID,
	)
	if err != nil {
		log.Println("Error al actualizar usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", user.ID)
	}

	return nil
}

func (r *MySQLUserRepository) Delete(id int32) error {
	query := `DELETE FROM user WHERE id_user = ?`
	result, err := r.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", id)
	}

	return nil
}

func (r *MySQLUserRepository) Login(username, password string) (*entities.User, error) {
	query := `
		SELECT id_user, name, lastname, email, password, username, id_rol_fk
		FROM user
		WHERE username = ? AND password = ?
	`
	row := r.conn.QueryRow(query, username, password)

	var user entities.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.IdRolFK,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("credenciales inválidas")
		}
		log.Println("Error al hacer login:", err)
		return nil, err
	}

	return &user, nil
}
