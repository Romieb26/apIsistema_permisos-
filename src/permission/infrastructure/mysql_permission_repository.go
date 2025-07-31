package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Romieb26/ApIsistema_permisos/src/core"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/permission/domain/entities"
)

type MySQLPermissionRepository struct {
	conn *sql.DB
}

func NewMySQLPermissionRepository() domain.IPermission {
	conn := core.GetDB()
	return &MySQLPermissionRepository{conn: conn}
}

func (repo *MySQLPermissionRepository) Save(permission *entities.Permission) error {
	query := `
		INSERT INTO permission (id_tutorado_fk, id_docente_fk, evidencia, date, motivo, estatus)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := repo.conn.Exec(
		query,
		permission.TutoradoID,
		permission.DocenteID,
		permission.Evidencia,
		permission.Date,
		permission.Motivo,
		permission.Estatus,
	)
	if err != nil {
		log.Println("Error inserting permission:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return err
	}
	permission.ID = int32(id)
	return nil
}

func (repo *MySQLPermissionRepository) Update(permission *entities.Permission) error {
	query := `
		UPDATE permission
		SET id_tutorado_fk = ?, id_docente_fk = ?, evidencia = ?, date = ?, motivo = ?, estatus = ?
		WHERE id_permission = ?
	`
	result, err := repo.conn.Exec(
		query,
		permission.TutoradoID,
		permission.DocenteID,
		permission.Evidencia,
		permission.Date,
		permission.Motivo,
		permission.Estatus,
		permission.ID,
	)
	if err != nil {
		log.Println("Error updating permission:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("permission with ID %d not found", permission.ID)
	}

	return nil
}

func (repo *MySQLPermissionRepository) Delete(id int32) error {
	query := "DELETE FROM permission WHERE id_permission = ?"
	result, err := repo.conn.Exec(query, id)
	if err != nil {
		log.Println("Error deleting permission:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("permission with ID %d not found", id)
	}

	return nil
}

func (repo *MySQLPermissionRepository) GetById(id int32) (*entities.Permission, error) {
	query := `
		SELECT id_permission, id_tutorado_fk, id_docente_fk, evidencia, date, motivo, estatus
		FROM permission
		WHERE id_permission = ?
	`
	row := repo.conn.QueryRow(query, id)

	var permission entities.Permission
	err := row.Scan(
		&permission.ID,
		&permission.TutoradoID,
		&permission.DocenteID,
		&permission.Evidencia,
		&permission.Date,
		&permission.Motivo,
		&permission.Estatus,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("permission with ID %d not found", id)
		}
		log.Println("Error scanning permission:", err)
		return nil, err
	}

	return &permission, nil
}

func (repo *MySQLPermissionRepository) GetAll() ([]entities.Permission, error) {
	query := `
		SELECT id_permission, id_tutorado_fk, id_docente_fk, evidencia, date, motivo, estatus
		FROM permission
	`
	rows, err := repo.conn.Query(query)
	if err != nil {
		log.Println("Error querying permissions:", err)
		return nil, err
	}
	defer rows.Close()

	var permissions []entities.Permission
	for rows.Next() {
		var permission entities.Permission
		err := rows.Scan(
			&permission.ID,
			&permission.TutoradoID,
			&permission.DocenteID,
			&permission.Evidencia,
			&permission.Date,
			&permission.Motivo,
			&permission.Estatus,
		)
		if err != nil {
			log.Println("Error scanning permission:", err)
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error after iterating rows:", err)
		return nil, err
	}

	return permissions, nil
}
