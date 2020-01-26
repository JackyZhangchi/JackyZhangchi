package main

import "fmt"

//if条件判断
func main() {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("网吧开黑")
	// } else {
	// 	fmt.Println("寒假作业")
	// }

	//多个判断条件
	// if age > 35 {
	// 	fmt.Println("人到中年")
	// } else if age > 18 {
	// 	fmt.Println("青年")
	// } else if age < 18 {
	// 	fmt.Println("寒假作业")
	// }

	//作用域
	//age变量此时只在IF判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("网吧开黑")
	} else {
		fmt.Println("寒假作业")
	}

	// fmt.Println(age) //在这里是找不到age
}
