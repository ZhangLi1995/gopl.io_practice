package main

import (
	"fmt"
	"reflect"
	"time"
)

func sayHello() {
	for {
		fmt.Println("Hello gorotine")
		time.Sleep(time.Second)
	}
}

func gen(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {
	//defer fmt.Println(runtime.NumGoroutine())
	//limit := time.NewTicker(1 * time.Second)
	//for i := 0; i < 10; i++ {
	//	<-limit.C
	//	fmt.Println(runtime.NumGoroutine())
	//	fmt.Printf("After 1s: %v\n", time.Now().Format("2006-01-02_15:04:05"))
	//}
	//limit.Stop()

	//defer func() {
	//	fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	//}()
	//
	//go sayHello()
	//fmt.Println("Hello main")

	//defer func() {
	//	time.Sleep(time.Second)
	//	fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	//}()
	//
	//// Set up the pipeline.
	//done := make(chan struct{})
	//defer close(done)
	//
	//out := gen(done, 2, 3)
	//
	//for n := range out {
	//	fmt.Println(n)              // 2
	//	time.Sleep(5 * time.Second) // done thing, 可能异常中断接收
	//	if true {                   // if err != nil
	//		break
	//	}
	//}

	//m := map[string]string{
	//	"a": "A",
	//}
	//fmt.Printf("%v, %T\n", m["b"], m["b"])

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	fmt.Println(v.String())
}
