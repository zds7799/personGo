package chapter8

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

/*
*
1、goroutine
2、并发时钟 服务器客户端启动，交互
3、并发回声
4、通道

	无缓冲通道 同步通道 定义
	管道 定义
	单项通道
	缓存通道  容量大小 阻塞情况

5、并发循环
6、并发web爬虫 示例
7、select多路复用
8、并发目录遍历 示例
9、取消
10、聊天服务器 示例
*/

// TestComplicated 并发循环 通道  WaitGroup
func TestComplicated() {
	var ch = make(chan int64)
	var wp sync.WaitGroup
	for i := 0; i < 15; i++ {
		wp.Add(1)
		go func(i int) {
			defer wp.Done()
			s := rand.Int() % 10
			time.Sleep(time.Duration(s) * time.Second)
			ch <- int64(s)
		}(i)
	}
	go func() {
		wp.Wait()
		close(ch)
	}()
	for t := range ch {
		fmt.Printf("sleep %d second\n", t)
	}
}

type urlString struct {
	url   string
	depth int
}

// CrewUrl 并发抓去网页链接
func CrewUrl(strs []string) {
	works := make(chan []urlString)
	go func() {
		var urls []urlString
		for _, u := range strs[:] {
			urls = append(urls, urlString{u, 1})
		}
		works <- urls
	}()
	var n int
	n++
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		work := <-works
		for _, s := range work {
			if !seen[s.url] && s.depth < 3 {
				n++
				go func(s urlString) {
					works <- crew(s)
				}(s)
			}
		}
	}
}

var token = make(chan struct{}, 20)

func crew(s urlString) []urlString {
	fmt.Println(s.url)
	token <- struct{}{}
	strings, err := extract(s.url)
	<-token
	if err != nil {
		log.Print(err)
	}
	var urls []urlString
	for _, s2 := range strings {
		urls = append(urls, urlString{s2, s.depth + 1})
	}
	return urls
}
func extract(url string) ([]string, error) {
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
					continue // ignore bad URLs
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

// TestTimer 定时执行 无法主动终止计数
func TestTimer() {
	fmt.Printf("")
	tick := time.Tick(2 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \t", i)
		fmt.Printf("time %v\n", <-tick)
	}
}

// TestTimer2 定时执行 可以终止
func TestTimer2() {
	tick := time.NewTicker(2 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-time.After(12 * time.Second):
			fmt.Println("Done!")
			return
		case t := <-tick.C:
			fmt.Println("Current time: ", t)
		}
	}
}
