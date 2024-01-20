package conveyor

import (
	"falcon/internal/rejector"
	"falcon/internal/sensor"
	"sync"
	"time"
)

type Product struct {
	Code       string
	Expiration *time.Timer
}

type ConveyorSystem struct {
	Camera       chan *Product
	Sensor       chan *Product
	Rejector     chan *Product
	ProductQueue chan *Product
	mu           sync.Mutex
}

func NewConveyorSystem() *ConveyorSystem {
	return &ConveyorSystem{
		Camera:       make(chan *Product),
		Sensor:       make(chan *Product),
		Rejector:     make(chan *Product),
		ProductQueue: make(chan *Product, 10),
	}
}

func (cs *ConveyorSystem) Start() {
	go camera.Start(cs.ProductQueue)
	go sensor.Start(cs.ProductQueue)
	go rejector.Start(cs.ProductQueue)
}

func (cs *ConveyorSystem) Stop() {
	close(cs.ProductQueue)
}
