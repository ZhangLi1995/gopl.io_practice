package memo3

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

// 并发安全: 有可能多个 goroutine 重复过去同一个 URL 的值，并造成覆盖更新 map
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Lock()
	if !ok {
		res.value, res.err = memo.f(key)

		// 在两个临界区之间，可能会有多个 goroutine 来计算 f(key) 并且更新 map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	memo.mu.Unlock()
	return res.value, res.err
}
