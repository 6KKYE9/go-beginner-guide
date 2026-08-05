// 待办事项（Todo List）命令行小工具
//
// 这是一个"稍微进阶"的练习，相比前面的示例，你会学到：
//   - 结构体 struct：把"一件事"的相关数据打包在一起
//   - 切片 slice（动态数组）：保存很多待办事项
//   - map：把命令行参数名映射到对应的处理函数
//   - 文件读写：把待办存到磁盘，关掉程序也不会丢
//   - flag 包：解析命令行参数（Go 官方推荐的参数解析标准库）
//
// 用法示例（在项目根目录执行）：
//   go run ./todo add "买牛奶"        # 添加一条待办
//   go run ./todo list               # 列出所有待办
//   go run ./todo done 1             # 把第 1 条标记为已完成
//   go run ./todo rm 2               # 删除第 2 条
//   go run ./todo clear             # 清空所有待办
package main

import (
	"encoding/json" // 把结构体转换成可保存的文本格式（JSON）
	"flag"           // 命令行参数解析
	"fmt"
	"os"  // 文件操作、退出程序
	"strconv"
)

// Todo 表示"一条待办事项"。
// 结构体字段后面的 `json:"..."` 是"标签"，告诉 json 包保存时用什么字段名。
type Todo struct {
	ID        int    `json:"id"`        // 编号，方便定位
	Title     string `json:"title"`     // 待办内容
	Completed bool   `json:"completed"` // 是否已完成
}

// 数据文件的保存路径（放在用户目录下，叫 todo.json）
// 这样每次运行都读同一个文件，数据就能持久保存。
const dataFile = ".todo.json"

// loadTodos 从文件读取所有待办事项。
// 如果文件不存在，就返回空的列表（不报错）。
func loadTodos() []Todo {
	// 尝试读取文件
	data, err := os.ReadFile(dataFile)
	if err != nil {
		// 文件不存在很正常（第一次运行），直接返回空列表
		return []Todo{}
	}
	var todos []Todo
	// 把 JSON 文本解析回结构体切片
	json.Unmarshal(data, &todos)
	return todos
}

// saveTodos 把待办列表写回文件（覆盖保存）。
func saveTodos(todos []Todo) {
	// 把结构体切片转换成 JSON 文本
	data, _ := json.MarshalIndent(todos, "", "  ")
	// 写入文件（如果不存在会创建，0644 是文件权限）
	os.WriteFile(dataFile, data, 0644)
}

// add 添加一条待办。
func add(title string) {
	todos := loadTodos()
	// 新编号 = 当前最大编号 + 1（空列表时从 1 开始）
	newID := 1
	if len(todos) > 0 {
		newID = todos[len(todos)-1].ID + 1
	}
	todos = append(todos, Todo{ID: newID, Title: title, Completed: false})
	saveTodos(todos)
	fmt.Printf("已添加：%s（编号 %d）\n", title, newID)
}

// list 列出所有待办。
func list() {
	todos := loadTodos()
	if len(todos) == 0 {
		fmt.Println("暂无待办，用 add 添加一条吧～")
		return
	}
	fmt.Println("===== 待办列表 =====")
	for _, t := range todos {
		// 用 [x] / [ ] 直观表示完成状态
		status := "[ ]"
		if t.Completed {
			status = "[x]"
		}
		fmt.Printf("%s %d. %s\n", status, t.ID, t.Title)
	}
}

// done 把指定编号的待办标记为已完成。
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
	saveTodos(todos)
	fmt.Printf("已标记 %d 为完成 ✅\n", id)
}

// rm 删除指定编号的待办。
func rm(id int) {
	todos := loadTodos()
	newTodos := []Todo{}
	found := false
	for _, t := range todos {
		if t.ID == id {
			found = true // 跳过这一条，等于删除
			continue
		}
		newTodos = append(newTodos, t)
	}
	if !found {
		fmt.Printf("找不到编号为 %d 的待办\n", id)
		return
	}
	saveTodos(newTodos)
	fmt.Printf("已删除编号 %d\n", id)
}

// clear 清空所有待办。
func clear() {
	saveTodos([]Todo{})
	fmt.Println("已清空所有待办")
}

// help 打印使用说明。
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
	// flag 包：第一个参数是"子命令"，后面跟着它自己的参数
	if len(os.Args) < 2 {
		help()
		return
	}

	// 子命令（add / list / done ...）
	cmd := os.Args[1]

	switch cmd {
	case "add":
		// 用 flag 解析 add 后面的文字参数
		fs := flag.NewFlagSet("add", flag.ExitOnError)
		fs.Parse(os.Args[2:])
		args := fs.Args()
		if len(args) == 0 {
			fmt.Println("用法：todo add \"要记的事\"")
			return
		}
		// 把多个词拼成一个标题
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
