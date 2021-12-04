package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	//开始！这里展示了如何写入一个字符串（或者只是一些字节）到一个文件。
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat2", d1, 0664)
	check(err)

	//对于更细粒度的写入，先打开一个文件。
	f, err2 := os.Create("/tmp/dat2")
	check(err2)

	//打开文件后，一个习惯性的操作是：立即使用 defer 调用文件的 Close。
	defer f.Close()

	//您可以按期望的那样 Write 字节切片(slice)
	d2 := []byte{115, 111, 109, 101, 10}
	n, err3 := f.Write(d2)
	check(err3)
	fmt.Printf("wrote %d bytes\n", n)

	//WriteString 也是可用的。
	n, err4 := f.WriteString("writes\n")
	check(err4)
	fmt.Printf("wrote %d n bytes\n", n)

	//调用 Sync 将缓冲区的数据写入硬盘。
	f.Sync()

	//与我们前面看到的带缓冲的 Reader 一样，bufio 还提供了的带缓冲的 Writer。
	w := bufio.NewWriter(f)
	i, err5 := w.WriteString("buffered\n")
	check(err5)
	fmt.Printf("wrote %d bytes\n", i)

	//使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer。
	w.Flush()

}
