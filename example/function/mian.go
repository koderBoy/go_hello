package main

import "fmt"

//

func main() {

	// 通过 函数名(参数列表) 来调用函数
	res := plus(1, 2)
	fmt.Printf("res: %v\n", res)

	res = plusPlus(1, 2, 3)
	fmt.Printf("res: %v\n", res)

	//一次函数调用返回两个输的差、和
	sum, sub := calcl(1, 2)
	fmt.Printf("sub: %v\n", sub)
	fmt.Printf("sum: %v\n", sum)

}

// 这里是一个函数，接受两个 int 并且以 int 返回它们的和
func plus(a int, b int) int {
	//Go 需要明确的 return，也就是说，它不会自动 return 最后一个表达式的值
	return a + b
}

// 当多个连续的参数为同样类型时， 可以仅声明最后一个参数的类型，忽略之前相同类型参数的类型声明。
func plusPlus(a, b, c int) int {
	return a + b + c
}

//Go 原生支持 _多返回值_。 这个特性在 Go 语言中经常用到，例如用来同时返回一个函数的结果和错误信息。
func calcl(a, b int) (int, int) {
	return a + b, a - b
}
