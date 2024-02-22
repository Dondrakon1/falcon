// В файле internal/service/service.go

package service

import "fmt"

type CodeStorage interface {
	AddCode(code string) error
}

// CodeStorageService представляет сервис для работы с кодами.
type CodeStorageService struct {
	storage CodeStorage
}

// NewCodeService создает новый экземпляр сервиса для работы с кодами.
func NewCodeService(storage CodeStorage) *CodeStorageService {
	return &CodeStorageService{storage: storage}
}

// AddCode добавляет новый код.
func (s *CodeStorageService) AddCode(code string) error {
	fmt.Println("AddCode")
	return s.storage.AddCode(code)
}
