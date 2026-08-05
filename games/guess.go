// 第三个示例：猜数字小游戏。
// 这是一个“综合练习”，用到了前面学到的 变量、循环、条件判断、随机数、输入读取。
// 程序会随机想一个 1~100 的数字，你来猜，它会提示“大了/小了”，直到猜中。
package main

import (
	"bufio"  // 带缓冲的读取，方便逐行读取键盘输入
	"fmt"
	"math/rand" // 生成随机数的标准库
	"os"
	"strconv"
	"strings" // 字符串处理，如去掉首尾空格
	"time"    // 用来给随机数“播种”，保证每次运行结果不同
)

// guessNumber 是游戏的核心逻辑，被 main 调用。
// 把逻辑写成单独的函数，结构更清晰，也方便写测试。
func guessNumber() {
	// 用当前时间作为随机种子，否则每次程序启动生成的“随机数”会一模一样
	rand.Seed(time.Now().UnixNano())
	// 生成 1~100 之间的随机整数
	target := rand.Intn(100) + 1

	fmt.Println("我已经想好了一个 1~100 之间的数字，来猜猜看吧！")

	// 创建一个读取器，从标准输入（键盘）读取内容
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入你的猜测：")
		// 读取一整行输入，直到按下回车
		text, _ := reader.ReadString('\n')
		// 去掉行尾的换行符和前后空格
		text = strings.TrimSpace(text)

		// 把输入的文字转换成整数
		guess, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("请输入一个有效的数字哦～")
			continue // 跳过本次循环，重新让用户输入
		}

		// 用 if 判断大小并给出提示
		if guess > target {
			fmt.Println("太大了！")
		} else if guess < target {
			fmt.Println("太小了！")
		} else {
			fmt.Println("恭喜你，猜中了！答案就是", target)
			break // 猜中后退出循环
		}
	}
}

func main() {
	guessNumber()
}
