package sqlite

import (
	"database/sql"

	"github.com/AmarjitKaranSharma/golang-student-api/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	age INTEGER NOT NULL)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil
}
