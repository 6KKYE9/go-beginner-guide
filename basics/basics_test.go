// 这是“测试文件”，用来验证我们的函数是否正常工作。
// Go 的测试约定：
//   - 文件名以 _test.go 结尾
//   - 函数名以 Test 开头，参数固定为 (t *testing.T)
// 运行测试的命令：go test ./...  （在项目根目录执行）
package main

import "testing" // testing 是 Go 自带的测试框架

// TestAdd 验证 add 函数：3+5 应该等于 8。
func TestAdd(t *testing.T) {
	got := add(3, 5) // 实际得到的值
	want := 8        // 我们期望的值
	if got != want {
		// t.Errorf 会在测试失败时记录错误，但不会立刻停止
		t.Errorf("add(3,5) = %d; 期望 %d", got, want)
	}
}

// TestAddNegative 再测一个负数的情况，确保函数更可靠。
func TestAddNegative(t *testing.T) {
	got := add(-2, 2)
	want := 0
	if got != want {
		t.Errorf("add(-2,2) = %d; 期望 %d", got, want)
	}
}
