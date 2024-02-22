package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	Db *sql.DB
}

func New(path string) (*Storage, error) {
	op := "sqlite.New"

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS codes(
  	  id INTEGER PRIMARY KEY ,
  	  payload TEXT NOT NULL UNIQUE,
  	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_payload ON codes(payload);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	_, err = stmt.Exec()

	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &Storage{Db: db}, nil
}
