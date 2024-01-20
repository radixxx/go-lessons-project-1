package main

import "fmt"

func main() {
	//	deadLoh()  # it's DEADLOCK example !
	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}

/*
	channel ch is created and we send 5 to the channel in line no. 6 ch <- 5.
	In this program no other Goroutine is receiving data from the channel ch.
*/

func deadLoh() {
	ch := make(chan int)
	ch <- 5
}

/*
	It is possible to convert a bidirectional channel to a send only or receive only channel
	but not the vice versa.
*/

func sendData(sendch chan<- int) {
	sendch <- 10
}
