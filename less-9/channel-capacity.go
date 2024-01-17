package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	go may(ch)
	log.Println("waiting for reception...")
	ch <- 55
	log.Println("received")
}

func may(c chan int) {
	time.Sleep(3 * time.Second)
	<-c
}

/*
Created an unbuffered channel of integers ch
A second goroutine is launched with go may(ch)
The goroutine will wait 3 seconds then receive data on the input channel.
The data is received is discarded (we do not store it into a variable)
In the main goroutine we send the value 55 into the channel.

This program outputs: bla bla

The send operation blocks the main goroutine
It is unblocked when the dummy goroutine receive the value sent
*/
