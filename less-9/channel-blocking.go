package main

import (
	"fmt"
	"time"
)

func helloChannel(done chan bool) {
	fmt.Println("1 hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("2 hello go routine awake and going to write to done")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go helloChannel(done)
	<-done
	fmt.Println("Main received data")
}
