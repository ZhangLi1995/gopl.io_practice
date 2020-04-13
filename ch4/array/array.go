package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	// 输出索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	symbol := [...]string{USD: "$", EUR: "e", GBP: "f", RMB: "y"}
	fmt.Println(RMB, symbol[RMB])
}
