package main

import (
	"errors"
	"fmt"
	"log"
)

type AppError struct {
	Message string
	Err     error
}

func main() {
	result := myArray()
	fmt.Println(result)
	devideWrapper(5 / 0)
}

// !
func (ae *AppError) Error() string {
	return ae.Message
}

func devideWrapper(a, b int) {
	defer func() {
		var appError *AppError
		if err := recover(); err != nil {
			log.Println("Panic apeared: ", err)
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appError) {
					fmt.Println("AAAA it's PANIC !", err)
				} else {
					fmt.Println("-> ")
					panic("Custom Panic mzf!")
				}
			default:
				panic("PanicHuanic !")
			}
			log.Println(devide(a, b))
		}
	}()
	fmt.Println(a / b)
}

func devide(a, b int) int {
	if b == 0 {
		panic(&AppError{
			Message: "Custom panic ! this is devide by zero...",
			Err:     nil,
		})
	}
	return a / b
}

func myArray() int {
	array := [3]int{3, 6, 9}
	return array[5]
}
