package bank3

import (
	"fmt"
	"runtime"
	"testing"
)

func TestBank3(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// 等待两笔交易完成
	<-done
	<-done
	fmt.Printf("goroutine numbers: %v\n", runtime.NumGoroutine())
	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
