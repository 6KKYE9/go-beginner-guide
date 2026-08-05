// 本文件演示 Go 语言最基础的语法：变量、数据类型、控制流、函数。
// 适合编程小白逐行阅读，每一块都有注释说明“这是什么、有什么用”。
//
// 在 Go 中：
//   - 每个可独立运行的程序都必须有一个名为 main 的包（package main）
//   - 程序入口是 main() 函数（没有参数、没有返回值）
package main

import "fmt" // fmt 是 Go 内置的“格式化输入输出”标准库，用来打印内容到屏幕

// ============ 1. 变量与数据类型 ============

// 演示变量的三种声明方式。
func demoVariables() {
	// 方式一：使用 var 关键字声明，并指定类型
	var name string = "小明" // string 表示“字符串”（一段文字）
	// 方式二：声明时让 Go 自动推断类型（不写类型，Go 根据你给的值判断）
	var age = 18 // 这里 Go 会自动判断 age 是 int（整数）类型
	// 方式三：简短声明 := （只能在函数内部使用，最常用、最简洁）
	height := 1.75 // 小数默认是 float64 类型（双精度浮点数）

	fmt.Println("===== 变量演示 =====")
	fmt.Println("姓名：", name)
	fmt.Println("年龄：", age)
	fmt.Println("身高：", height)

	// 常量：一旦赋值就不能再修改，用 const 声明
	const pi = 3.14
	fmt.Println("圆周率：", pi)
}

// ============ 2. 控制流（if / for / switch） ============

// 演示条件判断 if-else。
func demoIfElse(score int) {
	fmt.Println("===== if-else 演示 =====")
	if score >= 60 {
		fmt.Println("成绩", score, "分：及格啦！")
	} else {
		fmt.Println("成绩", score, "分：要加油哦～")
	}
}

// 演示最常用的循环 for（Go 只有 for，没有 while）。
func demoFor() {
	fmt.Println("===== for 循环演示 =====")
	// 经典三段式：初始值; 循环条件; 每轮结束后执行
	for i := 1; i <= 5; i++ {
		fmt.Println("这是第", i, "次循环")
	}
}

// 演示 switch 多分支选择。
func demoSwitch(day int) {
	fmt.Println("===== switch 演示 =====")
	switch day {
	case 1:
		fmt.Println("星期一：开始新的一周")
	case 6, 7: // 多个值可以写在一起
		fmt.Println("周末：好好休息")
	default:
		fmt.Println("工作日：认真写代码")
	}
}

// ============ 3. 函数 ============

// 定义一个“两数相加”的函数。
// 参数后面写类型，函数名后面写“返回值类型”。
func add(a int, b int) int {
	return a + b
}

// 演示函数调用。
func demoFunction() {
	fmt.Println("===== 函数演示 =====")
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)
}

// ============ 程序入口 ============

// main 是程序的入口，Go 运行时会自动从这里开始执行。
func main() {
	demoVariables()
	demoIfElse(75)
	demoFor()
	demoSwitch(7)
	demoFunction()
}
