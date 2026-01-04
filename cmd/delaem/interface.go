package payments

import "fmt"

type Auto interface {
	StepOnGas()
	StepOnBreak()
	StepOnReturn()
}

type BMW struct{}

func (b BMW) StepOnReturn() {
	fmt.Println("Задняя передача BMW!")
}

func (b BMW) StepOnBreak() {
	fmt.Println("Я останавливаюсь")
}

func (b BMW) StepOnGas() {
	fmt.Println("Я БМВ! 550 лошадиных сил! Жмем на газ")
}

type Zhiga struct{}

func (z Zhiga) StepOnReturn() {
	fmt.Println("Задняя передеча Zhiga!")
}

func (z Zhiga) StepOnBreak() {
	fmt.Println("Я останавлоиваюсь!")
}

func (z Zhiga) StepOnGas() {
	fmt.Println("Я жига! Попробую не разволиться!")
}
func ride(auto Auto) {
	fmt.Println("Я водитель!")
	fmt.Println("Я сожусь в свою машину!")
	fmt.Println("И нажимаю на газ...")
	auto.StepOnGas()
	auto.StepOnBreak()
	auto.StepOnReturn()
}
func Nelchan() {
	bmw := BMW{}
	ride(bmw)
	zhiga := Zhiga{}
	ride(zhiga)

}
