// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	sep = " "
	for i := 1; i < len(os.Args); i++ {
		s = os.Args[i]
		fmt.Println(i, sep, s)
	}
}
