// 文件读写入门：怎么把数据存到磁盘、再读回来。
// 用到 os（读写文件）、bufio（按行读）、配合前面的字符串处理。
// 跑一下：go run ./files
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const noteFile = ".notes.txt"

// 写一行到文件末尾，没有就新建
func appendNote(line string) error {
	f, err := os.OpenFile(noteFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(line + "\n")
	return err
}

// 把文件按行读回来，空文件返回空切片
func readNotes() ([]string, error) {
	f, err := os.Open(noteFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer f.Close()

	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines, sc.Err()
}

func RunFiles() {
	fmt.Println("文件读写演示（数据存在", noteFile, "）")

	_ = os.Remove(noteFile) // 先清掉，保证演示可重复
	if err := appendNote("买牛奶"); err != nil {
		fmt.Println("写失败:", err)
		return
	}
	_ = appendNote("写代码")
	_ = appendNote("散步")

	lines, err := readNotes()
	if err != nil {
		fmt.Println("读失败:", err)
		return
	}
	fmt.Printf("读到了 %d 行：\n", len(lines))
	for i, l := range lines {
		fmt.Printf("  %d. %s\n", i+1, l)
	}

	// 顺手演示：只读带"代码"的那几行
	var hit []string
	for _, l := range lines {
		if strings.Contains(l, "代码") {
			hit = append(hit, l)
		}
	}
	fmt.Printf("含\"代码\"的有 %d 行：%v\n", len(hit), hit)
}

func main() {
	RunFiles()
}
