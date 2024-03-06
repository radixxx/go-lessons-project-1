package main

import (
	"golang.org/x/exp/constraints"
)

// Note that the identifier of the type parameter is positioned before the type constraint -.
func maxGeneric[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
