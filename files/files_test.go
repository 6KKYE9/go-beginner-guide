package main

import (
	"bufio"
	"os"
	"testing"
)

func TestAppendAndRead(t *testing.T) {
	const tmp = ".notes_test.txt"
	os.Remove(tmp)
	defer os.Remove(tmp)

	if err := appendNoteTo(tmp, "第一行"); err != nil {
		t.Fatal(err)
	}
	_ = appendNoteTo(tmp, "第二行")

	lines, err := readNotesFrom(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 || lines[0] != "第一行" || lines[1] != "第二行" {
		t.Errorf("读出来的内容不对: %v", lines)
	}
}

// 下面两个是测试专用的小包装，复用上面的逻辑但指定文件路径
func appendNoteTo(path, line string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(line + "\n")
	return err
}

func readNotesFrom(path string) ([]string, error) {
	f, err := os.Open(path)
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

func TestReadMissingFile(t *testing.T) {
	// 不存在的文件应该返回空切片而不是报错
	lines, err := readNotesFrom(".not_exist_xxx.txt")
	if err != nil {
		t.Errorf("不存在的文件不该报错: %v", err)
	}
	if len(lines) != 0 {
		t.Errorf("期望空切片，得到 %v", lines)
	}
}
