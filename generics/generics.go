package main

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxGenerec[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
