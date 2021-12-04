package main

import (
	"fmt"
)

//zeroval 有一个 int 型参数，所以使用值传递。 zeroval 将从调用它的那个函数中得到一个实参的拷贝：i。
func zeroval(i int) {
	i = 0
}

//zeroptr 有一个和上面不同的参数：*int，这意味着它使用了 int 指针。
// 紧接着，函数体内的 *iptr 会 解引用 这个指针，从它的内存地址得到这个地址当前对应的值。
//  对解引用的指针赋值，会改变这个指针引用的真实地址的值。
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	//字符串指针,&变量名的方式获取变量指向的内存地址
	str := "hello world"
	fmt.Printf("str: %v\n", str)
	fmt.Printf("str 的地址:%v\n", &str)

	//定义并用&str初始化一个字符串指针
	pstr := &str
	fmt.Printf("pstr的类型是 %T\n", pstr)
	fmt.Printf("pstr: %v\n", pstr)

	var num int = 10

	//定义并初始化一个int类型的指针
	var pnum = &num
	fmt.Println(pnum)

	//用*指针变量来获取指针指向的数据
	fmt.Printf("num*100=%d\n", *pnum*100)

	i := 1
	//通过zerval并不能将 i 的值修改为0
	zeroval(i)
	fmt.Println(i)

	// 通过zeroptr代用可以修改变量i的值
	zeroptr(&i)
	fmt.Println(i)
}
