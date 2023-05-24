package main

import "fmt"

func justPanic() (x int, err error) {
	defer func() {
		p := recover()
		x = 1
		err = fmt.Errorf("internal error: %v", p)
	}()
	panic("Function always calls panic")
}
