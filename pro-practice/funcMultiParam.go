package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	varFunc(arguments...)
	sum := oneByOne("Adding numbers...", 1, 2, 3, 4, 5, -1, 10)
	fmt.Println("Sum: ", sum)
	s := []int{1, 2, 3}
	sum = oneByOne("Adding numbers...", s...)
	fmt.Println(s)
}

func varFunc(input ...string) {
	fmt.Println(input)
}

func oneByOne(msg string, s ...int) int {
	fmt.Println(msg)
	sum := 0
	for i, a := range s {
		fmt.Println(i, a)
		sum = sum + a
	}
	s[0] = -1000
	return sum
}
