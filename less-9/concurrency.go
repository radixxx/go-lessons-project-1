package main

import (
	"fmt"
	"time"
)

func main() {
	//panic("Show me panic bleat !")e
	fmt.Println("launch goroutine 1")
	go printNumber()
	fmt.Println("launch goroutine 2")
	go printNumber()
	fmt.Println("launch goroutine 3")
	go printNumber()
	fmt.Println("launch goroutine 4")
	go printNumber()
	time.Sleep(1 * time.Minute)
}

func printNumber() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		i++
		fmt.Println(i)
	}
}
