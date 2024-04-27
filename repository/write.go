package repository

import (
	"fmt"
	"test/domain"
)

func (h *PostgresHandler) AddTODO(todo domain.TODO) (int64, error) {
	var id int64

	db, err := h.Open()
	if err != nil {
		var id int64
		return id, fmt.Errorf("method db Open error: %s", err.Error())
	}

	defer db.Close()

	sql := `
		INSERT INTO 
			todo
			(	
				,desc
				,do
			) VALUES (

				,$2
				,$3
			)
	`

	result, err := db.Exec(sql, todo.Description, todo.Do)
	if err != nil {
		return id, fmt.Errorf("method db Exec error: %s", err.Error())
	}

	id, err = result.LastInsertId()
	if err != nil {
		return id, fmt.Errorf("method db LastInsertId error: %s", err.Error())
	}

	return id, nil
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
			desc = $2
			,do = $3
		WHERE
			id = $1
	`

	result, err := db.Exec(sql, todo.Id, todo.Description, todo.Do)
	if err != nil {
		return id, fmt.Errorf("method db Exec error: %s", err.Error())
	}

	id, err = result.LastInsertId()
	if err != nil {
		return id, fmt.Errorf("method db LastInsertId error: %s", err.Error())
	}

	return id, nil
}
