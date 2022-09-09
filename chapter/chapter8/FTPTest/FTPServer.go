package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8887")
	if err != nil {
		return
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			return
		}
		go handleCoon(accept)
	}
}
func handleCoon(c net.Conn) {
	defer c.Close()
	cwd := "."
	sc := bufio.NewScanner(c)
	for sc.Scan() {
		args := strings.Fields(sc.Text())
		switch args[0] {
		case "close":
			return
		case "ls":
			if len(args) < 2 {
				ls(c, cwd)
			} else {
				fileName := args[1]
				if err := ls(c, cwd+fileName); err != nil {
					fmt.Fprint(c, err)
				}
			}
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(c, "not enough argument")
			} else {
				cwd += "/" + args[1]
			}
		case "get":
			if len(args) < 2 {
				fmt.Fprintln(c, "not enough argument")
			} else {
				filename := args[1]
				data, err := os.ReadFile(filename)
				if err != nil {
					fmt.Fprint(c, err)
				}
				fmt.Fprintf(c, "%s\n", data)
			}
		default:
			io.WriteString(c, fmt.Sprintf("command not found: %v", args[0]))
		}
	}
}
func ls(c net.Conn, cur string) error {
	dir, err := os.ReadDir(cur)
	if err != nil {
		return err
	}
	for _, entry := range dir {
		io.WriteString(c, fmt.Sprintf("%s\n", entry.Name()))
	}
	return nil
}
