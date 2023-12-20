package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
}

func writeToChannel(c chan int, x int) {
	fmt.Println(x)

	c <- x
	close(c)
	fmt.Println(x)
}
