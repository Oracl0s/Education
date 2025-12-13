package p

import "fmt"

type User struct {
	Id   int64
	Name string
}

func Mapss() {
	var defaultMaps map[int64]string
	fmt.Println(defaultMaps) //map[]

	mapByMake := make(map[string]string)
	fmt.Println(mapByMake) //map[]

	mapByLiteral := map[string]int{"Vasya": 18, "Dima": 20}
	fmt.Println(mapByLiteral) //map[Dima:20 Vasya:18]

	mapWithNew := *new(map[string]string)
	fmt.Println(mapWithNew) //map[]

	mapByMake["First"] = "Vasya"
	fmt.Println(mapByMake) //map[First:Vasya]

	mapByMake["Second"] = "Dima"
	fmt.Println(mapByMake) //map[First:Vasya Second:Dima]

	mapByMake["First"] = "Ilya"
	fmt.Println(mapByMake) //map[First:Ilya Second:Dima]

	value, ok := mapByMake["First"] //Проверка переменых Ilya true
	fmt.Println(value, ok)

	mapForIteration := map[string]int{"first": 1, "second": 2, "third": 3, "fourth": 4} // Нету повторяемости последовательности (при каждом запуске разные последовательности)

	for key, val := range mapForIteration {
		fmt.Println(key, val)
	}
	users := []User{
		{
			Id:   1,
			Name: "Vasya",
		},
		{
			Id:   45,
			Name: "John",
		},
		{
			Id:   57,
			Name: "Petya",
		},
	}
	uniqueUsers := make(map[int64]struct{}, len(users))
	for _, user := range users {
		if _, ok := uniqueUsers[user.Id]; !ok {
			uniqueUsers[user.Id] = struct{}{}
		}
	}
	fmt.Println(uniqueUsers)

}
