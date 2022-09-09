package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8887")
	if err != nil {
		return
	}
	defer c.Close()
	go func() {
		if _, err := io.Copy(os.Stdout, c); err != nil {
			log.Fatal(err)
		}
	}()
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		switch args[0] {
		case "close":
			fmt.Fprintf(c, sc.Text())
			return
		default:
			fmt.Fprintln(c, sc.Text())
		}
	}
}
