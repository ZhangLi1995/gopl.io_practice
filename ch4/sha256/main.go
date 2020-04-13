package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"

	"gopl.io/ch2/popcount"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Println("==========>")
	fmt.Printf("c1 bit: %d, c2 bit: %d\n", CountSHA256Bit(c1), CountSHA256Bit(c2))

	fmt.Println("==========>")
	OutputSHA("x")
}

// 练习 4.1：编写一个函数，用于统计 SHA256 散列中不同的位数
func CountSHA256Bit(bytes [32]byte) int {
	ret := 0
	for _, v := range bytes {
		ret += popcount.PopCount3(uint64(v))
	}
	return ret
}

// 练习 4.2：编写一个程序，用于在默认情况下输出其标准输入的 SHA256 散列，但也支持一个输出 SHA384 或 SHA512 散列的命令行标记
func OutputSHA(s string) {
	if os.Args[1] == "384" {
		c := sha512.New384().Sum([]byte(s))
		fmt.Printf("%x\n%T\n", c, c)
	} else if os.Args[1] == "512" {
		c := sha512.Sum512([]byte(s))
		fmt.Printf("%x\n%T\n", c, c)
	} else {
		c := sha256.Sum256([]byte(s))
		fmt.Printf("%x\n%T\n", c, c)
	}
}
