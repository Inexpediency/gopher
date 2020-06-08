package types

func Reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// CycleShift of slice
func CycleShift(s []int, n int) {
	if n > len(s) {
		n = n % len(s)
	}

	Reverse(s[:n])
	Reverse(s[n:])
	Reverse(s)
}
