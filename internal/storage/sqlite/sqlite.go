package sqlite

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mukul-work/golang-students-api/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER )`)

	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	slog.Info(
		"students table created",
		slog.Int64("rows_affected", rows),
	)

	return &Sqlite{
		Db: db,
	}, nil
}
