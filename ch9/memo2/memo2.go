package memo2

import "sync"

// Memo 缓存了调用 Func 的结果
type Memo struct {
	f     Func
	mu    sync.Mutex // 保护 cache
	cache map[string]result
}

// Func 是用于记忆的函数类型
type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

// 并发安全：但是退化成串行
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
