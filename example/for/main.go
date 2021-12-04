package main

//for 是 Go 中唯一的循环结构。这里会展示 for 循环的三种基本使用方式。
func main() {

	//最基础方式，单个条件for循环
	i := 1
	for i <= 3 {
		println(i)
		i++
	}

	//经典的初始/条件/后续 for 循环。

	for j := 4; j <= 10; j++ {
		println(j)
	}

	// 不带条件的 for 循环将一直重复执行， 直到在循环体内使用了 break 或者 return 跳出循环。
	for {
		println("loop")
		break
	}

	//你也可以使用 continue 直接进入下一次循环。
	for n := 5; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		println(n)
	}
}
