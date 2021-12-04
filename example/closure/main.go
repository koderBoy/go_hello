package main

import "fmt"

//intSeq 函数返回一个在其函数体内定义的匿名函数。 返回的函数使用闭包的方式 隐藏 变量 i。 返回的函数 隐藏 变量 i 以形成闭包。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt。 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时，都会更新 i 的值。
	nextInt := intSeq()

	//通过多次调用 nextInt 来看看闭包的效果。
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
}
