package test

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(" ")
	fmt.Println(os.Args[1:])
}
