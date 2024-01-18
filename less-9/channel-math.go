package main

import "fmt"

func calcSquares(number int, square chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	square <- sum
}

func calcCubes(number int, cube chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cube <- sum
}

func main() {
	number := 925
	sqrch := make(chan int)
	cubeh := make(chan int)

	go calcSquares(number, sqrch)
	go calcCubes(number, cubeh)

	sqs := <-sqrch
	cubes := <-cubeh
	fmt.Println("Final result: ", sqs+cubes)
}
