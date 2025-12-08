package p

import (
	"fmt"
)

type Person struct {
	Age  int
	Name string
}

func Arrays() {

	var intArr [3]int

	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)
	intArr[0] = 5
	intArr[1] = 6
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)

	people := [2]Person{
		{
			Age:  30,
			Name: "Katy",
		},
	}
	fmt.Printf("Type: %T Value: %#v\n", people, people)
	stringsArr := [...]string{"Fiurst", "Second", "Third", "Fourth", "Second"}
	fmt.Printf("Type: %T Value: %#v\n", stringsArr, stringsArr)
	for i := 0; i < len(stringsArr); i++ {
		fmt.Printf("Type: %T Value: %#v\n", stringsArr[i], stringsArr[i])
	}
	for inx, value := range stringsArr {
		fmt.Printf("Type: %T Value: %#v\n", inx, value)
	}
	for _, value := range stringsArr {
		fmt.Printf("Type: %T Value: %#v\n", value, value)
	}

}
