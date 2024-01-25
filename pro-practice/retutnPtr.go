package main

import (
	"fmt"
)

func main() {
	sq := returnPointer(10)
	fmt.Println("Sq value: ", sq)
	fmt.Println("Sq value: ", *sq)
}

func returnPointer(x int) *int {
	y := x * x
	return &y
}
