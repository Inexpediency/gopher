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

// AppendInt to slice
func AppendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room for growth. Expanding the slice
		z = x[:zlen]
	} else {
		// There is no room for growth. Selecting a new array.
		// Double it for linear amortized complexity,
		zcap := zlen
		if zcap	< 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)

	return z
}

// NonEmpty return array without empty strings
func NonEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
