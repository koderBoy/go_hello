package main

import "fmt"

//关闭 一个通道意味着不能再向这个通道发送值了。 该特性可以向通道的接收方传达工作已经完成的信息。

//在这个例子中，我们将使用一个 jobs 通道，将工作内容， 从 main() 协程传递到一个工作协程中。
//当我们没有更多的任务传递给工作协程时，我们将 close 这个 jobs 通道。
func main() {
	jobs := make(chan int, 5)
	down := make(chan bool)

	//这是工作协程。使用 j, more := <- jobs 循环的从 jobs 接收数据。
	// 根据接收的第二个值，如果 jobs 已经关闭了， 并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。
	//当我们完成所有的任务时，会使用这个特性通过 done 通道通知 main 协程。
	go func() {
		for {
			if i, more := <-jobs; more {
				fmt.Println("received job:", i)
			} else {
				fmt.Print("received all jobs")
				down <- true
			}
		}
	}()

	//使用 jobs 发送 3 个任务到工作协程中，然后关闭 jobs。
	for i := 0; i < 3; i++ {
		jobs <- i
	}
	close(jobs)

	//使用前面学到的通道同步方法等待任务结束。
	<-down
}
