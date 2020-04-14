package main

import "fmt"

// squares 函数返回一个函数，后者包含下一次要用到的平方数
// the next squares number each time it is called.

func squares() func() int {
	var x int
	//var y = 1
	//fmt.Println(y)
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
