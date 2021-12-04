package main

import "fmt"

// 定义一个person结构体,这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

func main() {

	//使用这个语法创建新的结构体元素。
	fmt.Println(person{"changfe", 18}.name)

	// 你可以在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{age: 18, name: "常飞"})

	//省略的字段将被初始化为零值。
	fmt.Println(person{name: "Fred"}.age)

	//& 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40})

	//使用.来访问结构体字段。
	p := person{age: 50}
	fmt.Println(p.name)

	// 也可以对结构体指针使用. - 指针会被自动解引用。结构体是可变(mutable)的。
	sp := &p
	sp.age = 51
	fmt.Println(p.age)

}
