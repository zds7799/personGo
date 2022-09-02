package test

import (
	"fmt"
	"os"
	"strings"
)

func Dup() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

	}
	for line, arg := range counts {
		fmt.Printf("%d\t%s\n", arg, line)
	}
}
