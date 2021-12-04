package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//下面我们将使用这两个结构体来演示自定义类型的编码和解码。
//只有 可导出 的字段才会被 JSON 编码/解码。必须以大写字母开头的字段才是可导出的。
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	//首先我们来看一下基本数据类型到 JSON 字符串的编码过程。 这是一些原子值的例子。
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.35)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//这是一些切片和 map 编码成 JSON 数组和对象的例子。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "peach": 3}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//JSON 包可以自动的编码你的自定义类型。 编码的输出只包含可导出的字段，并且默认使用字段名作为 JSON 数据的键名。
	resp1D := response1{1, []string{"apple", "peach", "pear"}}
	resp1B, _ := json.Marshal(&resp1D)
	fmt.Println(string(resp1B))

	//你可以给结构字段声明标签来自定义编码的 JSON 数据的键名。 上面 Response2 的定义，就是这种标签的一个例子。
	resp2D := response2{1, []string{"apple", "peach", "pear"}}
	resp2B, _ := json.Marshal(&resp2D)
	fmt.Println(string(resp2B))

	//现在来看看将 JSON 数据解码为 Go 值的过程。 这是一个普通数据结构的解码例子。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	//我们需要提供一个 JSON 包可以存放解码数据的变量。 这里的 map[string]interface{} 是一个键为 string，值为任意值的 map。
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	//访问嵌套的值需要一系列的转化。
	numI := dat["num"]
	num := numI.(float64) //断言
	println(num)

	strs := dat["strs"].([]interface{}) //断言
	str1 := strs[0].(string)
	println(str1)

	//我们还可以将 JSON 解码为自定义数据类型。 这样做的好处是，可以为我们的程序增加附加的类型安全性， 并在访问解码后的数据时不需要类型断言。
	str := `{"page":1,"fruits":["apple","peach","pear"]}`
	res := response2{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}
	fmt.Println(res.Fruits)

	//在上面例子的标准输出上， 我们总是使用 byte和 string 作为数据和 JSON 表示形式之间的中介。
	//当然，我们也可以像 os.Stdout 一样直接将 JSON 编码流传输到 os.Writer 甚至 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	if err := enc.Encode(d); err != nil {
		panic(err)
	}

}
