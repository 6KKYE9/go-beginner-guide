package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestSplitJoin(t *testing.T) {
	parts := strings.Split("a,b,c", ",")
	if strings.Join(parts, "-") != "a-b-c" {
		t.Error("Split/Join 结果不对")
	}
}

func TestAtoiRoundtrip(t *testing.T) {
	n, err := strconv.Atoi("42")
	if err != nil || n != 42 {
		t.Error("Atoi 期望得到 42")
	}
	if strconv.Itoa(n+8) != "50" {
		t.Error("Itoa 期望得到 50")
	}
}
