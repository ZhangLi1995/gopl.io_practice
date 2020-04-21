package bank1

import (
	"fmt"
	"testing"
)

func TestBank1(t *testing.T) {
	done := make(chan struct{})
	//var wg sync.WaitGroup

	// Alice
	//wg.Add(1)
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
		//wg.Done()
	}()

	// Bob
	//wg.Add(1)
	go func() {
		Deposit(100)
		done <- struct{}{}
		//wg.Done()
	}()

	// 等待两笔交易完成
	<-done
	<-done
	// wg.Wait()

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
