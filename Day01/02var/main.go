package main

import "fmt"

//批量声明
var (
	name string
	age  int
	isok bool
)

//声明的变量之后必须使用，不然编译不过去
func main() {
	name = "wang"
	age = 16
	isok = true

	fmt.Print(isok)
	fmt.Printf("name:%s\n", name)
	fmt.Println(age)

	//声明变量同时赋予
	var s1 string = "zhang"
	fmt.Println(s1)
	//类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)

	//简短变量声明,只能在函数里面使用
	s3 := "hahaha"
	fmt.Println(s3)

}
