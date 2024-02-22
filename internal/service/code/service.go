// В файле internal/service/service.go

package code

import (
	"encoding/json"
	"falcon/internal/storage/sqlite"
	"fmt"
)

type Storage interface {
	AddCode(code string) error
	GetCodeByPayload(payload string) (*sqlite.Code, error)
}

// CodeStorageService представляет сервис для работы с кодами.
type StorageService struct {
	storage Storage
}

// NewCodeService создает новый экземпляр сервиса для работы с кодами.
func NewCodeService(storage Storage) *StorageService {
	return &StorageService{storage: storage}
}

// AddCode добавляет новый код.
func (s *StorageService) AddCode(code string) error {
	fmt.Printf("AddCode: %s\n", code)

	return s.storage.AddCode(code)
}

// GetCodeByPayload возвращает код по его payload.
func (s *StorageService) GetCodeByPayload(payload string) ([]byte, error) {
	fmt.Println("GetCodeByPayload")

	code, err := s.storage.GetCodeByPayload(payload)
	if err != nil {
		return nil, err
	}

	codePayload, err := json.Marshal(code)
	if err != nil {
		return nil, err
	}
	return codePayload, nil

}
