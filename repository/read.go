package repository

import (
	"fmt"
	"test/domain"
)

func (h *PostgresHandler) GetTODOList() ([]domain.TODO, error) {
	db, err := h.Open()
	if err != nil {
		return nil, fmt.Errorf("method db Open error: %s", err.Error())
	}

	defer db.Close()

	sql := `
		SELECT
			id
			,description
			,status
		FROM
			todo
		ORDER BY description ASC
	`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("method db Query error: %s", err.Error())
	}
	defer rows.Close()

	result := make([]domain.TODO, 0)

	for rows.Next() {
		el := domain.TODO{}
		err := rows.Scan(&el.Id, &el.Description, &el.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result = append(result, el)
	}

	return result, nil
}

func (h *PostgresHandler) ExistsTODO(id int) (bool, error) {
	db, err := h.Open()
	if err != nil {
		return false, fmt.Errorf("method db Open error: %s", err.Error())
	}

	defer db.Close()

	sql := `
		SELECT EXITS (
			SELECT
				*
			FROM
				todo
			WHERE 
				id = $1
		)
	`

	row := db.QueryRow(sql, id)
	if err != nil {
		return false, fmt.Errorf("method db Query error: %s", err.Error())
	}

	var isExists bool

	err = row.Scan(&isExists)
	if err != nil {
		return false, err
	}

	return isExists, nil
}
