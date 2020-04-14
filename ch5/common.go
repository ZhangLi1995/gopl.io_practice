package ch5

import (
	"fmt"

	"golang.org/x/net/html"
)

// Visit 函数会将 n 节点中的每个链接添加到结果中
func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}

// forEachNode 调用 pre(x) 和 post(x) 遍历以 n 为根的树中的每个节点 x
// 两个函数是可选的
// pre 在子节点被访问前（前序）调用
// post 在子节点被访问后（后序）调用
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func F(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	F(x - 1)
}
