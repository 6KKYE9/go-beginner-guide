// 本文件演示 Go 语言常用的“标准库”（官方自带、无需额外安装的功能包）。
// 重点演示三个小白最常用的包：
//   - fmt：打印输出
//   - os ：读取命令行参数、与环境交互
//   - strconv：字符串与数字之间的相互转换
package main

import (
	"fmt"
	"os"
	"strconv"
)

// 演示 os 包：读取运行程序时传入的命令行参数。
// 例如：go run stdlib.go 你好 18
// 其中 "你好" 和 "18" 就是命令行参数。
func demoOS() {
	fmt.Println("===== os 包：命令行参数演示 =====")
	// os.Args 是一个“字符串切片”（可以理解为字符串组成的列表）
	// os.Args[0] 是程序自身路径，真正传给我们的参数从下标 1 开始
	args := os.Args
	if len(args) < 3 {
		fmt.Println("提示：请这样运行 -> go run stdlib.go 你的名字 你的年龄")
		return
	}
	name := args[1]        // 第一个参数：名字
	ageStr := args[2]      // 第二个参数：年龄（此时还是字符串）
	fmt.Println("你好，", name, "！你输入的年龄字符串是：", ageStr)
}

// 演示 strconv 包：把字符串转换成整数。
func demoStrconv() {
	fmt.Println("===== strconv 包：字符串转数字演示 =====")
	ageStr := "18"
	// strconv.Atoi 把字符串转成 int，返回两个值：结果 和 错误
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		// Go 用“错误值”来表示出错了，nil 表示“没有错误”
		fmt.Println("转换失败：", err)
		return
	}
	// 转换成功后就可以做数学运算了
	fmt.Println("明年你就", age+1, "岁了")
}

// 演示 fmt 的多种格式化输出方式。
func demoFmt() {
	fmt.Println("===== fmt 包：格式化输出演示 =====")
	name := "小红"
	score := 92.5
	// %s 占位字符串，%v 占位任意值，\n 是换行
	fmt.Printf("姓名：%s，成绩：%v 分\n", name, score)
	// Sprintf 不直接打印，而是“生成”一个字符串并返回
	msg := fmt.Sprintf("恭喜 %s 考了 %.1f 分", name, score)
	fmt.Println(msg)
}

// 程序入口。
func main() {
	demoOS()
	demoStrconv()
	demoFmt()
}
