//URL 提供了统一资源定位方式。 这里展示了在 Go 中是如何解析 URL 的。
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	//我们将解析这个 URL 示例，它包含了一个 scheme、 认证信息、主机名、端口、路径、查询参数以及查询片段
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	//直接访问 scheme。
	s2 := u.Scheme
	fmt.Printf("schema: %v\n", s2)

	//User 包含了所有的认证信息， 这里调用 Username 和 Password 来获取单独的值。
	u2 := u.User
	fmt.Printf("user: %v\n", u2.Username())
	password, _ := u2.Password()
	fmt.Printf("password: %v\n", password)

	//Host 包含主机名和端口号（如果存在）。使用 SplitHostPort 提取它们。
	s3 := u.Host
	fmt.Printf("host: %v\n", s3)
	host, port, _ := net.SplitHostPort(s3)
	fmt.Printf("host: %v\n", host)
	fmt.Printf("port: %v\n", port)

	//这里我们提取路径和 # 后面的查询片段信息。
	s4 := u.Path
	fmt.Printf("path: %v\n", s4)

	s5 := u.Fragment
	fmt.Printf("fragment: %v\n", s5)

	//要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。
	//你也可以将查询参数解析为一个 map。
	//已解析的查询参数 map 以查询字符串为键， 已解析的查询参数会从字符串映射到到字符串的切片， 因此如果您只想要第一个值，则索引为 [0]。
	s6 := u.RawQuery
	fmt.Printf("query: %v\n", s6)

	values := u.Query()

	for k, v := range values {
		fmt.Printf("%v=%v", k, v)
	}

}
