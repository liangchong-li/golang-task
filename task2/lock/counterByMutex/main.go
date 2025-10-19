package main

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

var (
	counter int = 0
	lock sync.Mutex
	wg   sync.WaitGroup // 添加 WaitGroup
)

func increment() {
	lock.Lock()
	counter++
	lock.Unlock()
}

func main() {
	fmt.Println("befor opt", counter)
	
	// 设置 WaitGroup 的计数器为 1000
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done() // 确保在函数退出时调用 Done
			for i := 0; i < 1000; i++ {
				increment()
			}
		}()
	}

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("after opt", counter)
}

