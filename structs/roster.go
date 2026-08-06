// struct 和方法（method）的演示，用"学生成绩册"当例子，
// 顺带把前面学的切片、循环、排序都串起来。
// 跑一下：go run ./structs
package main

import (
	"fmt"
	"sort"
	"strings"
)

// Student 一个学生的信息，struct 就是把姓名和分数打包在一起
type Student struct {
	Name  string  // 姓名
	Score float64 // 分数，允许小数
}

// Average 演示"方法"的写法：(s Student) 把函数绑到类型上，调用像 s.Average()
// 这里只有一个学生，平均值就是他自己分数
func (s Student) Average() float64 {
	return s.Score
}

// Roster 成绩册，本质是个 Student 切片
type Roster struct {
	Students []Student
}

// Add 往里塞一个学生。注意用 *Roster 指针接收者，改动才影响原件
func (r *Roster) Add(name string, score float64) {
	r.Students = append(r.Students, Student{Name: name, Score: score})
}

// Top 按分数从高到低排，不动原数据
func (r Roster) Top() []Student {
	cp := make([]Student, len(r.Students))
	copy(cp, r.Students)
	// 分数大的排前面
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Score > cp[j].Score
	})
	return cp
}

// Best 返回最高分学生，空册子返回 false
func (r Roster) Best() (Student, bool) {
	if len(r.Students) == 0 {
		return Student{}, false
	}
	top := r.Top()
	return top[0], true
}

// AverageScore 全班平均分
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

// Names 把所有姓名用顿号拼起来，顺手演示 strings.Join
func (r Roster) Names() string {
	names := make([]string, len(r.Students))
	for i, s := range r.Students {
		names[i] = s.Name
	}
	return strings.Join(names, "、")
}

// Print 把成绩册打出来，用 %-10s 左对齐让表格齐整
func (r Roster) Print() {
	fmt.Println("学生成绩册：")
	if len(r.Students) == 0 {
		fmt.Println("（还没有学生，先用 Add 添加吧）")
		return
	}
	for i, s := range r.Students {
		fmt.Printf("  %d. %-10s 分数：%.1f\n", i+1, s.Name, s.Score)
	}
	fmt.Printf("  全班平均分：%.2f\n", r.AverageScore())
	if best, ok := r.Best(); ok {
		fmt.Printf("  第一名：%s（%.1f 分）\n", best.Name, best.Score)
	}
}

func main() {
	roster := Roster{}
	roster.Add("小明", 88.5)
	roster.Add("小红", 95.0)
	roster.Add("小刚", 76.0)
	roster.Add("小美", 91.5)

	roster.Print()

	fmt.Println("\n按分数排序后的名单：")
	for i, s := range roster.Top() {
		fmt.Printf("  %d. %s %.1f\n", i+1, s.Name, s.Score)
	}

	fmt.Println("\n全部学生姓名：")
	fmt.Println("  ", roster.Names())
}
