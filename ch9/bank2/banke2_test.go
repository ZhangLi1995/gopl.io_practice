package bank2

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestBank2(t *testing.T) {
	var wg sync.WaitGroup

	// Alice
	wg.Add(1)
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		wg.Done()
	}()

	// Bob
	wg.Add(1)
	go func() {
		Deposit(100)
		wg.Done()
	}()

	// 等待两笔交易完成
	wg.Wait()

	if got, want := Balance(), 300; got != want {
		fmt.Printf("goroutine numbers: %v\n", runtime.NumGoroutine())
		t.Errorf("Balance = %d, want %d", got, want)
	} else {
		fmt.Printf("goroutine numbers: %v\n", runtime.NumGoroutine())
	}
}
