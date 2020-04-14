package main

import (
	"fmt"
	"os"

	"gopl.io/ch5"

	"golang.org/x/net/html"
)

// Findlinks1 输出从标准输入中读入的 HTML 文档中的所有链接

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range ch5.Visit(nil, doc) {
		fmt.Println(link)
	}
}
