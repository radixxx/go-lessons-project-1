package main

import "fmt"

type Instrument interface {
	Play()
}

func PlayInstrument(i Instrument) {
	i.Play()
}

type Guitar struct {
	Type string
}

func (g Guitar) Play() {
	fmt.Println("Guitar Sounds")
}

func main() {
	g := Guitar{"Musica bleat"}
	PlayInstrument(g)
}

/*
	Golang supports polymorphism through interfaces only.
	A type implements an interface if it provides definitions for all the methods declared in the interface.
*/
