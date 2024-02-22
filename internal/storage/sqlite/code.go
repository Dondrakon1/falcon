package sqlite

import (
	"database/sql"
	"errors"
	"falcon/internal/storage"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"time"
)

// Code представляет собой структуру для хранения информации о коде в базе данных.
type Code struct {
	ID        int64     `json:"id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Storage) AddCode(code string) error {
	op := "sqlite.Storage.AddCode"

	// Подготовка SQL-запроса для вставки кода в таблицу
	stmt, err := s.Db.Prepare("INSERT INTO codes(payload) VALUES(?)")
	if err != nil {
		return fmt.Errorf("%s: preparing statement: %s", op, err)
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для добавления кода
	_, err = stmt.Exec(code)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return fmt.Errorf("%s: %s ", op, storage.ErrCodeExists)
		}
		return fmt.Errorf("%s: executing statement: %s", op, err)
	}

	return nil
}

func (s *Storage) GetCodeByPayload(payload string) (*Code, error) {
	op := "sqlite.Storage.GetCodeByPayload"

	var code Code

	stmt, err := s.Db.Prepare("SELECT id, payload, created_at FROM codes WHERE payload = ?")
	if err != nil {
		return nil, fmt.Errorf("%s: preparing statement: %s", op, err)
	}

	if err := stmt.QueryRow(payload).Scan(&code.ID, &code.Payload, &code.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", storage.ErrCodeNotFound, err)
		}
		return nil, fmt.Errorf("%s: executing statement: %s", op, err)
	}

	defer stmt.Close()

	return &code, nil

}
