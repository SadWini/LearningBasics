package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...)
	fmt.Println(x)
	x = []int{1}
	var y = []int{1, 2, 3}
	x = appendIntLong(x, y...)
	fmt.Println(x)
}

func appendIntLong(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	z = make([]int, zlen, 2*zlen)
	copy(z[:len(x)], x)
	copy(z[len(x):], y)
	return z
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
