// time 包几乎是每个 Go 程序都会碰到的：拿当前时间、算时间差、格式化输出。
// 这里挑几个最实用的玩法演示一遍。
package main

import (
	"fmt"
	"time"
)

// 拿当前时间，再算"几秒后"和"格式化"
func demoNow() {
	now := time.Now()
	fmt.Println("现在：", now.Format("2006-01-02 15:04:05"))

	// Go 的参考时间就是 2006-01-02 15:04:05，记这个模板而不是别的年份
	later := now.Add(90 * time.Second)
	fmt.Println("90 秒后：", later.Format("15:04:05"))

	// 算两个时间差，直接减就得到 Duration
	diff := later.Sub(now)
	fmt.Println("差了：", diff.Seconds(), "秒")
}

// 定时器：睡一会儿再继续，模拟"等 1 秒"
func demoSleep() {
	start := time.Now()
	time.Sleep(1 * time.Second) // 真会停 1 秒，别在循环里乱用
	fmt.Printf("睡了 %.0f 秒\n", time.Since(start).Seconds())
}

// 把字符串按格式解析回时间，常用于读日志/配置文件里的日期
func demoParse() {
	layout := "2006-01-02"
	t, err := time.Parse(layout, "2026-08-06")
	if err != nil {
		fmt.Println("解析失败：", err)
		return
	}
	fmt.Println("解析出来的星期：", t.Weekday())
}

func main() {
	demoNow()
	demoSleep()
	demoParse()
}
