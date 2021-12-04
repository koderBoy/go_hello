package main

import "fmt"

//
func main() {
	nums := []int{1, 2, 3, 4, 5}

	// 这里我们使用 range 来对 slice 中的元素求和。 数组也可以用这种方法初始化并赋初值。
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum=", sum)

	//range 在数组和 slice 中提供对每项的索引和值的访问。 上面我们不需要索引，所以我们使用 空白标识符 _ 将其忽略。
	//  实际上，我们有时候是需要这个索引的。
	for i, v := range nums {
		if i%2 == 0 {
			println("index: %d value:%", i, v)
		}
	}

	// range 也可以只遍历 map 的键。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range 在字符串中迭代 unicode 码点(code point)。 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
	for ops, c := range "我们都有一个家，名字叫中国" {
		fmt.Printf("%d:%v:%c\n", ops, c, c)
	}

}
