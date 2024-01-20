package sensor

import (
	"fmt"
	"github.com/Dondrakon1/falcon/internal/conveyor"
)

func Start(productQueue chan *conveyor.Product) {
	for {
		select {
		case product := <-productQueue:
			fmt.Println("Датчик активирован.")
			// Принять решение о браке или принятии продукта, учитывая считанный штрих-код
			if product != nil && product.Code != "" {
				// Дополнительная логика датчика
			}
		}
	}
}
