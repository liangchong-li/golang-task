package main

import (
	"fmt"
	"time"
)

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

func main() {
	ch := make(chan int, 10)
	go productor(ch)
	go consumer(ch)
	
	timeout := time.After(2 * time.Second)
	for {
        select {
        case v, ok := <-ch:
            if !ok {
                fmt.Println("Channel已关闭")
                return
            }
            fmt.Printf("主goroutine接收到: %d\n", v)
        case <-timeout:
            fmt.Println("操作超时")
            return
        default:
            fmt.Println("没有数据，等待中...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

// 生产者
func productor(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}


// 消费者
func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Printf("消费%d\n", v)
	}
}