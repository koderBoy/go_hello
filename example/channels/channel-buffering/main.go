package main

import "fmt"

//默认情况下，通道是 无缓冲 的，这意味着只有对应的接收（<- chan） 通道准备好接收时，才允许进行发送（chan <-）。
//有缓冲通道 允许在没有对应接收者的情况下，缓存一定数量的值。
func main() {
	//这里我们 make 了一个字符串通道，最多允许缓存 2 个值。
	messages := make(chan string, 2)

	// 由于此通道是有缓冲的， 因此我们可以将这些值发送到通道中，而无需并发的接收。
	go func() {
		messages <- "buffered"
		messages <- "channel"
	}()

	// 然后我们可以正常接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
