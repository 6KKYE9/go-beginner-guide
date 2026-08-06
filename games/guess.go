// 猜数字小游戏：综合练一下变量、循环、判断、随机数、读输入。
// 程序想一个 1~100 的数，你猜，它提示大/小，直到猜中。
// 跑一下：go run ./games
package main

import (
	"bufio"     // 带缓冲，方便逐行读键盘
	"fmt"
	"math/rand" // 生成随机数
	"os"
	"strconv"
	"strings"
	"time" // 给随机数"播种"，不然每次跑结果都一样
)

// 游戏核心逻辑单独拎出来，结构清楚也好测
func guessNumber() {
	// 用当前时间当种子，否则每次启动随机数序列一模一样
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1 // 1~100 的随机整数

	fmt.Println("我已经想好了一个 1~100 之间的数字，来猜猜看吧！")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入你的猜测：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		guess, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("请输入一个有效的数字哦～")
			continue
		}

		if guess > target {
			fmt.Println("太大了！")
		} else if guess < target {
			fmt.Println("太小了！")
		} else {
			fmt.Println("恭喜你，猜中了！答案就是", target)
			break
		}
	}
}

func main() {
	guessNumber()
}
