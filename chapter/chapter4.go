package chapter

import (
	"fmt"
	"unicode/utf8"
)

const resStr = `"原生字符串
变量"`

func BufTest() {
	var a int = 1
	var b byte = 1
	c := 3
	d := new(int)
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("c: %T\n", &c)
	fmt.Printf("d: %T\n", d)
	var str string = "as中文jkl"
	fmt.Println(str[3])
	fmt.Println(resStr)
}
func init() {

}

// BuffTest 数组
func BuffTest() {
	//数组定义
	var a [3]int
	var b = [...]int{1, 2, 3}
	var c = [...]string{2: "ag", 3: "中文", 1: "shd"}
	var d = [...]int{1, 2, 3, 5: 6}
	//数组指针
	var pa = &a
	//数组常用方法
	fmt.Printf("数组长度len(c) : %d\n", len(c)) //数组的长度
	fmt.Printf("数组空间cap(c) : %d\n", cap(c)) //等于数组的长度
	//数组遍历
	fmt.Print("数组遍历pa: ")
	for i, v := range pa {
		fmt.Printf("%d:%d\t", i, v)
	}
	fmt.Println()
	fmt.Print("数组遍历b: ")
	for i := range b {
		fmt.Printf("%d:%d\t", i, b[i])
	}
	fmt.Println()
	fmt.Print("数组遍历d: ")
	for i := 0; i < len(d); i++ {
		fmt.Printf("%d:%d\t", i, d[i])
	}
	//接口数组 管道数组
}

// ParameterBuf 数组为参数
func ParameterBuf() {
	var buf = [3]int32{1, 2, 3}
	fmt.Println("修改前：", buf[0])
	parameterBuf2(buf)
	fmt.Println("修改后：", buf[0])
}

func parameterBuf2(buf [3]int32) {
	if len(buf) == 0 {
		fmt.Println("buf size : 0")
		return
	}
	buf[0] = 245
	fmt.Println("修改：", buf[0])
}

// StringTest 字符串
func StringTest() {
	//字符串定义
	var str string = "hello, 世界!"
	// var str1 string
	//字符串常用方法
	//字符串遍历
	fmt.Println("range 遍历：")
	for i, v := range str {
		fmt.Printf("index: %d, q: %q, c: %c \n", i, v, v)
	}
	fmt.Println("utf.DecodeRuneInString 遍历：")
	fmt.Println("字符数：", utf8.RuneCountInString(str))
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("index: %d, q: %q, c: %c \n", i, r, r)
		i += size
	}
}

// SliceTest 切片
func SliceTest() {
	s := "a s d c 中 文 "
	var a []int
	fmt.Printf("a == nil : %t\n", a == nil)
	fmt.Printf("len(a) : %d\n", len(a))
	fmt.Printf("cap(a) : %d\n", cap(a))
	var strs []rune
	for _, r := range s {
		strs = append(strs, r)
	}
	out := strs[:0]
	for _, v := range strs {
		if v != ' ' {
			out = append(out, v)
		}
	}
	fmt.Println("out slice: ")
	for _, r := range out {
		fmt.Printf("%c ", r)
	}
	fmt.Println("strs slice: ")
	for _, r := range strs {
		fmt.Printf("%c ", r)
	}
}
func MapTest() {
	//创建map
	args1 := make(map[string]int32)
	args2 := map[string]int32{
		"alice":   1,
		"charlie": 2,
	}
	//增加
	args1["alice"] = 2
	args1["alice"]++

	//修改
	args2["alice"]++

	//删除
	delete(args2, "charlie")
	//访问 遍历
	fmt.Printf(string(args2["alie"]))
	for k, v := range args2 {
		fmt.Println(k, " ", v)
	}
}
