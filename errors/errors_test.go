package main

import "testing"

// 测试 divide：正常除法应返回正确结果且无错误；除 0 应返回错误。
func TestDivide(t *testing.T) {
	if r, err := divide(10, 2); err != nil || r != 5 {
		t.Errorf("divide(10,2) = (%v, %v)，期望 (5, nil)", r, err)
	}
	if _, err := divide(10, 0); err == nil {
		t.Errorf("divide(10,0) 期望返回错误，但得到 nil")
	}
}

// 测试 sqrt：负数应返回错误，正数应返回近似正确结果。
func TestSqrt(t *testing.T) {
	if _, err := sqrt(-1); err == nil {
		t.Errorf("sqrt(-1) 期望返回错误，但得到 nil")
	}
	s, err := sqrt(9)
	if err != nil {
		t.Fatalf("sqrt(9) 意外返回错误: %v", err)
	}
	// 3 的平方是 9，结果应非常接近 3。
	if s < 2.99 || s > 3.01 {
		t.Errorf("sqrt(9) = %v，期望约等于 3", s)
	}
}
