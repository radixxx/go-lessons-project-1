package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	rangeType()
	anMap()

}

func rangeType() {
	storage := [3]int{64, 128, 256}
	for index, element := range storage {
		fmt.Printf("element at index %d is %d\n", index, element)
	}

	//the blank identifier
	for _, element := range storage {
		fmt.Printf("element with index _ is %d\n", element)
	}

	index := [4]int{4, 8, 32, 64}
	for i := 0; i < len(index); i++ {
		fmt.Printf("Ascending index - %d element: %d\n", i, index[i])
	}

	for i := len(index) - 1; i >= 0; i-- {
		fmt.Printf("Descending index - %d element: %d\n", i, index[i])
	}
}

func anMap() {
	pointMap := map[string]Point{}
	anotherPointsMap := make(map[int]Point)

	pointMap["a"] = Point{
		x: 2,
		y: 4,
	}

	fmt.Println(pointMap)
	fmt.Println(pointMap["a"])
	anotherPointsMap[1] = Point{3, 6}
	fmt.Println(anotherPointsMap)
	fmt.Println(anotherPointsMap[1])
}
