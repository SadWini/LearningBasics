package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopBack
	FlagPointToPoint
	FlagMultiCast
)

func main() {
	var v Flags = FlagMultiCast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadCast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}
func TurnDown(v *Flags) {
	*v &^= FlagUp
}
func SetBroadCast(v *Flags) {
	*v |= FlagBroadcast
}
func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMultiCast) != 0
}
