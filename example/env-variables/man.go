// 环境变量 是一种向 Unix 程序传递配置信息的常见方式。 让我们来看看如何设置、获取以及列出环境变量。
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//使用 os.Setenv 来设置一个键值对。 使用 os.Getenv获取一个键对应的值。 如果键不存在，将会返回一个空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	//使用 os.Environ 来列出所有环境变量键值对。 这个函数会返回一个 KEY=value 形式的字符串切片。
	// 你可以使用 strings.SplitN 来得到键和值。这里我们打印所有的键。
	fmt.Println()

	for _, e := range os.Environ() {
		pari := strings.SplitN(e, "=", 2)
		fmt.Println(pari[0])
	}

}
