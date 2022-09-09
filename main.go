package main

import (
	"flag"
	"log"
	"os"
	"sample-app/chapter"
	"time"
)

const path string = "../xml配置"

var period = flag.Duration("period", 1*time.Second, "sleep period")
var celsius = chapter.CelsiusFlag("celsius", 20.0, "celsius")

func main() {
	//test.Helloword()
	//fmt.Println()
	//log.Fatal(string([]byte("s")))
	//chapter.BufTest()
	//chapter.StringTest()
	//chapter.SliceTest()
	//chapter.TestParameter()
	//chapter.TopSort()
	//chapter.MethodTest()
	// test.TestString()
	//fmt.Println(test.CtoF(test.Celsius(123)))
	// fmt.Println(test.SameString("go程序", "go程序"))
	// chapter.BuffTest()
	// chapter.SliceTest()
	/**chaper7*/
	//flag.Parse()
	//fmt.Println("sleep time ", *period)
	//time.Sleep(*period)
	//
	//fmt.Println("celsius ", *celsius)
	//chapter.TestInterfaceNil()
	//chapter.TestSort()
	/**chaper8*/
	//chapter8.TestComplicated()
	//chapter8.CrewUrl([]string{"https://www.baidu.com/"})
	//chapter8.TestTimer2()
	create, err := os.Create("aa.txt")
	if err != nil {
		return
	}
	log.SetOutput(create)
	log.Println("aaa")
}
