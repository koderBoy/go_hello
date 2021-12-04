package main

import "fmt"

func main() {
	//go 声明（定义）变量的三种方式

	//第一种 是用类型推断
	var str = "hello world"
	fmt.Printf("str: %v\n", str)

	//第二种方式是用显示类型声明
	var str2 string = "hello world"
	fmt.Printf("str2: %v\n", str2)

	//第三种方式可以省略 var 等价于 var str3 str3="hello world"
	str3 := "hello world"
	fmt.Printf("str3: %v\n", str3)
}
