package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
	"gopl.io/ch5"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查 content-Type 是 HTML (如 "text/html; charset=utf-8")
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	ch5.ForEachNode(doc, visitNode, nil)
	return nil
}

func main() {
	for _, url := range os.Args[1:] {
		if err := title(url); err != nil {
			log.Fatalf("%v", err)
		}
	}
}
