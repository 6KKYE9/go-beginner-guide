package main

import (
	"sort"
	"testing"
)

// TestRunConcurrency 验证：5 个任务（1..5）每个翻倍，结果应为 2,4,6,8,10。
// 由于并发完成顺序不确定，先排序再比较，只关心“最终集合”是否正确。
func TestRunConcurrency(t *testing.T) {
	got := RunConcurrency()
	if len(got) != 5 {
		t.Fatalf("期望 5 个结果，实际得到 %d 个: %v", len(got), got)
	}
	sort.Ints(got)
	want := []int{2, 4, 6, 8, 10}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("位置 %d: 期望 %d，实际 %d", i, want[i], got[i])
		}
	}
}
