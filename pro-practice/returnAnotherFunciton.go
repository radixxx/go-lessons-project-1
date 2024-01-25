package main

import "fmt"

func main() {
	i := funcReturn()
	j := funcReturn()

	fmt.Println("i:", i())
	fmt.Println("i-1: ", i())
	fmt.Println("j:", j())
	fmt.Println("j-1: ", j())
	fmt.Println("i-3: ", i())
}

func funcReturn() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}
