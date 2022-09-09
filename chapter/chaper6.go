package chapter

import (
	"bytes"
	"fmt"
)

/* 方法
1、方法声明
2、指针接受者方法 指针与非指针应用区别
3、通过结构体内嵌组成类型
4、方法变量与表达式
5、位面量 一种解决特定问题的数据结构
6、封装
**/

type IntSet struct {
	words []int64
}

// Has 非负数检查
func (s *IntSet) Has(x int) bool {
	return true
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range s.words {
		if v == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if v&(1<<j) == 0 {
				continue
			}
			buf.WriteByte(' ')
			fmt.Fprintf(&buf, "%d", 64*i+j)
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Add 添加元素
func (s *IntSet) Add(x int) {
	w, index := x/64, x%64
	for w >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[w] |= 1 << index
}

// Remove 删除元素
func (s IntSet) Remove(x int) {
	w, index := x/64, x%64
	if w >= len(s.words) {
		return
	}
	var wor int64 = 0
	wor ^= 1 << index
	s.words[w] &= ^wor
}

// 获取元素数量
func (s IntSet) len() int {
	var num = 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if (word & (1 << i)) != 0 {
				num++
			}
		}
	}
	return num
}

// Clear 清空元素
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] = 0
	}
}

// Clear2 清空元素
func (s IntSet) Clear2() {
	for i, _ := range s.words {
		s.words[i] = 0
	}
}

// Copy 获取元素集合的副本
func (s *IntSet) Copy() *IntSet {
	var c IntSet
	for _, v := range s.words {
		c.words = append(c.words, v)
	}
	return &c
}
func MethodTest() {
	var c IntSet
	c1 := c.Add
	c1(2)
	c1(76)
	c1(455)
	c1(34)
	fmt.Printf("c 数据：%s\n", c)
	fmt.Printf("c 元素个数：%d\n", c.len())
	d := c.Copy()
	c2 := IntSet.Remove
	c2(c, 1)
	c2(*d, 2)
	fmt.Printf("remove c 数据：%s\n", c.String())
	fmt.Printf("remove d 数据：%s\n", d.String())

	d.Clear2()
	c.Clear()
	fmt.Printf("Clear c 数据：%s\n", c.String())
	fmt.Printf("Clear2 d 数据：%s\n", d.String())
}
