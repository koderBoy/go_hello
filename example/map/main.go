package main

import "fmt"

//map 是 Go 内建的关联数据类型 （在一些其他的语言中也被称为 哈希(hash) 或者 字典(dict) ）。
// 注意，使用 fmt.Println 打印一个 map 的时候， 是以 map[k:v k:v] 的格式输出的。
func main() {

	//要创建一个空 map，需要使用内建函数 make：make(map[key-type]val-type)。
	m := make(map[string]string)

	// 使用典型的 name[key] = val 语法来设置键值对。
	m["name"] = "change fei"
	m["age"] = "30"

	// 打印 map。例如，使用 fmt.Println 打印一个 map，会输出它所有的键值对。
	fmt.Println("m=", m)

	// 使用 name[key] 来获取一个键的值。
	fmt.Println("name", m["name"])

	// 内建函数 len 可以返回一个 map 的键值对数量。
	fmt.Println("len:", len(m))

	// 内建函数 delete 可以从一个 map 中移除键值对。
	delete(m, "name")
	fmt.Println("map", m)

	// 当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。
	// 这可以用来消除 键不存在 和 键的值为零值 产生的歧义， 例如 0 和 ""。
	// 这里我们不需要值，所以用 空白标识符(blank identifier) _ 将其忽略。
	_, psr := m["changfei"]
	fmt.Println("psr:", psr)

}
