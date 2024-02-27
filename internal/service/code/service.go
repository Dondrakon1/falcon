// В файле internal/service/service.go

package code

import (
	"encoding/json"
	"falcon/internal/storage/models"
	"fmt"
	"log/slog"
	"os"
)

const codeLength = 30

type Storage interface {
	AddCode(code string) error
	GetCodeByPayload(payload string) (*models.Code, error)
	GetByOrderID(orderID int64) (*models.Codes, error)
}

// CodeStorageService представляет сервис для работы с кодами.
type StorageService struct {
	storage   Storage
	Log       *slog.Logger
	codeCache map[string]struct{}
}

// NewCodeService создает новый экземпляр сервиса для работы с кодами.
func NewCodeService(storage Storage) *StorageService {
	return &StorageService{
		storage:   storage,
		codeCache: make(map[string]struct{}),
	}
}

// AddCode добавляет новый код в хранилище проверяя на уникальность в течении одного заказа"
func (s *StorageService) AddCode(newCode string) error {
	if len(newCode) != codeLength {
		return fmt.Errorf("invalid code length: %d", len(newCode))
	}
	if _, exists := s.codeCache[newCode]; exists {
		return fmt.Errorf("code already exists: %s", newCode)

	}

	if err := s.storage.AddCode(newCode); err != nil {
		return err
	}

	s.codeCache[newCode] = struct{}{}
	s.Log.Debug("code added", slog.String("code", newCode))
	return nil
}

// GetCodeByPayload retrieves the code by its payload.
func (s *StorageService) GetCodeByPayload(payload string) ([]byte, error) {
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

func (s *StorageService) GetByOrderID(orderID int64) (*models.Codes, error) {
	if orderID == 0 {
		return nil, fmt.Errorf("invalid order ID: %d", orderID)
	}
	codes, err := s.storage.GetByOrderID(orderID)
	if err != nil {
		return nil, err
	}
	if len(codes.Codes) == 0 {
		return nil, fmt.Errorf("order ID not found: %d", orderID)
	}

	return codes, nil

}

// Функция для сохранения payload кодов в файл
func (s *StorageService) SavePayloadsToFile(orderID int64, filename string) error {
	codes, err := s.GetByOrderID(orderID)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	defer file.Close()

	for _, code := range codes.Codes {
		_, err := file.WriteString(code.Payload + "\n")
		if err != nil {
			return fmt.Errorf("unable to write to file: %v", err)
		}
	}

	return nil
}
