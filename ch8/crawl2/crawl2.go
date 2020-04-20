package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func main() {
	worklist := make(chan []string)
	var n int // 等待发送到任务列表的数量

	// 从命令行参数开始
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// 并发爬取 Web
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// 令牌是一个计数信号量
// 确保并发请求限制在 20 个以内
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 令牌
	list, err := links.Extract(url)
	<-tokens // 释放令牌

	if err != nil {
		log.Print(err)
	}
	return list
}
