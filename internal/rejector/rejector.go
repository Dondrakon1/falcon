package rejector

import (
	"github.com/Dondrakon1/falcon/internal/conveyor"
)

func Start(productQueue chan *conveyor.Product) {
	for {
		select {
		case product := <-productQueue:
			// Отбраковать или принять продукт, учитывая считанный штрих-код
			if product != nil && product.Code != "" {
				// Дополнительная логика отбраковывателя
			}
		}
	}
}
