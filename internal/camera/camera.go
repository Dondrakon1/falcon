package camera

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type CodeService interface {
	AddCode(code string) error
	GetCodeByPayload(payload string) ([]byte, error)
}

type Camera struct {
	client  net.Conn
	service CodeService
}

func NewCamera(address string, service CodeService) (*Camera, error) {
	op := "camera.NewCamera"

	conn, err := net.Dial("tcp4", address)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}
	return &Camera{client: conn, service: service}, nil
}

func (c *Camera) StartListening() {
	reader := bufio.NewReader(c.client)
	for {
		data, err := reader.ReadString('\r')
		if err != nil {
			log.Println("Ошибка чтения данных:", err)
			break
		}
		result := strings.Trim(data, "\r")

		if err := c.service.AddCode(result); err != nil {
			fmt.Printf("failed to add code to storage: %v\n", err)
		}
	}
}

func (c *Camera) Close() {
	c.client.Close()
}
