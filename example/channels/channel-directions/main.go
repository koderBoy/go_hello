package main

import "fmt"

//当使用通道作为函数的参数时，你可以指定这个通道是否为只读或只写。 该特性可以提升程序的类型安全。

//ping 函数定义了一个只能发送数据的（只写）通道。 尝试从这个通道接收数据会是一个编译时错误。
func ping(ping chan<- string, msg string) {
	ping <- msg
}

//pong 函数接收两个通道，pings 仅用于接收数据（只读），pongs 仅用于发送数据（只写）。
func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	go ping(pings, "passed message")
	go pong(pings, pongs)
	fmt.Println(<-pongs)
}
