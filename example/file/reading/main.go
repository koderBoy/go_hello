//读写文件在很多程序中都是必须的基本任务。 首先我们来看一些读文件的例子
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//最基本的文件读取任务或许就是将文件内容读取到内存中。
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))

	//您通常会希望对文件的读取方式和内容进行更多控制。 对于这个任务，首先使用 Open 打开一个文件，以获取一个 os.File 值。
	f, err2 := os.Open("/tmp/dat")
	check(err2)
	defer f.Close()

	//从文件的开始位置读取一些字节。 最多允许读取 5 个字节，但还要注意实际读取了多少个。
	b1 := make([]byte, 5)
	n1, err3 := f.Read(b1)
	check(err3)
	fmt.Printf("%d bytes:%s\n", n1, string(b1[:n1]))

	//你也可以 Seek 到一个文件中已知的位置，并从这个位置开始读取。
	o1, err4 := f.Seek(6, 0)
	check(err4)
	b2 := make([]byte, 2)
	n, err5 := f.Read(b2)
	check(err5)
	fmt.Printf("%d bytes @%d:%s\n", n, o1, string(b2[:n]))

	//例如，io 包提供了一个更健壮的实现 ReadAtLeast，用于读取上面那种文件。
	o3, err6 := f.Seek(6, 0)
	check(err6)
	b3 := make([]byte, 2)
	n2, err7 := io.ReadAtLeast(f, b3, 2)
	check(err7)
	fmt.Printf("%d bytes @%d:%s\n", n2, o3, string(b3[:n2]))

	//没有内建的倒带，但Seek(0,0)实现了这一功能
	_, err8 := f.Seek(0, 0)
	check(err8)

	//bufio 包实现了一个缓冲读取器，这可能有助于提高许多小读操作的效率，以及它提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err9 := r4.Peek(5)
	check(err9)
	fmt.Printf("5 bytes:%s\n", string(b4))

}

//读取文件需要经常进行错误检查， 这个帮助方法可以精简下面的错误检查过程。
func check(err error) {
	if err != nil {
		panic(err)
	}

}
