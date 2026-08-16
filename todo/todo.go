// 待办事项命令行小工具：struct 打包一条待办，slice 存一摞，JSON 落盘不丢。
// 用法：
//
//	go run ./todo add "买牛奶"   # 加一条
//	go run ./todo list           # 看全部
//	go run ./todo done 1         # 标记第 1 条完成
//	go run ./todo rm 2           # 删第 2 条
//	go run ./todo clear          # 清空
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Todo 是一条待办，字段后的 json 标签决定存盘时用什么字段名。
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// 数据落在本目录的 .todo.json，每次运行读同一个文件。
// 用变量而非常量，便于测试时重定向到临时文件。
var dataFile = ".todo.json"

// 读盘；文件不存在（第一次跑）就当空列表返回。
func loadTodos() []Todo {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return []Todo{}
	}
	var todos []Todo
	// 解析失败要明确提示，而不是悄悄返回半截数据
	if err := json.Unmarshal(data, &todos); err != nil {
		fmt.Printf("警告：%s 内容损坏，已按空列表处理：%v\n", dataFile, err)
		return []Todo{}
	}
	return todos
}

// 覆盖写回磁盘。0644 是文件权限，不存在就新建。
// 返回错误，避免磁盘满/无权限时用户误以为保存成功。
func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func add(title string) {
	todos := loadTodos()
	// 新编号取当前最大 ID +1（而非最后一条），即使中间删过也绝不重复
	newID := 1
	for _, t := range todos {
		if t.ID >= newID {
			newID = t.ID + 1
		}
	}
	todos = append(todos, Todo{ID: newID, Title: title})
	if err := saveTodos(todos); err != nil {
		fmt.Printf("保存失败：%v\n", err)
		return
	}
	fmt.Printf("已添加：%s（编号 %d）\n", title, newID)
}

func list() {
	todos := loadTodos()
	if len(todos) == 0 {
		fmt.Println("暂无待办，用 add 添加一条吧～")
		return
	}
	for _, t := range todos {
		// [x] / [ ] 一眼看出完成没
		status := "[ ]"
		if t.Completed {
			status = "[x]"
		}
		fmt.Printf("%s %d. %s\n", status, t.ID, t.Title)
	}
}

func done(id int) {
	todos := loadTodos()
	found := false
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("找不到编号为 %d 的待办\n", id)
		return
	}
	if err := saveTodos(todos); err != nil {
		fmt.Printf("保存失败：%v\n", err)
		return
	}
	fmt.Printf("已标记 %d 为完成 ✅\n", id)
}

func rm(id int) {
	todos := loadTodos()
	newTodos := []Todo{}
	found := false
	for _, t := range todos {
		if t.ID == id {
			found = true // 跳过这条就等于删掉
			continue
		}
		newTodos = append(newTodos, t)
	}
	if !found {
		fmt.Printf("找不到编号为 %d 的待办\n", id)
		return
	}
	if err := saveTodos(newTodos); err != nil {
		fmt.Printf("保存失败：%v\n", err)
		return
	}
	fmt.Printf("已删除编号 %d\n", id)
}

func clear() {
	if err := saveTodos([]Todo{}); err != nil {
		fmt.Printf("清空失败：%v\n", err)
		return
	}
	fmt.Println("已清空所有待办")
}

func help() {
	fmt.Println(`待办事项小工具 - 用法：
  todo add "内容"   添加一条待办
  todo list         列出所有待办
  todo done <编号>  标记完成
  todo rm <编号>    删除一条
  todo clear        清空全部
  todo help         显示本帮助`)
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	cmd := os.Args[1]

	switch cmd {
	case "add":
		fs := flag.NewFlagSet("add", flag.ExitOnError)
		fs.Parse(os.Args[2:])
		args := fs.Args()
		if len(args) == 0 {
			fmt.Println("用法：todo add \"要记的事\"")
			return
		}
		// 多词拼成一条标题
		title := ""
		for i, a := range args {
			if i > 0 {
				title += " "
			}
			title += a
		}
		add(title)

	case "list":
		list()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("用法：todo done <编号>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("编号必须是数字")
			return
		}
		done(id)

	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("用法：todo rm <编号>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("编号必须是数字")
			return
		}
		rm(id)

	case "clear":
		clear()

	case "help":
		help()

	default:
		fmt.Printf("未知命令：%s\n", cmd)
		help()
	}
}
