package main

import (
	"fmt"
	"time"
)

//我们经常需要在未来的某个时间点运行 Go 代码，或者每隔一定时间重复运行代码。
//Go 内置的 定时器 和 打点器 特性让这些变得很简单。 我们会先学习定时器，然后再学习打点器。

func main() {

	// 定时器表示在未来某一时刻的独立事件。 你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。 这里的定时器将等待 2 秒。
	t1 := time.NewTicker(2 * time.Second)

	// <-t1.C 会一直阻塞， 直到定时器的通道 C 明确的发送了定时器失效的值。
	<-t1.C
	fmt.Println("Timer t1 fired")

	//如果你需要的仅仅是单纯的等待，使用 time.Sleep 就够了。
	//使用定时器的原因之一就是，你可以在定时器触发之前将其取消。 例如这样
	t2 := time.NewTimer(time.Second)
	go func() {
		<-t2.C
		fmt.Println("Timer t2 fired")
	}()

	if t2.Stop() {
		fmt.Println("Timer t2 stopped")
	}
	//给 timer2 足够的时间来触发它，以证明它实际上已经停止了。
	time.Sleep(2 * time.Second)

}
