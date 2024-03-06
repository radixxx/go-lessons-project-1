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

func existsInSlice(val int, values []int) bool {
	for _, v := range values {
		if val == v {
			return true
		}
	}

	return false
}

func stringExistsInSlice(val string, values []string) bool {
	for _, v := range values {
		if val == v {
			return true
		}
	}

	return false
}

func existsInSliceGen[T comparable](val T, values []T) bool {
	for _, v := range values {
		if val == v {
			return true
		}
	}

	return false
}
