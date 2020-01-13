package main

import (
	"crypto/sha256"
	"fmt"

	"gopl.io/ch2/popcount"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Println("==========>")
	fmt.Printf("c1 bit: %d, c2 bit: %d", CountSHA256Bit(c1), CountSHA256Bit(c2))
}

// 练习 4.1：编写一个函数，用于统计 SHA256 散列中不同的位数
func CountSHA256Bit(bytes [32]byte) int {
	ret := 0
	for _, v := range bytes {
		ret += popcount.PopCount3(uint64(v))
	}
	return ret
}
