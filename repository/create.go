package repository

import "fmt"

func (h *PostgresHandler) CreateTable() error {

	db, err := h.Open()
	if err != nil {
		return fmt.Errorf("method db Open error: %s", err.Error())
	}

	defer db.Close()

	sql := `
		CREATE TABLE IF NOT EXISTS todo (
			id    BIGINT PRIMARY KEY,
			desc  TEXT NOT NULL,
			do boolean NOT NULL,
		);
	`

	_, err = db.Exec(sql)
	if err != nil {
		return fmt.Errorf("method db Exec error: %s", err.Error())
	}

	return nil
}
