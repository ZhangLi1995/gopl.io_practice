package main

import (
	"os"
	"runtime"

	"gopl.io/ch5"
)

func main() {
	ch5.F(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
