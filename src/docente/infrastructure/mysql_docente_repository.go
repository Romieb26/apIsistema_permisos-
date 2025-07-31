package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Romieb26/ApIsistema_permisos/src/core"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
)

type MySQLDocenteRepository struct {
	conn *sql.DB
}

func NewMySQLDocenteRepository() domain.IDocente {
	conn := core.GetDB()
	return &MySQLDocenteRepository{conn: conn}
}

func (repo *MySQLDocenteRepository) Save(docente *entities.Docente) error {
	query := `
		INSERT INTO docente (name, lastname, email)
		VALUES (?, ?, ?)
	`
	result, err := repo.conn.Exec(
		query,
		docente.Name,
		docente.LastName,
		docente.Email,
	)
	if err != nil {
		log.Println("Error al guardar docente:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID:", err)
		return err
	}

	docente.ID = int32(id)
	return nil
}

func (repo *MySQLDocenteRepository) Update(docente *entities.Docente) error {
	query := `
		UPDATE docente
		SET name = ?, lastname = ?, email = ?
		WHERE id_docente = ?
	`
	result, err := repo.conn.Exec(
		query,
		docente.Name,
		docente.LastName,
		docente.Email,
		docente.ID,
	)
	if err != nil {
		log.Println("Error al actualizar docente:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al verificar filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("docente con ID %d no encontrado", docente.ID)
	}

	return nil
}

func (repo *MySQLDocenteRepository) Delete(id int32) error {
	query := `DELETE FROM docente WHERE id_docente = ?`
	result, err := repo.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar docente:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al verificar filas eliminadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("docente con ID %d no encontrado", id)
	}

	return nil
}

func (repo *MySQLDocenteRepository) GetById(id int32) (*entities.Docente, error) {
	query := `
		SELECT id_docente, name, lastname, email
		FROM docente
		WHERE id_docente = ?
	`

	row := repo.conn.QueryRow(query, id)

	var docente entities.Docente
	err := row.Scan(
		&docente.ID,
		&docente.Name,
		&docente.LastName,
		&docente.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("docente con ID %d no encontrado", id)
		}
		log.Println("Error al obtener docente por ID:", err)
		return nil, err
	}

	return &docente, nil
}

func (repo *MySQLDocenteRepository) GetAll() ([]entities.Docente, error) {
	query := `
		SELECT id_docente, name, lastname, email
		FROM docente
	`

	rows, err := repo.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los docentes:", err)
		return nil, err
	}
	defer rows.Close()

	var docentes []entities.Docente
	for rows.Next() {
		var docente entities.Docente
		err := rows.Scan(
			&docente.ID,
			&docente.Name,
			&docente.LastName,
			&docente.Email,
		)
		if err != nil {
			log.Println("Error al escanear docente:", err)
			return nil, err
		}
		docentes = append(docentes, docente)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error al iterar resultados de docentes:", err)
		return nil, err
	}

	return docentes, nil
}
