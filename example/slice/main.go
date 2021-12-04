package main

import "fmt"

//Slice 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式。

func main() {
	s := make([]string, 5)
	fmt.Println("empt s:", s)

	// ßlen 返回 slice 的长度
	fmt.Println("len:", len(s))

	//我们可以和数组一样设置和得到值
	s[0] = "jack"
	s[1] = "ma"
	s[2] = "poniy"
	s[3] = "dong"
	fmt.Println("s[3]=", s[3])
	fmt.Println("s=", s)

	//除了基本操作外，slice 支持比数组更丰富的操作。比如 slice 支持内建函数 append，
	//该函数会返回一个包含了一个或者多个新值的 slice。 注意由于 append 可能返回一个新的 slice，我们需要接收其返回值。
	s = append(s, "hello")
	s = append(s, "jimi", "meimei")
	fmt.Println("s=", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c=", c)

	//slice 支持通过 slice[low:high] 语法进行“切片”操作。 例如，右边的操作可以得到一个包含元素 s[2]、s[3] 和 s[4] 的 slice。
	l := s[2:5]
	fmt.Println("sl1:", l)

	//这个 slice 包含从 s[0] 到 s[5]（不包含 5）的元素。
	l = s[:5]
	fmt.Println("sl2:", l)

	//这个 slice 包含从 s[2]（包含 2）之后的元素。
	l = s[2:]
	fmt.Println("sl3:", l)

	//我们可以通过c[0:0]获取一个空的slice并赋值给c方式来清空slice
	c = c[0:0]

	//我们可以在一行代码中声明并初始化一个 slice 变量。
	t := []string{"hi", "say", "hello"}
	fmt.Printf("t: %v\n", t)

	fmt.Println("c=", c)
	fmt.Println("s=", s)

}
