package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need a time duration")
		return //!
	}

	var waiter sync.WaitGroup
	waiter.Add(1)

	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Duration(int32(t) * int32(time.Millisecond))
	fmt.Println("Timeout period is %s\n", duration)

	if timeout(&waiter, duration) {
		fmt.Println("Time out !")
	} else {
		fmt.Println("Ok !")
	}

	waiter.Done()
	if timeout(&waiter, duration) {
		fmt.Println("Time out! ")
	} else {
		fmt.Println("OK !")
	}
}

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(5 * time.Second)
		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}
