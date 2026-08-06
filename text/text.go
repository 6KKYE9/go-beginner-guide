// 文本处理入门：strings 和 strconv 两个常用标准库
//
// 平时写程序，八成时间都在跟"字符串"打交道——切分、拼接、替换、转数字。
// Go 标准库里 strings 管文本操作，strconv 管字符串和数字的互转。
// 跑一下看效果：go run ./text
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 这是本示例的主入口，main 里直接调它
func RunText() {
	fmt.Println("===== 文本处理演示 =====")

	// 一、strings：大小写、前后缀、包含、替换
	s := "  Hello, Go World!  "
	fmt.Printf("原文: %q\n", s)
	fmt.Printf("去首尾空格: %q\n", strings.TrimSpace(s))
	fmt.Printf("转小写: %s\n", strings.ToLower(s))
	fmt.Printf("是否以 Hello 开头: %v\n", strings.HasPrefix(strings.TrimSpace(s), "Hello"))
	fmt.Printf("World 出现了吗: %v\n", strings.Contains(s, "World"))
	// Replace 默认只换第一个，要全换得给最后一个参数 -1
	fmt.Printf("把 o 换成 0: %s\n", strings.ReplaceAll(s, "o", "0"))

	// 二、Split / Join：字符串拆成切片，再拼回去
	csv := "apple,banana,cherry"
	fruits := strings.Split(csv, ",")
	fmt.Printf("切分后: %v\n", fruits)
	fmt.Printf("用 | 拼回: %s\n", strings.Join(fruits, " | "))

	// 三、strconv：字符串 <-> 数字
	numStr := "42"
	n, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("转数字失败:", err)
	} else {
		fmt.Printf("Atoi 结果: %d，加 8 等于 %d\n", n, n+8)
	}
	// 数字转字符串用 Itoa（其实就是 FormatInt 的简写）
	fmt.Printf("Itoa(100) -> %s\n", strconv.Itoa(100))
	// 浮点转字符串记得给位数，不然 Go 会按最短表示来
	f := 3.14159
	fmt.Printf("FormatFloat: %s\n", strconv.FormatFloat(f, 'f', 2, 64))
}

// main 让本示例能直接用 `go run ./text` 运行。
func main() {
	RunText()
}
