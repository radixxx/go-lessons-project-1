package main

import "fmt"

func doOrNot(callback func(int, int) int, s string) int {
	result := callback(5, 2)
	fmt.Println(s)
	return result
}

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sumCallback := func(n, m int) int {
		return n + m
	}

	result := doOrNot(sumCallback, "plus")
	fmt.Println(result)

	multipleCallback := func(n, m int) int {
		return n * m
	}

	result = doOrNot(multipleCallback, "multiple")
	fmt.Println(result)

	//Closures
	orderPrice := totalPrice(1)
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(10))
	fmt.Println(orderPrice(110))
	fmt.Println(orderPrice(1110))
	fmt.Println(orderPrice(11100))
	fmt.Println(orderPrice(111000))
}
