// 最基础的语法：变量、数据类型、控制流、函数。
// 小白可以一行行看，每块注释都讲"这是啥、干啥用"。
// 跑一下：go run ./basics
package main

import "fmt" // fmt 负责把内容打到屏幕上

// 变量有几种写法，看一眼就懂
func demoVariables() {
	var name string = "小明" // 最正式：var + 类型
	var age = 18              // 让 Go 自己猜类型，这里会是 int
	height := 1.75            // := 最常用，函数里直接这么写，小数默认 float64

	fmt.Println("姓名：", name)
	fmt.Println("年龄：", age)
	fmt.Println("身高：", height)

	const pi = 3.14 // 常量，定了就改不了
	fmt.Println("圆周率：", pi)
}

// 看分数给句评价
func demoIfElse(score int) {
	if score >= 60 {
		fmt.Println("成绩", score, "分：及格啦！")
	} else {
		fmt.Println("成绩", score, "分：要加油哦～")
	}
}

// Go 只有 for 一种循环，没有 while
func demoFor() {
	for i := 1; i <= 5; i++ {
		fmt.Println("这是第", i, "次循环")
	}
}

// switch 多分支，比一长串 if 清爽
func demoSwitch(day int) {
	switch day {
	case 1:
		fmt.Println("星期一：开始新的一周")
	case 6, 7: // 多个值可以写一块
		fmt.Println("周末：好好休息")
	default:
		fmt.Println("工作日：认真写代码")
	}
}

// 两个数相加，参数后写类型，函数名后写返回值类型
func add(a int, b int) int {
	return a + b
}

func demoFunction() {
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)
}

func main() {
	demoVariables()
	demoIfElse(75)
	demoFor()
	demoSwitch(7)
	demoFunction()
}
