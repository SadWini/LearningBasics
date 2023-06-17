package studentpairs

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", n)
	var x, y []int
	for i := 0; i < 2*n; i++ {
		fmt.Scanf("%d%d", x[i], y[i])
	}
	var pairCost []int
	for i := 0; i < 2*n; i++ {
		for j := 0; j < 2*n; j++ {
			pairCost[2*n*i+j] = (x[i]-x[j])*(x[i]-x[j]) + (y[i]-y[j])*(y[i]*y[j])
		}
	}
	/*
		I realized that an exhaustive search takes too long.
		16! / 2^8 sets is more than 10^10, so I get the time limit exception.
		Thus, I am an uncompetitive type A programmer
	*/
	// ex 1.1.1 {(1, 0), (3,0), (4, 0), (6, 0)}
	// ex 1.1.2 O(N!/2^(N/2))
}
