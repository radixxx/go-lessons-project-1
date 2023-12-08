package main

import "fmt"

func main() {
	johnPc := computerPrice(3, 5)
	tomPc := computerPrice(5, 7)
	alexPc := computerPrice(7, 7)

	fmt.Println(johnPc, tomPc, alexPc)
}

func computerPrice(rate int, power int) int {
	return rate * power
}
