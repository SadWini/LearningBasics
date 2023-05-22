package main

import (
	"fmt"
	strings2 "strings"
)

func main() {
	data := []string{"one", "", "threee"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
}

func elimDups(strings []string) []string {
	i := 0
	for v, s := range strings {
		if !(strings2.Contains(s[:v-1], s)) {
			strings[i] = strings[v]
			i++
		}
	}
	return strings[:i]
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
