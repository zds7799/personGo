package main

import (
	"sample-app/chapter"
	"sample-app/test"
)

const path string = "../xml配置"

func main() {
	test.Helloword()
	chapter.BufTest()
	chapter.StringTest()
	chapter.SliceTest()
	// test.TestString()
	//fmt.Println(test.CtoF(test.Celsius(123)))
	// fmt.Println(test.SameString("go程序", "go程序"))
	// chapter.BuffTest()
	// chapter.SliceTest()
}
