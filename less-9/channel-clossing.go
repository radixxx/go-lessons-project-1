package main

import "fmt"

func main() {
	go simple()
	go better()
}

func simple() {
	ch := make(chan int)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

func better() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
