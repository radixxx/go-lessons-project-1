package main

import (
	"fmt"
)

func main() {
	fmt.Println("function1:", funIntofunc(first, 123))
	fmt.Println("function2:", funIntofunc(second, 123))
	fmt.Println("Inline:", funIntofunc(func(i int) int { return i * i * i }, 123))

}

func first(i int) int {
	return i + i
}

func second(i int) int {
	return i * i
}

func funIntofunc(f func(int) int, v int) int {
	return f(v)
}
