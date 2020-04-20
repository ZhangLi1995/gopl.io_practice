package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // 通道关闭并且读完
			}
			squares <- x * x
		}
		close(squares)
	}()

	// print
	for {
		fmt.Println(<-squares)
	}
}
