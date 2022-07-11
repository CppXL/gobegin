package main

import (
	"bytes"
	"fmt"
)

func main() {
	var a, b []byte
	a = []byte("aaabbb")
	b = []byte("bbb")
	// a<b return -1 , a>b return 1, a==b return 0
	// -1
	fmt.Println(bytes.Compare(a, b))
	// true
	fmt.Println(bytes.Contains(a, b))

	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "fÄo!"))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "去是伟大的."))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), ""))
	fmt.Println(bytes.ContainsAny([]byte(""), ""))

	// 计算非重复的切片数量
	a = []byte("aaabbbb")
	b = []byte("bbb")
	// 1
	fmt.Println(bytes.Count(a, b))
	a = []byte("aaabbbbbb")
	b = []byte("bbb")
	// 2
	fmt.Println(bytes.Count(a, b))
	a = []byte("aaabbbbbb")
	b = []byte("")
	// 10 len(a)+1
	fmt.Println(bytes.Count(a, b))

	// print aaa bbccc true
	a = []byte("aaabbbbbccc")
	b = []byte("bbb")
	c, d, result := bytes.Cut(a, b)
	fmt.Println(string(c), string(d), result)

	// true false
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++")))

	a = []byte("aaabbb")
	b = []byte("bbb")
	fmt.Println(bytes.EqualFold(a, b))

	// 返回utf-8码点 使用'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0 分割
	fmt.Println(bytes.Fields([]byte("  foo bar  baz  ")))

	a = []byte("aaabbbccc")
	b = []byte("aaa")
	c = []byte("ccc")
	fmt.Println(bytes.HasPrefix(a, b))
	fmt.Println(bytes.HasSuffix(a, c))
	// bytes.Index()
} // end main