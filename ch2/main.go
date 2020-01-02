package main

import (
	"fmt"
	"time"

	"gopl.io/ch2/popcount"
)

func main() {
	x := uint64(34646836576236000)
	now := time.Now().UnixNano()
	fmt.Printf("original PopCount result: %v, cost time: %vns\n", popcount.PopCount(x), time.Now().UnixNano()-now)
	now = time.Now().UnixNano()
	fmt.Printf("Exercise PopCount1 result: %v, cost time: %vns\n", popcount.PopCount1(x), time.Now().UnixNano()-now)
	now = time.Now().UnixNano()
	fmt.Printf("original PopCount2 result: %v, cost time: %vns\n", popcount.PopCount2(x), time.Now().UnixNano()-now)
	now = time.Now().UnixNano()
	fmt.Printf("original PopCount3 result: %v, cost time: %vns\n", popcount.PopCount3(x), time.Now().UnixNano()-now)
}
