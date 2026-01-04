package main

import (
	"fmt"

	payments "github.com/Oracl0s/Education.git/cmd/delaem"
	"github.com/Oracl0s/Education.git/cmd/delaem/methods"
)

func main() {
	method := methods.NewBank()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Бургер", 5)
	paymentModule.Pay("Телефон", 500)
	paymentModule.Pay("Игра", 20)

	allInfo := paymentModule.AllInfo()
	fmt.Println(allInfo)

}
