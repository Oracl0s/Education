package p

import "fmt"

func Slices() {
	//var defaultSlic [40]int
	//fmt.Println(defaultSlic)
	//var def [13]int
	//fmt.Println(def)
	intSlice := [...]int{1, 2, 3, 4, 5, 6, 8, 9, 0}
	fmt.Println(intSlice)
	intArr := [...]int{1, 2, 3, 4, 5, 6}
	fullSlice := intArr[:]
	fmt.Println(fullSlice)
}
