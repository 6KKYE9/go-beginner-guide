// 本文件是 roster.go 的测试。运行：go test ./structs
// 测试能帮你确认“代码按预期工作”，是程序员的好习惯。
package main

import "testing"

// 测试 Add 与 AverageScore 是否正确累加求平均。
func TestRosterAverage(t *testing.T) {
	r := Roster{}
	r.Add("A", 80)
	r.Add("B", 90)
	got := r.AverageScore()
	want := 85.0
	if got != want {
		t.Errorf("平均分错误：得到 %.2f，期望 %.2f", got, want)
	}
}

// 测试 Top 排序是否“分数高的在前”。
func TestRosterTop(t *testing.T) {
	r := Roster{}
	r.Add("低分", 50)
	r.Add("高分", 99)
	r.Add("中分", 75)
	top := r.Top()
	if top[0].Score != 99 {
		t.Errorf("第一名应该是 99 分，实际是 %.1f", top[0].Score)
	}
}

// 测试空成绩册时 Best 返回 false（避免越界）。
func TestRosterBestEmpty(t *testing.T) {
	r := Roster{}
	_, ok := r.Best()
	if ok {
		t.Error("空成绩册不应返回有效学生")
	}
}
