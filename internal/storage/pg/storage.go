package pg

import (
	"context"
	"errors"
	"falcon/internal/storage/models"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Storage struct {
	Db *pgxpool.Pool
}

func сonnectDB(path string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), path)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Connected!")
	return pool
}

func New(path string) (*Storage, error) {
	pool := сonnectDB(path)
	return &Storage{Db: pool}, nil
}

func (s *Storage) AddCode(code string) error {
	const query = `INSERT INTO codes (order_id, payload, created_at) VALUES ($1,$2, NOW())`
	_, err := s.Db.Exec(context.Background(), query, 3, code)
	if err != nil {
		return fmt.Errorf("failed to add code: %v", err)
	}
	return nil
}

func (s *Storage) GetCodeByPayload(payload string) (*models.Code, error) {
	var code models.Code
	const query = `SELECT id, payload, created_at FROM codes WHERE payload = $1`
	err := s.Db.QueryRow(context.Background(), query, payload).Scan(&code.ID, &code.Payload, &code.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("code not found")
		}
		return nil, fmt.Errorf("failed to get code: %v", err)
	}
	return &code, nil
}

func (s *Storage) GetByOrderID(orderID int64) (*models.Codes, error) {
	var codes models.Codes

	const query = `SELECT id, payload, created_at FROM codes WHERE order_id = $1`
	rows, err := s.Db.Query(context.Background(), query, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get codes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var code models.Code
		code.OrderID = orderID
		if err := rows.Scan(&code.ID, &code.Payload, &code.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan code: %v", err)
		}

		codes.Codes = append(codes.Codes, code)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %v", err)
	}

	return &codes, nil
}
