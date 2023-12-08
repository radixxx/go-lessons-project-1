package main

import (
	"fmt"
	"time"
)

func main() {
	slices()
	borschSlice()
}

func slices() {
	cars := []string{"bmw", "audi", "porsche", "toyta", "opel", "reno"}
	cars[4] = "mercedes"
	cars = append(cars, "lexus", "vw")

	fmt.Println(cars[7])

	var interfaceSlice []interface{}
	start := time.Now() //append different data types
	interfaceSlice = append(interfaceSlice, 77, 3.14, "bmw", start)
	fmt.Println("Borsch: ", interfaceSlice)

	for i, j := 0, len(interfaceSlice)-1; i < j; i, j = i+1, j-1 {
		interfaceSlice[i], interfaceSlice[j] = interfaceSlice[j], interfaceSlice[i]
	}

	slicelengh := len(interfaceSlice)
	fmt.Println("Slice lengh: ", slicelengh)

	slieceCapacity := cap(interfaceSlice)
	fmt.Println("Slice capacity: ", slieceCapacity)

	interfaceSlice = append(interfaceSlice, "1812")
	fmt.Println(interfaceSlice)

	fmt.Println("Slice lengh: ", len(interfaceSlice))
	fmt.Println("Slice capacity: ", cap(interfaceSlice))
}

func borschSlice() {
	vegetablesArray := [7]string{"potato", "bean", "lettuce", "carrot", "garlic", "onion", "kvas"}

	vegetablesSlice := vegetablesArray[:4]
	var anotherVagetableSlice []string = vegetablesArray[2:]

	fmt.Println(vegetablesSlice)
	fmt.Println(anotherVagetableSlice)

	vegetablesSlice[2] = "mushroom"
	fmt.Println("Updated slice: ", vegetablesSlice)
	fmt.Println("Updated slice", anotherVagetableSlice)

	borshMixed := vegetablesArray[:]
	fmt.Println("Show all vegetables :", borshMixed)
}

func arrays() {
	var array [3]string
	array[0] = "Hola"
	array[1] = "Get Up"
	array[2] = "Mudila"

	fmt.Println(array[2])

	numbers := [...]int{1, 3, 4, 5, 7}
	numbers[3] = 9
}
