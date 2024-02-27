package models

import (
	"time"
)

// Code представляет собой структуру для хранения информации о коде в базе данных.
type Code struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"order_id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

type Codes struct {
	OrderID int64
	Codes   []Code
}
