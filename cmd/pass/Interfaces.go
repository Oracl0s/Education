package p

import "fmt"

type Greeter interface {
	greet(string) string
}
type Russian struct{}
type American struct{}

func (r *Russian) greet(name string) string {
	return fmt.Sprintf("Привет, %s", name)
}
func (a *American) greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
func SayHello(g Greeter, name string) {
	fmt.Println(g.greet(name))
}

type Animal interface {
	makeSound()
}
type cat struct{}
type dog struct{}

func (c *cat) makeSound() {
	fmt.Println("I am Cat")
}
func (d *dog) makeSound() {
	fmt.Println("I am Dog")
}
func Inter() {
	var c, d Animal = &cat{}, &dog{}
	c.makeSound()
	d.makeSound()
}
