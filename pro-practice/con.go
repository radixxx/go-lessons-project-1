package main

import (
	"math/rand"
)

var CLOSE = false
var DATA = make(map[int]bool)

func main() {
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSE {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
