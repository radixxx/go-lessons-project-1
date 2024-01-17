package main

import (
	"fmt"
	"time"
)

func numbers() {
	for iter := 1; iter <= 5; iter++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d ", iter)
	}
}

func alpha() {
	for iter := 'a'; iter <= 'z'; iter++ {
		time.Sleep(2 * time.Second)
		fmt.Printf("%c ", iter)
	}
}

func main() {
	go numbers()
	go alpha()
	time.Sleep(5 * time.Second)
	fmt.Println("main-terminated")
}
