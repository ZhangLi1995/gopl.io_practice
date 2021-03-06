package popcount

// pc[i] 是 i 的种群统计
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount 返回 x 的种群统计
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount1(x uint64) int {
	var ret byte
	for i := 0; i < 8; i++ {
		ret += pc[byte(x>>(i*8))]
	}
	return int(ret)
}

func PopCount2(x uint64) int {
	ret := 0
	for x != 0 {
		if x&1 == 1 {
			ret++
		}
		x >>= 1
	}
	return ret
}

func PopCount3(x uint64) int {
	ret := 0
	for x != 0 {
		ret++
		x &= x - 1
	}
	return ret
}
