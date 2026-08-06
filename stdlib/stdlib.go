// 标准库就是 Go 官方自带、不用额外装的东西。
// 这里挑了三个平时最常用的：fmt 打印、os 读命令行参数、strconv 做字符串和数字互转。
package main

import (
	"fmt"
	"os"
	"strconv"
)

// 读命令行参数：go run stdlib.go 名字 年龄
func demoOS() {
	// os.Args[0] 是程序自己的路径，真正传进来的参数从下标 1 开始
	args := os.Args
	if len(args) < 3 {
		fmt.Println("提示：请这样运行 -> go run stdlib.go 你的名字 你的年龄")
		return
	}
	name := args[1]   // 第一个参数：名字
	ageStr := args[2] // 第二个参数：年龄，这时候还是字符串
	fmt.Println("你好，", name, "！你输入的年龄字符串是：", ageStr)
}

// 字符串转数字，strconv 干这个最方便
func demoStrconv() {
	ageStr := "18"
	// Atoi 返回两个数：结果和错误。出错时 err 不是 nil
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("转换失败：", err)
		return
	}
	fmt.Println("明年你就", age+1, "岁了")
}

// fmt 不止有 Println，还有按格式拼接的 Printf / Sprintf
func demoFmt() {
	name := "小红"
	score := 92.5
	// %s 塞字符串，%v 塞任意值，%.1f 保留一位小数
	fmt.Printf("姓名：%s，成绩：%v 分\n", name, score)
	// Sprintf 不打印，而是先攒成一个字符串再返回，方便后面接着用
	msg := fmt.Sprintf("恭喜 %s 考了 %.1f 分", name, score)
	fmt.Println(msg)
}

func main() {
	demoOS()
	demoStrconv()
	demoFmt()
}
