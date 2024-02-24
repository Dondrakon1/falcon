package code

import "time"

// Code представляет собой структуру для хранения информации о коде в базе данных.
type Code struct {
	ID        int64     `json:"id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}
