package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (c PayPal) Pay(usd int) int {
	fmt.Println("Оплата PayPal!")
	fmt.Println("Размер оплаты:", usd, "USDT")

	return rand.Int()
}
func (c PayPal) Cancel(id int) {
	fmt.Println("Отмена PayPal-операции! ID:", id)
}
