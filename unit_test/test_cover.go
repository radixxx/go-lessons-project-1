package un_test

import (
	"fmt"
	"sync"
)

func wait(number int) int {
	var waitGroup sync.WaitGroup
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go concurent(number, &waitGroup)
	}
	waitGroup.Wait()
	return number
}

func concurent(number int, waitGroup *sync.WaitGroup) {
	useless := number + 2
	fmt.Println(useless)
	waitGroup.Done()
}
