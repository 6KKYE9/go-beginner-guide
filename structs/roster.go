// 本文件演示 Go 语言中非常重要的概念：结构体（struct）与方法（method）。
// 我们用一个“学生成绩册”作为例子，把前面学的变量、切片、循环串起来。
//
// 你将学到：
//   - 用 struct 把多个相关的数据打包成一个“对象”（如：姓名 + 分数）
//   - 给 struct 定义“方法”（函数绑定到类型上，像 o.Average() 这样调用）
//   - 用 sort.Slice 对切片按规则排序
//   - 用 strings.Join 把多个字符串拼成一个
//   - 用 fmt.Printf 做更精致的格式化输出
//
// 运行：go run ./structs
package main

import (
	"fmt"
	"sort"
	"strings"
)

// Student 表示一个学生的信息。
// struct 就像一个“自定义模板”，把姓名和分数放在一起。
type Student struct {
	Name  string  // 姓名
	Score float64 // 分数（允许小数，如 92.5）
}

// Average 是 Student 的“方法”：计算该学生的平均得分。
// 这里只有一个学生，平均值就是他自己的分数；写成方法是为了演示“方法”的写法。
// 语法：(s Student) 表示这个方法绑定在 Student 类型上，调用时像 s.Average()。
func (s Student) Average() float64 {
	return s.Score
}

// Roster 是“成绩册”，本质是一个 Student 切片（存放很多学生）。
type Roster struct {
	Students []Student
}

// Add 往成绩册里添加一个学生。
func (r *Roster) Add(name string, score float64) {
	// 用 append 往切片末尾追加元素。注意接收者用 *Roster（指针），
	// 这样修改才能影响原始的成绩册，而不是它的副本。
	r.Students = append(r.Students, Student{Name: name, Score: score})
}

// Top 返回按分数从高到低排序后的学生列表（不修改原数据）。
func (r Roster) Top() []Student {
	// 复制一份，避免改动原始顺序
	cp := make([]Student, len(r.Students))
	copy(cp, r.Students)
	// sort.Slice 允许你自定义“谁排在前面”的规则。
	// 这里用匿名函数告诉它：分数大的排在前面（a.Score > b.Score）。
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Score > cp[j].Score
	})
	return cp
}

// Best 返回分数最高的学生；如果成绩册为空，返回 false。
func (r Roster) Best() (Student, bool) {
	if len(r.Students) == 0 {
		return Student{}, false
	}
	top := r.Top()
	return top[0], true
}

// AverageScore 计算全班平均分（用到了 for 循环累加）。
func (r Roster) AverageScore() float64 {
	if len(r.Students) == 0 {
		return 0
	}
	var sum float64
	for _, s := range r.Students {
		sum += s.Score
	}
	return sum / float64(len(r.Students))
}

// Names 把所有学生姓名拼成一个用顿号分隔的字符串。
// 演示 strings.Join 的用法：把 []string 拼成一个字符串。
func (r Roster) Names() string {
	names := make([]string, len(r.Students))
	for i, s := range r.Students {
		names[i] = s.Name
	}
	return strings.Join(names, "、")
}

// Print 把成绩册漂亮地打印出来。
func (r Roster) Print() {
	fmt.Println("===== 学生成绩册 =====")
	if len(r.Students) == 0 {
		fmt.Println("（还没有学生，先用 Add 添加吧）")
		return
	}
	// 用 fmt.Printf 的 %-10s（左对齐、占 10 字符宽度）让表格整齐。
	for i, s := range r.Students {
		fmt.Printf("  %d. %-10s 分数：%.1f\n", i+1, s.Name, s.Score)
	}
	fmt.Printf("  全班平均分：%.2f\n", r.AverageScore())
	if best, ok := r.Best(); ok {
		fmt.Printf("  第一名：%s（%.1f 分）\n", best.Name, best.Score)
	}
}

// 程序入口：构造一个成绩册并演示所有功能。
func main() {
	roster := Roster{}
	roster.Add("小明", 88.5)
	roster.Add("小红", 95.0)
	roster.Add("小刚", 76.0)
	roster.Add("小美", 91.5)

	roster.Print()

	fmt.Println("\n===== 按分数排序后的名单 =====")
	for i, s := range roster.Top() {
		fmt.Printf("  %d. %s %.1f\n", i+1, s.Name, s.Score)
	}

	fmt.Println("\n===== 全部学生姓名 =====")
	fmt.Println("  ", roster.Names())
}
