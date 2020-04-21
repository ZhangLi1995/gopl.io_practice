package bank3

import "sync"

var (
	mu      sync.Mutex // 保护 balance
	balance int
)

// Go语言互斥锁无法重入，所以可能需要利用封装将一些函数拆成两部分
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // 余额不足
	}
	return true
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

// 调用这个函数要求获取互斥锁
func deposit(amount int) {
	balance += amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
