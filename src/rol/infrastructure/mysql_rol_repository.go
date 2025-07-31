package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Romieb26/ApIsistema_permisos/src/core"
	repositories "github.com/Romieb26/ApIsistema_permisos/src/rol/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/rol/domain/entities"
)

type MySQLRolRepository struct {
	conn *sql.DB
}

func NewMySQLRolRepository() repositories.IRol {
	conn := core.GetDB()
	return &MySQLRolRepository{conn: conn}
}

func (mysql *MySQLRolRepository) Save(rol *entities.Rol) error {
	query := `
		INSERT INTO rol (title, description)
		VALUES (?, ?)
	`
	result, err := mysql.conn.Exec(query, rol.Title, rol.Description)
	if err != nil {
		log.Println("Error al guardar el rol:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	rol.ID = int32(id)
	return nil
}

func (mysql *MySQLRolRepository) Update(rol *entities.Rol) error {
	query := `
		UPDATE rol
		SET title = ?, description = ?
		WHERE id_rol = ?
	`
	result, err := mysql.conn.Exec(query, rol.Title, rol.Description, rol.ID)
	if err != nil {
		log.Println("Error al actualizar el rol:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rol con ID %d no encontrado", rol.ID)
	}

	return nil
}

func (mysql *MySQLRolRepository) Delete(id int32) error {
	query := "DELETE FROM rol WHERE id_rol = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el rol:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rol con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MySQLRolRepository) GetById(id int32) (*entities.Rol, error) {
	query := `
		SELECT id_rol, title, description
		FROM rol
		WHERE id_rol = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var rol entities.Rol
	err := row.Scan(&rol.ID, &rol.Title, &rol.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("rol con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el rol por ID:", err)
		return nil, err
	}

	return &rol, nil
}

func (mysql *MySQLRolRepository) GetAll() ([]entities.Rol, error) {
	query := `
		SELECT id_rol, title, description
		FROM rol
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los roles:", err)
		return nil, err
	}
	defer rows.Close()

	var roles []entities.Rol
	for rows.Next() {
		var rol entities.Rol
		err := rows.Scan(&rol.ID, &rol.Title, &rol.Description)
		if err != nil {
			log.Println("Error al escanear el rol:", err)
			return nil, err
		}
		roles = append(roles, rol)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return roles, nil
}
