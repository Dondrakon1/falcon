package main

import (
	"github.com/Dondrakon1/falcon/internal/conveyor"
)

func main() {
	conveyorSystem := conveyor.NewConveyorSystem()
	defer conveyorSystem.Stop()

	conveyorSystem.Start()

	// Пример добавления продукта на конвейер
	product := &conveyor.Product{Code: "123456"}
	conveyorSystem.ProductQueue <- product

	// Ждем завершения программы
	select {}
}
