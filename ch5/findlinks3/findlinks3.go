package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

// breadthFirst 对每个 worklist 元素调用 f
// 并将返回的内容添加到 worklist 中，对每一个元素，最多调用一次f
// f is called at mos once for each item
func breadthFist(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// 开始广度遍历
	// 从命令行参数开始
	breadthFist(crawl, os.Args[1:])
}
