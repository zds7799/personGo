package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8887")
	if err != nil {
		return
	}
	defer func(con net.Conn) {
		err := con.Close()
		if err != nil {
			log.Println(err)
		}
	}(con)
	go func() {
		if _, err := io.Copy(os.Stdout, con); err != nil {
			log.Println(err)
		}
	}()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		_, err := fmt.Fprintf(con, sc.Text())
		if err != nil {
			log.Println(err)
			return
		}
	}
}
