package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Romieb26/ApIsistema_permisos/src/core"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/tutorados/domain/entities"
)

type MySQLTutoradoRepository struct {
	conn *sql.DB
}

func NewMySQLTutoradoRepository() domain.ITutorado {
	conn := core.GetDB()
	return &MySQLTutoradoRepository{conn: conn}
}

func (repo *MySQLTutoradoRepository) Save(tutorado *entities.Tutorado) error {
	query := `
		INSERT INTO tutorados (id_user_fk, name, lastname, matricula, email, estatus)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := repo.conn.Exec(
		query,
		tutorado.IdUserFK,
		tutorado.Name,
		tutorado.LastName,
		tutorado.Matricula,
		tutorado.Email,
		tutorado.Estatus,
	)
	if err != nil {
		log.Println("Error al guardar tutorado:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	tutorado.ID = int32(id)
	return nil
}

func (repo *MySQLTutoradoRepository) Update(tutorado *entities.Tutorado) error {
	query := `
		UPDATE tutorados
		SET id_user_fk = ?, name = ?, lastname = ?, matricula = ?, email = ?, estatus = ?
		WHERE id_tutorado = ?
	`
	result, err := repo.conn.Exec(
		query,
		tutorado.IdUserFK,
		tutorado.Name,
		tutorado.LastName,
		tutorado.Matricula,
		tutorado.Email,
		tutorado.Estatus,
		tutorado.ID,
	)
	if err != nil {
		log.Println("Error al actualizar tutorado:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tutorado con ID %d no encontrado", tutorado.ID)
	}

	return nil
}

func (repo *MySQLTutoradoRepository) Delete(id int32) error {
	query := "DELETE FROM tutorados WHERE id_tutorado = ?"
	result, err := repo.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar tutorado:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tutorado con ID %d no encontrado", id)
	}

	return nil
}

func (repo *MySQLTutoradoRepository) GetById(id int32) (*entities.Tutorado, error) {
	query := `
		SELECT id_tutorado, id_user_fk, name, lastname, matricula, email, estatus
		FROM tutorados
		WHERE id_tutorado = ?
	`
	row := repo.conn.QueryRow(query, id)

	var tutorado entities.Tutorado
	err := row.Scan(
		&tutorado.ID,
		&tutorado.IdUserFK,
		&tutorado.Name,
		&tutorado.LastName,
		&tutorado.Matricula,
		&tutorado.Email,
		&tutorado.Estatus,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tutorado con ID %d no encontrado", id)
		}
		log.Println("Error al buscar tutorado por ID:", err)
		return nil, err
	}

	return &tutorado, nil
}

func (repo *MySQLTutoradoRepository) GetAll() ([]entities.Tutorado, error) {
	query := `
		SELECT id_tutorado, id_user_fk, name, lastname, matricula, email, estatus
		FROM tutorados
	`
	rows, err := repo.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los tutorados:", err)
		return nil, err
	}
	defer rows.Close()

	var tutorados []entities.Tutorado
	for rows.Next() {
		var tutorado entities.Tutorado
		err := rows.Scan(
			&tutorado.ID,
			&tutorado.IdUserFK,
			&tutorado.Name,
			&tutorado.LastName,
			&tutorado.Matricula,
			&tutorado.Email,
			&tutorado.Estatus,
		)
		if err != nil {
			log.Println("Error al escanear tutorado:", err)
			return nil, err
		}
		tutorados = append(tutorados, tutorado)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return tutorados, nil
}
