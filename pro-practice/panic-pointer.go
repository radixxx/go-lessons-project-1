package main

import (
	"fmt"
)

type Shark struct {
	Name string
}

func main() {
	s := &Shark{"Bobby"}
	s.SayHello()
	// ! nullpointer
	s = nil
	s.SayHello()
}

func (s *Shark) SayHello() {
	fmt.Println("Hi! My name is", s.Name)
}
