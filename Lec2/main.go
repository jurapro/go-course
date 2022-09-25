package main

import (
	"fmt"
	"math"
)

func main() {
	var age bool
	fmt.Println("Hello ", age)
	age = true
	fmt.Println("Hello ", age)
	fmt.Printf("%T\n", age)

	var (
		personName = "Иван"
		personAge  = 20
	)
	personUID := 1234

	fmt.Printf("Имя пользователя: %s, возраст: %d, UID: %d \n", personName, personAge, personUID)

	x, y := 100.23, -12.00

	fmt.Printf("Минимум двух чисел %.4f\n", math.Min(x, y))
}
