package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	goroutine1()

	var tasks []func(a, b int) int = []func(a, b int) int{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a - b },
		func(a, b int) int { return a * b },
		func(a, b int) int { return a / b },
	}
	goroutine2(6, 4, tasks)
}


// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func goroutine1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i := 1; i < 10 ; i +=2 {
			fmt.Println("goroutine a print:", i)
		}
	}()

	wg.Add(1)
	go func(){
		defer wg.Done()
		for i := 2; i <= 10 ; i +=2 {
			fmt.Println("goroutine b print:", i)
		}
	}()

	wg.Wait()
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，
// 同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
func goroutine2(a int, b int, tasks []func(a, b int) int) {
	var wg sync.WaitGroup
	fmt.Println("进入任务调度器")
	for i := 0; i < len(tasks); i++ {
		// go tasks[i](6, 4)
		wg.Add(1)
		fmt.Println("遍历执行任务")
		go func() {
			start := time.Now()
			defer wg.Done()
			res := tasks[i](b, a)
			fmt.Println("任务", i, "执行结果：", res, "耗时：", time.Since(start))
		}()
		wg.Wait()
	}
}