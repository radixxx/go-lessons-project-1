package main

import (
	"testing"
)

/*
**

	`-v` show detailed command line result
	`-cover` show test covereage

**
*/

func TtestStringQuantityShort(t *testing.T) {
	if s1("123456789") != 9 {
		t.Error(`s1("123456789") != 9`)
	}
	if s1("") != 0 {
		t.Error(`s1("") != 0`)
	}
}

func TtestStringQuantityLong(t *testing.T) {
	if s2("123456789") != 9 {
		t.Error(`s2("123456789") != 9`)
	}
	if s2("") != 0 {
		t.Error(`s2("") != 0`)
	}
}

func TtestFirst(t *testing.T) {
	if first(0) != 0 {
		t.Error(`first(0) != 0`)
	}
	if first(1) != 1 {
		t.Error(`first(1) != 1`)
	}
	if first(2) != 1 {
		t.Error(`first(2) != 1`)
	}
	if first(10) != 55 {
		t.Error(`first(10) != 55`)
	}
}

func TtestSecond(t *testing.T) {
	if second(0) != 0 {
		t.Error(`second(0) != 0`)
	}
	if second(1) != 1 {
		t.Error(`second(1) != 1`)
	}
	if second(2) != 1 {
		t.Error(`second(2) != 1`)
	}
	if second(10) != 55 {
		t.Error(`second(10) != 55`)
	}
}
