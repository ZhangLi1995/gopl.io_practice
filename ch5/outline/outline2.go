package main

import (
	"fmt"
	"os"

	"gopl.io/ch5"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	ch5.ForEachNode(doc, startElement, endElement)
}

// forEachNode 调用 pre(x) 和 post(x) 遍历以 n 为根的树中的每个节点 x
// 两个函数是可选的
// pre 在子节点被访问前（前序）调用
// post 在子节点被访问后（后序）调用
//func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
//	if pre != nil {
//		pre(n)
//	}
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		forEachNode(c, pre, post)
//	}
//	if post != nil {
//		post(n)
//	}
//}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}
