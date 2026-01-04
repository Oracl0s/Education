package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Оплата банковской картой!")
	fmt.Println("Размер оплаты:", usd, "USDT")

	return rand.Int()
}
func (c Bank) Cancel(id int) {
	fmt.Println("Отмена операции по карте! ID:", id)
}
