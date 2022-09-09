package chapter

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

/*
*接口
1、接口类型
2、实现接口 隐式实现
3、flag.Value 接受终端传入的参数
4、接口值 动态类型，动态值
5、sort.Interface 排序的接口
6、http.Handler
7、error
8、表达式求值器
9、类型断言 类似java中的父类对象强转子类对象

	识别错误 查询特性

10、类型分支
11、xml解析
*/
type celsius float64

const debug = false

// 自定义实现Flag 实现Flag.value接口
type celsiusFlag struct {
	celsius
}
type Track struct {
	title  string
	artist string
	album  string
	year   int
	length time.Duration
}
type TrackSort struct {
	tacks []*Track
	less  func(x, y *Track) bool
}

func (t TrackSort) Len() int {
	return len(t.tacks)
}
func (t TrackSort) Less(i, j int) bool {
	return t.less(t.tacks[i], t.tacks[j])
}
func (t TrackSort) Swap(i, j int) {
	t.tacks[i], t.tacks[j] = t.tacks[j], t.tacks[i]
}

func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) //Sscanf 拆分字符串
	switch unit {
	case "C", "c", "ºC", "ºc":
		c.celsius = celsius(value)
		return nil
	case "F", "f", "ºF", "ºf":
		c.celsius = fToC(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
func (c celsius) String() string {
	s := fmt.Sprintf("%.1fºC", c)
	return s
}
func CelsiusFlag(name string, value celsius, usage string) *celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.celsius
}
func fToC(v float64) celsius {
	v = (v - 32) * 5 / 9
	return celsius(v)
}

// TestInterfaceNil 空指针的非空接口
func TestInterfaceNil() {
	//var buf *bytes.Buffer //非nil动态类型
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) //非nil的动态值
	}
	f(buf)
}
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done"))
	}
}

// TestSort 自定义排序
func TestSort() {
	var tracks = []*Track{{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")}}
	fmt.Println("排序前：")
	printTracks(tracks)
	//排序
	trackSort := TrackSort{tracks, func(x, y *Track) bool {
		if x.title != y.title {
			return x.title > y.title
		}
		if x.year != y.year {
			return x.year > y.year
		}
		return x.artist > y.artist
	}}
	sort.Sort(trackSort)
	fmt.Println("排序后：")
	printTracks(tracks)
	fmt.Println("反向排序后：")
	sort.Sort(sort.Reverse(trackSort))
	printTracks(tracks)
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "title", "artist", "album", "year", "length")
	fmt.Fprintf(tw, format, "_____", "_____", "_____", "_____", "_____")
	for _, track := range tracks {
		fmt.Fprintf(tw, format, track.title, track.artist, track.album, track.year, track.length)
	}
	tw.Flush()
}
