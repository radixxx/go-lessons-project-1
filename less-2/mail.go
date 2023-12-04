package main

import "fmt"

const pi float32 = 3.14
const carMaxSpeed = 310

func main() {
	fmt.Println(pi, carMaxSpeed)
	fmt.Println(threeArguments())
	fmt.Println(nonDeclaredReturn())
	fmt.Println(loop())
}

func threeArguments() (string, string, string) {
	a := "Hello"
	b := "programmer"
	c := "mundo"

	return a, b, c
}

func nonDeclaredReturn() (f string) {
	f = "Return Type !"
	return
}

func loop() (sum int) {
	sum = 0
	count := 10
	for i := 0; i < count; i++ {
		sum += i
	}
	return sum
}
