package main

import (
	"bufio"
	"fmt"
	"net"
)

type client chan string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	fmt.Println("聊天服务器启动")
	listen, err := net.Listen("tcp", "localhost:8887")
	if err != nil {
		return
	}
	go broadcaster()
	for {
		accept, err := listen.Accept()
		if err != nil {
			return
		}
		handleConn(accept)
	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWrite(conn, ch)
	who := conn.RemoteAddr().String()
	entering <- ch
	messages <- who + " entering"
	buf := bufio.NewScanner(conn)
	for buf.Scan() {
		messages <- ":" + buf.Text()
	}
	leaving <- ch
	messages <- who + " leaving"
	conn.Close()
}
func clientWrite(conn net.Conn, ch <-chan string) {
	for mess := range ch {
		fmt.Fprintf(conn, mess)
	}
}
func broadcaster() {
	clients := make(map[client]bool)
	select {
	case mess := <-messages:
		for cli, _ := range clients {
			cli <- mess
		}
	case cli := <-entering:
		clients[cli] = true
	case cli := <-leaving:
		delete(clients, cli)
		close(cli)
	}
}
