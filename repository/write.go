package repository

import (
	"fmt"
	"test/domain"
)

func (h *PostgresHandler) AddTODO(todo domain.TODO) error {

	db, err := h.Open()
	if err != nil {
		return fmt.Errorf("method db Open error: %s", err.Error())
	}

	defer db.Close()

	sql := `
		INSERT INTO todo
		(	
			description
			,status
		) VALUES (
			$1
			,$2
		)
	`

	_, err = db.Exec(sql, todo.Description, todo.Status)
	if err != nil {
		return fmt.Errorf("method db Exec error: %s", err.Error())
	}

	return nil
}

func (h *PostgresHandler) UpdateTODO(todo domain.TODO) (int64, error) {
	var id int64

	db, err := h.Open()
	if err != nil {
		var id int64
		return id, fmt.Errorf("method db Open error: %s", err.Error())
	}

	defer db.Close()

	sql := `
		UPDATE
			todo
		SET
			description = $2
			,status = $3
		WHERE
			id = $1
	`

	if _, err = db.Exec(sql, todo.Id, todo.Description, todo.Status); err != nil {
		return id, fmt.Errorf("method db Exec error: %s", err.Error())
	}

	return id, nil
}

func (h *PostgresHandler) DeleteTODO(id uint64) error {

	db, err := h.Open()
	if err != nil {
		return fmt.Errorf("method db Open error: %s", err.Error())
	}

	defer db.Close()

	sql := `
		DELETE FROM
			todo
		WHERE
			id = $1
	`

	_, err = db.Exec(sql, id)
	if err != nil {
		return fmt.Errorf("method db Exec error: %s", err.Error())
	}

	return nil
}
