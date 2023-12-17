package main

import "fmt"

func main() {
	ch5 := make(chan int, 2)
	ch5 <- 42
	close(ch5)
	received, ok := <-ch5
	fmt.Println(received, ok)
}
