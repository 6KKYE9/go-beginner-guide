// 本文件演示 Go 语言中最重要、也最容易让初学者困惑的概念：错误处理（error）。
// Go 不用 try/catch，而是把“可能出错”作为函数的一个普通返回值，由调用者决定怎么处理。
package main

import (
	"errors"
	"fmt"
)

// ============ 1. 返回错误 ============

// divide 演示“可能失败”的函数：除数为 0 时返回错误。
// 约定：Go 函数遇到错误时，错误值放在最后一个返回值，类型通常是 error。
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// errors.New 创建一个最简单的错误，里面只有一句说明文字。
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil // nil 表示“没有错误”
}

// ============ 2. 用错误包装带上下文的信息 ============

// sqrt 演示用 fmt.Errorf 生成带变量的错误（%v 把数值拼进提示里）。
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("不能对负数开平方，收到的值: %v", x)
	}
	// 这里用近似法（牛顿迭代）求平方根，只是为了演示，不用纠结数学细节。
	guess := x / 2
	for i := 0; i < 10; i++ {
		guess = (guess + x/guess) / 2
	}
	return guess, nil
}

// ============ 3. 调用者如何处理错误 ============

func demoHandle() {
	fmt.Println("===== 错误处理演示 =====")

	// 情况一：正常情况，错误是 nil，直接用结果。
	r, err := divide(10, 2)
	if err != nil {
		// 约定俗成的写法：先判断 err 是不是 nil，不是就说明出错了。
		fmt.Println("出错了:", err)
		return
	}
	fmt.Printf("10 / 2 = %.2f\n", r)

	// 情况二：出错情况，err 不是 nil，打印错误信息，不继续使用结果。
	r, err = divide(10, 0)
	if err != nil {
		fmt.Println("出错了:", err) // 会打印：出错了: 除数不能为 0
	} else {
		fmt.Printf("10 / 0 = %.2f\n", r)
	}

	// 情况三：带上下文的错误。
	s, err := sqrt(-4)
	if err != nil {
		fmt.Println("出错了:", err) // 会打印：不能对负数开平方，收到的值: -4
	} else {
		fmt.Printf("sqrt(-4) = %.2f\n", s)
	}

	s, err = sqrt(9)
	if err == nil {
		fmt.Printf("sqrt(9) = %.2f\n", s)
	}
}

// ============ 程序入口 ============

func main() {
	demoHandle()
}
