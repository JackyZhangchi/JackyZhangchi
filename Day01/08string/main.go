package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "\"D:\\Gopath\\src\\github.com\\JackyZhangchi\\day01\""
	fmt.Println(path)
	path1 := "'D:\\Gopath\\src\\github.com\\JackyZhangchi\\day01'"
	fmt.Println(path1)

	s := "i'm ok"
	fmt.Println(s)

	//多行字符串
	s2 := `
世情薄
人情恶
雨送黄昏花易落
	`
	fmt.Println(s2)
	s3 := `D:\Gopath\src\github.com\JackyZhangchi\Day01`
	fmt.Println(s3)

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "wang"
	word := "dsb"

	ss := name + word
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, word) //把拼接的字符串重新变成函数
	fmt.Println(ss1)
	//分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(ss, "wang"))
	fmt.Println(strings.Contains(ss, "li"))
	//前缀和后缀判断
	fmt.Println(strings.HasPrefix(ss, "wang"))
	fmt.Println(strings.HasSuffix(ss, "wang"))

	s4 := "abcde"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.Index(s4, "e"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
