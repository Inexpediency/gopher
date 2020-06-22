package types

import "fmt"

// iota practice

// Each of the lower five bits of an unsigned integer is assigned a unique name:
type Flags uint

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func FlagsToPrint() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // ...
	TiB
	PiB
	EiB
)

func PrintDegreesOfTwo() {
	fmt.Printf("\tKiB %d\n\tMiB %d\n\tGiB %d\n\tTiB %d\n\tPiB %d\n\tEiB %d\n",
		KiB, MiB, GiB, TiB, PiB, EiB)
}
