package camera

import (
	"fmt"
	"github.com/Dondrakon1/falcon/internal/conveyor"
)

func Start(productQueue chan *conveyor.Product) {
	for {
		select {
		case product := <-productQueue:
			fmt.Printf("Камера считала штрих-код: %s\n", product.Code)
			// Дополнительная логика обработки камеры
		}
	}
}
