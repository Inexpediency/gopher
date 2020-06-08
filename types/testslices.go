package types

import "fmt"

func TestCycleShift() {
	s := []int{0, 1, 2, 3, 4, 5}
	CycleShift(s, 3)
	fmt.Println(s) // Output: [3 4 5 0 1 2]
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

func TestNonEmpty() {
	strings := []string{"Hello", "", "World", "", "!"}

	// Before: []string{"Hello", "", "World", "", "!"}
	fmt.Printf("Before: %#v\n", strings)

	//After: []string{"Hello", "World", "!"}
	fmt.Printf("After: %#v", NonEmpty(strings))
}
