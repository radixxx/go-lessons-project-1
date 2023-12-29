package main

func firs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return firs(n-1) + firs(n-2)
}

func second(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2 // error!
	}
	return second(n-1) + second(n-2)
}

func stringQuantityLong(s string) int {
	if s == "" {
		return 0
	}
	n := 1
	for range s {
		n++
	}
	return n
}

func stringQuantityShort(s string) int {
	return len(s)
}
