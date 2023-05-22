package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountCycle(x uint64) int {
	var res int
	for i := range [8]int{} {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

func PopCountBitShifts(x uint64) int {
	var res int
	for i := range [65]int{} {
		res += int(pc[i] & 1)
	}
	return res
}

func PopCountBitSmart(x uint64) int {
	var res int
	for x > 0 {
		if x > x&(x-1) {
			res++
		}
		x = x & (x - 1)
	}
	return res
}
