package camera

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"strings"
	"time"
)

type codeService interface {
	AddCode(code string) error
	GetCodeByPayload(payload string) ([]byte, error)
}

type Camera struct {
	Log         *slog.Logger
	client      net.Conn
	service     codeService
	address     string
	isConnected bool
}

func NewCamera(address string, service codeService) (*Camera, error) {
	// ...
	return &Camera{address: address, service: service}, nil
}

func (c *Camera) connect() error {
	conn, err := net.Dial("tcp4", c.address)
	if err != nil {
		return err
	}
	c.client = conn
	c.isConnected = true
	return nil
}

func (c *Camera) StartListening() {
	for {
		if !c.isConnected {
			err := c.connect()
			if err != nil {
				c.Log.Info("failed to connect to camera", slog.String("error", err.Error()))
				time.Sleep(10 * time.Second) // wait before retrying
				continue
			}
		}

		reader := bufio.NewReader(c.client)
		data, err := reader.ReadString('\r')
		if err != nil {
			c.Log.Info("failed to read from camera", slog.String("error", err.Error()))
			c.isConnected = false
			continue
		}

		result := strings.Trim(data, "\r")
		if err := c.service.AddCode(result); err != nil {
			fmt.Printf("failed to add code to storage: %v\n", err)
		}
	}
}

func (c *Camera) Close() {
	if c.client != nil {
		c.client.Close()
	}
	c.isConnected = false
}
