package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// square
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// printer: 在主 goroutine 中
	for x := range squares {
		fmt.Println(x)
	}
}
