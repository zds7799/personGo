package chapter

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"sort"
)

var progress = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func parse() {
}

// TestParameter 测试参数传递
func TestParameter() {
	var a = 1
	var str = "strs"
	var b []byte
	var bp []byte
	var bb [3]byte
	var bbp [3]byte
	fmt.Printf("原始数据：a = %d; str = %s; b = %v; bp = %v; bb = %v ;  bbp = %v\n", a, str, b, bp, bb, bbp)
	testParameter2(a, str, b, &bp, bb, &bbp)
	fmt.Printf("修改后数据：a = %d; str = %s; b = %v; bp = %v; bb = %v ;  bbp = %v\n", a, str, b, bp, bb, bbp)
}
func testParameter2(a int, str string, b []byte, bp *[]byte, bb [3]byte, bbp *[3]byte) {
	a = 2
	str = "s"
	b = append(b, 1, 2, 3)
	*bp = append(*bp, 2, 3)
	bb[1] = 2
	(*bbp)[2] = 3
	fmt.Printf("修改后数据：a = %d; str = %s; b = %v; bp = %v; bb = %v ;  bbp = %v\n", a, str, b, *bp, bb, *bbp)
}
func buffToString(buff []int32) string {
	var buf bytes.Buffer
	for _, v := range buff {
		buf.WriteRune(v)
	}
	return buf.String()
}

// TopSort 匿名函数变量 拓扑排序
func TopSort() {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(item []string)
	visitAll = func(item []string) {
		for _, k := range item {
			if !seen[k] {
				seen[k] = true
				visitAll(progress[k])
				order = append(order, k)
			}
		}
	}
	var key []string
	for k, _ := range progress {
		key = append(key, k)
	}
	sort.Strings(key)
	visitAll(key)
	fmt.Println("拓扑排序：")
	for i, s := range order {
		fmt.Printf("%d\t%s\n", i, s)
	}
}

// Extract 解析url
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parse %s as HTML: %v", url, err)
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
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
