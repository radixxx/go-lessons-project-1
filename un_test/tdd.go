package un_test

import (
	"testing"
)

var vowels = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true}

func TestVowelCount(t *testing.T) {
	expected := uint(5)
	actual := VowelCount("Bla bla !")
	if actual != expected {
		t.Error("actual %d, expected %d", actual, expected)
	}
}

func VowelCount(sentence string) uint {
	var count uint
	for _, char := range sentence {
		if vowels[string(char)] {
			count++
		}
	}
	return count
}
