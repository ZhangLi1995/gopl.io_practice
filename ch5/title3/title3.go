package main

import (
	"fmt"

	"golang.org/x/net/html"
	"gopl.io/ch5"
)

// soleTitle 返回文档中的第一个非空标题元素
// 如果没有标题则返回错误
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		//没有宕机
		case bailout{}:
			// 预期的宕机
			err = fmt.Errorf("mutiple title elemnts")
		default:
			panic(p) // 未预期的宕机；继续宕机过程
		}
	}()

	// 如果发现多余一个非空标题，退出递归
	ch5.ForEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // 多个标题元素
			}
		}
		title = n.FirstChild.Data
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func main() {

}
