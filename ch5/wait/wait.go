package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		if err := WaitForServer(url); err != nil {
			//fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
			//os.Exit(1)
			log.Fatalf("Site is down: %v\n", err)
		}
	}
}

// WaitForServer 尝试连接 URL 对应的服务器
// 在一分钟内使用指数退避策略进行重试
// 所有的尝试失败后返回错误
func WaitForServer(url string) error {
	const timeout = 10 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // 成功
		}
		log.Printf("server not response (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // 指数退避策略
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
