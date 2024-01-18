package main

import (
	"fmt"
)

func helloChannel(done chan bool) {
	fmt.Println("Hello 1-st Channel")
	done <- true
}

func main() {
	done := make(chan bool)
	go helloChannel(done)
	<-done
	fmt.Println("main-function")
}
