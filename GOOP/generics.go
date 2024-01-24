package main

import "fmt"

func PlayInstrument[I Instrument](i I) {
	i.Play()
}

type Instrument interface {
	Play()
}

type Guitar struct {
	Sound string
}

func (g Guitar) Play() {
	fmt.Println(g.Sound)
}

func main() {
	g := Guitar{"Guitar Sound"}
	PlayInstrument(g)
}
