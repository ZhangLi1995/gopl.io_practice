package links

import (
	"fmt"
	"net/http"

	"gopl.io/ch5"

	"golang.org/x/net/html"
)

// link 包提供了解析链接的函数
// Extract 函数向给定的 URL 发起 HTTP GET 请求
// 解析 HTML 并返回 HTML 文档中存在的链接

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // 忽略不合法的 url
				}
				links = append(links, link.String())
			}
		}
	}
	ch5.ForEachNode(doc, visitNode, nil)
	return links, nil
}
