package main

import "fmt"

//fact 函数在到达 fact(0) 前一直调用自身。
func fact(n, s int) int {
	if n <= 0 {
		return 1
	} else if n == 1 {
		return s
	} else {
		return fact(n-1, s*n)
	}
}

func main() {
	f := fact(5, 1)
	fmt.Printf("f: %v\n", f)
}
