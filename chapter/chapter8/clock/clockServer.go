package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8887")
	if err != nil {
		return
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Printf("链接终止\n")
			continue
		}
		go handleConn(accept)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
