package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。 对于更加复杂的情况，我们可以使用一个互斥锁 来在 Go 协程间安全的访问数据
func main() {

	var readOps, writeOps uint64

	//在这个例子中，state 是一个 map。
	var state = make(map[int]int)

	//mutex 将同步对 state 的访问。
	var mutex = &sync.Mutex{}

	//这里我们启动 100 个协程， 每个协程以每 1ms 一次的频率来重复读取 state。
	for r := 0; r < 100; r++ {
		go func() {
			//每次循环读取，我们使用一个键来进行访问， Lock() 这个 mutex 来确保对 state 的独占访问， 读取选定的键的值，Unlock() 这个 mutex，并将 ops 值加 1。
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				//在下次读取前等待片刻。
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//同样的，我们启动 10 个协程来模拟写入操作， 使用和读取相同的模式。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//让这 10 个协程对 state 和 mutex 的操作持续 1 s。
	time.Sleep(time.Second)

	fmt.Println("readOps:", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps:", atomic.LoadUint64(&writeOps))

	//对 state 使用一个最终的锁，展示它是如何结束的。
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

}
