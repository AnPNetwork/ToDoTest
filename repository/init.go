package repository

import (
	"database/sql"
	"fmt"
	"test/domain"

	_ "github.com/lib/pq"
)

type IDbHandler interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Close() error
}

type PostgresHandler struct {
	connection string
}

func NewConnectionDB(settings domain.DBSettings) (*PostgresHandler, error) {

	connPG := fmt.Sprintf(
		`
		user=postgres 
		host=%s 
		port=%d 
		password=%s 
		dbname=%s 
		sslmode=%s
		`,
		settings.Host,
		settings.Port,
		settings.Password,
		settings.Name,
		settings.SSL,
	)

	postgresHandler := &PostgresHandler{
		connection: connPG,
	}

	db, err := postgresHandler.Open()
	if err != nil {
		return nil, err
	}

	db.Close()

	return postgresHandler, nil
}

func (psql *PostgresHandler) Open() (IDbHandler, error) {
	db, err := sql.Open("postgres", psql.connection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	return db, nil
}
