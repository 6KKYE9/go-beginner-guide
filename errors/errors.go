// 错误处理这块是 Go 里挺容易劝退新人的地方：它没有 try/catch，
// 而是把"可能出错"当作函数的一个普通返回值，交给调用方自己处理。
package main

import (
	"errors"
	"fmt"
)

// divide 除一下，除数为 0 就返回错误。
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

// sqrt 自己写个开平方（牛顿迭代随便迭代几次够用就行），负数没法开就报错。
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("不能对负数开平方，收到的值: %v", x)
	}
	guess := x / 2
	for i := 0; i < 10; i++ {
		guess = (guess + x/guess) / 2
	}
	return guess, nil
}

func demoHandle() {
	fmt.Println("===== 错误处理演示 =====")

	// 正常情况：err 是 nil 就直接用结果
	r, err := divide(10, 2)
	if err != nil {
		fmt.Println("出错了:", err)
		return
	}
	fmt.Printf("10 / 2 = %.2f\n", r)

	// 除 0 会进错误分支
	r, err = divide(10, 0)
	if err != nil {
		fmt.Println("出错了:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", r)
	}

	// fmt.Errorf 能把具体数值拼进错误信息里，方便排查
	s, err := sqrt(-4)
	if err != nil {
		fmt.Println("出错了:", err)
	} else {
		fmt.Printf("sqrt(-4) = %.2f\n", s)
	}

	s, err = sqrt(9)
	if err == nil {
		fmt.Printf("sqrt(9) = %.2f\n", s)
	}
}

func main() {
	demoHandle()
}
