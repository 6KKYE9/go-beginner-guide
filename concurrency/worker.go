// 示例 5（扩展）：并发编程入门
//
// Go 最吸引人的特性之一就是「并发」——用很少的代码同时做多件事。
// 本示例带你认识三个核心概念：
//  1. goroutine  —— 用 go 关键字启动的“轻量级线程”，廉价到可以同时开成千上万个。
//  2. channel    —— goroutine 之间传递数据的“管道”，用 <- 收发。
//  3. sync.WaitGroup —— 一种“等所有人干完活再下班”的计数工具。
//
// 建议先读懂 basics 和 structs 两个示例再来看本文件。
package main

import (
	"fmt"
	"sync"
	"time"
)

// worker 模拟一个“打工人”：从 jobs 通道里取任务，处理完把结果放进 results。
// 每处理一个任务就 sleep 一下，假装在工作。
// 注意参数 wg *sync.WaitGroup：所有 worker 共享同一个“签到表”，干完活要签到(wg.Done)。
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	// 函数退出前必须调用 wg.Done()，告诉 WaitGroup“我又少了一个活要等”。
	// defer 保证即使函数中途 return 也会执行，非常安全。
	defer wg.Done()

	for job := range jobs { // 不断从通道取任务，通道关闭后循环自动结束
		fmt.Printf("工人 %d 开始处理任务 %d\n", id, job)
		time.Sleep(time.Millisecond * 200) // 模拟耗时工作
		results <- job * 2                 // 把“结果”放进结果通道
	}
}

// RunConcurrency 是示例的主入口，供 README 与测试调用。
// 它启动 3 个 worker 并发处理 5 个任务，并返回所有结果（已排序无关，这里按完成顺序收集）。
func RunConcurrency() []int {
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int, numJobs)    // 缓冲通道，容量 5，可一次性放入 5 个任务
	results := make(chan int, numJobs) // 结果通道，容量 5
	var wg sync.WaitGroup              // 签到表：记录还有几个 worker 没下班

	// 启动 3 个 worker，每个都在后台 goroutine 里跑
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1) // 每启动一个 worker，签到表 +1
		go worker(w, jobs, results, &wg)
	}

	// 发送 5 个任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // 任务发完了，关闭通道，worker 取完后循环会退出

	// 单独开一个 goroutine 等所有 worker 下班，然后关闭 results 通道。
	// 否则下面“从 results 取数据”的循环会一直等，程序卡死。
	go func() {
		wg.Wait()      // 阻塞，直到所有 worker 都 wg.Done()
		close(results) // 全员下班，结果通道可以关闭了
	}()

	// 主 goroutine 从 results 收集所有结果
	var got []int
	for r := range results {
		got = append(got, r)
	}
	return got
}

func main() {
	fmt.Println("===== 并发入门：3 个工人处理 5 个任务 =====")
	out := RunConcurrency()
	fmt.Println("全部结果（按完成顺序）:", out)
}
