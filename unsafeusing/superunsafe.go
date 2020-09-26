package unsafeusing

import (
	"fmt"
	"unsafe"
)

// TestSuperUnsafe ...
func TestSuperUnsafe() {
	// Output is: 0 1 -2 3 4 824634142496
	array := []int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Print(*pointer, " ")
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

	for range array {
		pointer = (*int)(unsafe.Pointer(memoryAddress)) // go-vet
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}

	fmt.Println()
}
