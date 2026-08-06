package main

import (
	"os"
	"path/filepath"
	"testing"
)

// withTempFile 把 dataFile 重定向到临时文件，返回清理函数。
func withTempFile(t *testing.T) func() {
	t.Helper()
	dir := t.TempDir()
	dataFile = filepath.Join(dir, ".todo.json")
	return func() {
		dataFile = ".todo.json"
	}
}

func TestAddAndList(t *testing.T) {
	cleanup := withTempFile(t)
	defer cleanup()

	add("买牛奶")
	add("写代码")
	todos := loadTodos()
	if len(todos) != 2 {
		t.Fatalf("应有 2 条，实际 %d", len(todos))
	}
	if todos[0].Title != "买牛奶" || todos[1].Title != "写代码" {
		t.Fatalf("标题顺序错误: %v", todos)
	}
	// 编号应从 1 递增
	if todos[0].ID != 1 || todos[1].ID != 2 {
		t.Fatalf("编号错误: %v", todos)
	}
}

func TestAddIDNoReuse(t *testing.T) {
	cleanup := withTempFile(t)
	defer cleanup()

	add("a")
	add("b")
	add("c")
	// 删掉中间一条（id=2），再添加应取 max+1 = 4，而不是复用 2
	rm(2)
	add("d")
	todos := loadTodos()
	ids := map[int]bool{}
	for _, x := range todos {
		if ids[x.ID] {
			t.Fatalf("出现重复编号: %d", x.ID)
		}
		ids[x.ID] = true
	}
	if _, ok := ids[4]; !ok {
		t.Fatalf("新编号应为 4，实际: %v", ids)
	}
}

func TestDone(t *testing.T) {
	cleanup := withTempFile(t)
	defer cleanup()

	add("task")
	done(1)
	todos := loadTodos()
	if !todos[0].Completed {
		t.Fatal("id=1 应被标记完成")
	}
	done(99) // 不存在的编号不应 panic
	if len(loadTodos()) != 1 {
		t.Fatal("标记不存在的编号不应改变数据")
	}
}

func TestRm(t *testing.T) {
	cleanup := withTempFile(t)
	defer cleanup()

	add("a")
	add("b")
	rm(1)
	todos := loadTodos()
	if len(todos) != 1 || todos[0].Title != "b" {
		t.Fatalf("删除后剩余异常: %v", todos)
	}
	rm(99) // 不存在不应 panic
}

func TestClear(t *testing.T) {
	cleanup := withTempFile(t)
	defer cleanup()

	add("a")
	add("b")
	clear()
	if len(loadTodos()) != 0 {
		t.Fatal("clear 后应为空")
	}
}

func TestLoadTodosMissingFile(t *testing.T) {
	cleanup := withTempFile(t)
	defer cleanup()
	// 临时目录里没有数据文件，应返回空列表且不报错
	if got := loadTodos(); len(got) != 0 {
		t.Fatalf("缺失文件应返回空列表, got %v", got)
	}
}

func TestLoadTodosCorrupt(t *testing.T) {
	cleanup := withTempFile(t)
	defer cleanup()
	// 写入非法 JSON，loadTodos 应回退为空列表（内部打印警告但不 panic）
	if err := os.WriteFile(dataFile, []byte("not valid json{{{"), 0644); err != nil {
		t.Fatal(err)
	}
	if got := loadTodos(); len(got) != 0 {
		t.Fatalf("损坏文件应回退空列表, got %v", got)
	}
}
