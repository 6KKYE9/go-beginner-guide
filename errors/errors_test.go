package main

import "testing"

func TestDivide(t *testing.T) {
	r, err := divide(10, 2)
	if err != nil || r != 5 {
		t.Errorf("divide(10,2) = (%v, %v)，期望 (5, nil)", r, err)
	}
	// 除 0 必须报错
	if _, err := divide(10, 0); err == nil {
		t.Errorf("divide(10,0) 期望返回错误，但得到 nil")
	}
}

func TestSqrt(t *testing.T) {
	if _, err := sqrt(-1); err == nil {
		t.Errorf("sqrt(-1) 期望返回错误，但得到 nil")
	}
	s, err := sqrt(9)
	if err != nil {
		t.Fatalf("sqrt(9) 意外返回错误: %v", err)
	}
	// 近似值得在 3 附近
	if s < 2.99 || s > 3.01 {
		t.Errorf("sqrt(9) = %v，期望约等于 3", s)
	}
}
