package unsafeusing

// #include <stdio.h>
// void callC() {
//    printf("Calling C code!\n");
// }
import "C"

import "fmt"

// CGO runs C and Go code!
func CGO() {
	fmt.Println("A Go statement!")
	C.callC()
}
