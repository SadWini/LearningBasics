package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	for range [5]int{} {
		fmt.Println(f())
	}
}
