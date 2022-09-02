package test

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

const res = `sssldmmdlsss
zmdknd`

func TestString() {
	s := "tt无标题sss"
	fmt.Println("字节数量: ", len(s))
	fmt.Println("String : ", s[0:])
	fmt.Println(res)
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()
	fmt.Println("字符数量: ", utf8.RuneCountInString(s))

	for a, b := range s {
		fmt.Printf("%d\t%q\t%d\n", a, b, b)
	}
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		i += size
	}

	index := strings.LastIndex(s, "标")
	fmt.Println("标的起始位置：", index)
}

// 使用buf 链接字符串
func LinkString(str1 string) string {
	var buf bytes.Buffer
	buf.WriteRune('[')
	for i, v := range str1 {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(v)
	}
	buf.WriteRune(']')
	return buf.String()
}

// 检查字符串是否同文异构
func SameString(str1, str2 string) bool {
	for i := 0; i < len(str1); {
		r, size := utf8.DecodeRuneInString(str1[i:])
		index := strings.LastIndex(str2, string(r))
		if index < 0 {
			return false
		}
		i += size
	}
	for i := 0; i < len(str2); {
		r, size := utf8.DecodeRuneInString(str2[i:])
		index := strings.LastIndex(str1, string(r))
		if index < 0 {
			return false
		}
		i += size
	}
	return true
}
