//我们经常需要程序对数据集合执行操作， 例如选择满足给定条件的全部 item， 或通过自定义函数将全部 item 映射到一个新的集合。
// 在其它语言中，通常会使用泛型数据结构和算法。 但 Go 不支持泛型，如果你的程序或者数据类型有需要，通常的做法是提供函数集。
// 这是一些 strings 切片的组合函数示例。 你可以使用这些例子来构建自己的函数。
// 注意，在某些情况下，最简单明了的方法是： 直接内联操作方法集，而不是创建并调用帮助函数。
package main

import (
	"fmt"
	"strings"
)

//Index 返回目标字符串 t 在 vs 中第一次出现位置的索引， 或者在没有匹配值时返回 -1。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include 如果目标字符串 t 存在于切片 vs 中，则返回 true。
func Include(vs []string, t string) bool {
	return Index(vs, t) > -1
}

//Any 如果切片 vs 中的任意一个字符串满足条件 f，则返回 true。
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//All 如果切片 vs 中的所有字符串都满足条件 f，则返回 true。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//Filter 返回一个新的切片，新切片由原切片 vs 中满足条件 f 的字符串构成。
func Filter(vs []string, f func(s string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//Map 返回一个新的切片，新切片的字符串由原切片 vs 中的字符串经过函数 f 映射后得到。
func Map(vs []string, f func(s string) string) []string {
	vsf := make([]string, len(vs))
	for i, v := range vs {
		vsf[i] = f(v)
	}
	return vsf
}

//试试各种组合函数。
func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

}
