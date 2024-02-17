package main

import "fmt"

type Teacher struct {
	in        int
	firstname string
	lastname  string
}

func (t *Teacher) say() {
	fmt.Println("Hi class my name is %s %s", t.firstname, t.lastname)
}

func (s *StatusCode) String() string {
	return fmt.Sprintf("Status code %d", s)
}
