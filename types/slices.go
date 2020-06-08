package types

import "fmt"

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

func TestCycleShift() {
	s := []int{0, 1, 2, 3, 4, 5}
	CycleShift(s, 3)
	fmt.Println(s) // Output: [3 4 5 0 1 2]
}

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

func TestAppendInt() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = AppendInt(x, i)
		fmt.Printf("%d cap = %d\t%v\n", i, cap(y), y)
		x = y
	}

	fmt.Print("\n\n")

	// Append one more elements with ... syntax
	x, y = []int{}, []int{}
	for i := 1; i <= 3; i++ {
		x = append(x, i*2)
		y = AppendInt(x, x...)
		fmt.Printf("%d cap = %d\t%v\n", i, cap(y), y)
		x = y
	}
}