package main

import "fmt"

func doOrNot(callback func(int, int) int, s string) int {
	result := callback(5, 2)
	fmt.Println(s)
	return result
}

func main() {
	sumCallback := func(n, m int) int {
		return n + m
	}

	result := doOrNot(sumCallback, "plus")
	fmt.Println(result)

	johnPc := computerPrice(3, 5)
	tomPc := computerPrice(5, 7)
	alexPc := computerPrice(7, 7)

	fmt.Println(johnPc, tomPc, alexPc)
}

func computerPrice(rate int, power int) int {
	return rate * power
}
