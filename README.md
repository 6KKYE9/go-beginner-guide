# 🐹 Go 初学者学习项目（go-beginner-guide）

> 这是一个专门为**编程小白**设计的 Go 语言入门项目。每一行代码都配有详细中文注释，
> 帮助你从零开始理解变量、函数、控制流、标准库，以及如何用 Git 把代码传到 GitHub。

---

## 一、项目介绍与学习价值

很多人学编程卡在"看了教程却不知道怎么动手"。本项目反其道而行：

- **边读边跑**：每个示例都能直接运行，看到真实输出，比纯文字更直观。
- **注释即教程**：代码里的注释就是讲解，不需要额外查资料。
- **覆盖核心**：变量、数据类型、if/for/switch、函数、常用标准库，一次学全。
- **带测试**：学会写测试，养成"验证代码正确性"的好习惯（程序员的核心素养）。

学完本项目，你将能够：**独立写出可运行的 Go 小程序，并用 Git/GitHub 管理自己的代码**。

---

## 二、环境要求

| 工具 | 版本要求 | 作用 |
|------|----------|------|
| Go   | >= 1.21（本项目用 1.26 测试） | 编译运行 Go 代码 |
| Git  | 任意较新版本 | 把代码上传到 GitHub |

> 本项目**没有任何第三方依赖**，只用 Go 自带的"标准库"，所以你不需要 `go get` 安装任何东西。

**检查 Go 是否安装成功：**
```powershell
go version
```
预期输出类似：`go version go1.26.1 windows/amd64`（版本号可能不同，能显示即成功）。

---

## 三、安装与运行步骤

### 第 1 步：获取项目代码
如果你是从 GitHub 克隆（后面会讲怎么上传），用：
```powershell
git clone <你的仓库地址>
cd go-beginner-guide
```
如果是本地已有本文件夹，直接 `cd` 进目录即可。

### 第 2 步：运行三个示例程序

**示例 1 —— 基础语法演示（变量/控制流/函数）：**
```powershell
go run ./basics
```
预期输出：
```
===== 变量演示 =====
姓名： 小明
年龄： 18
身高： 1.75
圆周率： 3.14
===== if-else 演示 =====
成绩 75 分：及格啦！
===== for 循环演示 =====
这是第 1 次循环
这是第 2 次循环
这是第 3 次循环
这是第 4 次循环
这是第 5 次循环
===== switch 演示 =====
周末：好好休息
===== 函数演示 =====
3 + 5 = 8
```

**示例 2 —— 标准库演示（fmt / os / strconv）：**
```powershell
go run ./stdlib 小红 20
```
> 注意后面跟了两个参数：名字和年龄。
预期输出：
```
===== os 包：命令行参数演示 =====
你好， 小红 ！你输入的年龄字符串是： 20
===== strconv 包：字符串转数字演示 =====
明年你就 21 岁了
===== fmt 包：格式化输出演示 =====
姓名：小红，成绩：92.5 分
恭喜 小红 考了 92.5 分
```

**示例 3 —— 猜数字小游戏（综合练习）：**
```powershell
go run ./games
```
预期交互（你输入的数字可能不同）：
```
我已经想好了一个 1~100 之间的数字，来猜猜看吧！
请输入你的猜测：50
太大了！
请输入你的猜测：25
太小了！
请输入你的猜测：37
恭喜你，猜中了！答案就是 37
```

**示例 4 —— 结构体与方法（学生成绩册）：**
```powershell
go run ./structs
```
预期输出：
```
===== 学生成绩册 =====
  1. 小明        分数：88.5
  2. 小红        分数：95.0
  3. 小刚        分数：76.0
  4. 小美        分数：91.5
  全班平均分：87.75
  第一名：小红（95.0 分）

===== 按分数排序后的名单 =====
  1. 小红 95.0
  2. 小美 91.5
  3. 小明 88.5
  4. 小刚 76.0

===== 全部学生姓名 =====
   小明、小红、小刚、小美
```

**示例 5 —— 并发编程入门（goroutine + channel + WaitGroup）：**
```powershell
go run ./concurrency
```
预期输出（完成顺序可能不同，因为 3 个工人并发工作）：
```
===== 并发入门：3 个工人处理 5 个任务 =====
工人 1 开始处理任务 1
工人 2 开始处理任务 2
工人 3 开始处理任务 3
工人 1 开始处理任务 4
工人 2 开始处理任务 5
全部结果（按完成顺序）: [2 4 6 8 10]
```
> 这是 Go 的"杀手锏"。`go 函数()` 就能开一个并发任务；`chan` 是它们之间的传声筒；
> `sync.WaitGroup` 像点名册，主程序等所有工人干完再收工。注释里有逐行讲解。

**示例 6 —— 错误处理入门（error / errors.New / fmt.Errorf）：**
```powershell
go run ./errors
```
预期输出：
```
===== 错误处理演示 =====
10 / 2 = 5.00
出错了: 除数不能为 0
出错了: 不能对负数开平方，收到的值: -4
sqrt(9) = 3.00
```
> Go 没有 try/catch，而是把"可能出错"作为函数的一个普通返回值（类型 `error`）。
> 约定写法：`if err != nil { 处理错误 } else { 用结果 }`。学会它，你的程序才健壮。

**示例 7 —— 文本处理入门（strings / strconv）：**
```powershell
go run ./text
```
预期输出：
```text
===== 文本处理演示 =====
原文: "  Hello, Go World!  "
去首尾空格: "Hello, Go World!"
转小写: "  hello, go world!  "
是否以 Hello 开头: true
World 出现了吗: true
把 o 换成 0: "  Hell0, G0 W0rld!  "
切分后: [apple banana cherry]
用 | 拼回: apple | banana | cherry
Atoi 结果: 42，加 8 等于 50
Itoa(100) -> 100
FormatFloat: 3.14
```
> 八成时间都在跟字符串打交道。cuts：`strings.TrimSpace` 去空格、`Split/Join` 切了再拼、`ReplaceAll` 全量替换；`strconv.Atoi/Itoa` 做字符串和数字互转。注释里都标了坑。

### 第 3 步：运行测试
```powershell
go test ./...
```
预期输出：
```
ok  	go-beginner-guide/basics	0.019s
ok  	go-beginner-guide/structs	0.00Xs
?   	go-beginner-guide/games	[no test files]
?   	go-beginner-guide/stdlib	[no test files]
```
`ok` 表示测试通过；`?` 表示这个包还没有写测试文件（正常）。

---

## 四、项目文件结构

```
go-beginner-guide/
├── go.mod              # 模块定义文件（Go 项目的"身份证"）
├── .gitignore          # 告诉 Git 哪些文件不要上传（如编译产物）
├── README.md           # 你正在看的说明文档
├── basics/
│   ├── basics.go       # 示例1：变量、数据类型、控制流、函数
│   └── basics_test.go  # 示例1 的测试文件
├── stdlib/
│   └── stdlib.go       # 示例2：fmt / os / strconv 标准库
├── games/
│   └── guess.go        # 示例3：猜数字小游戏（综合练习）
├── structs/
│   ├── roster.go       # 示例4：结构体与方法——学生成绩册
│   └── roster_test.go  # 示例4 的测试文件
├── concurrency/
│   ├── worker.go       # 示例5：并发入门——goroutine / channel / WaitGroup
│   └── worker_test.go  # 示例5 的测试文件
├── errors/
│   ├── errors.go       # 示例6：错误处理入门——error / errors.New / fmt.Errorf
│   └── errors_test.go  # 示例6 的测试文件
├── text/
│   ├── text.go         # 示例7：文本处理入门——strings / strconv
│   └── text_test.go    # 示例7 的测试文件
└── todo/
    └── todo.go         # 示例5：命令行待办事项工具（struct/slice/JSON持久化）
```

---

## 五、各示例程序说明

| 文件 | 学到的知识点 | 关键函数 |
|------|--------------|----------|
| `basics/basics.go` | 变量三种声明方式、const 常量、if-else、for 循环、switch、函数定义 | `add()` |
| `stdlib/stdlib.go` | 命令行参数读取、字符串转数字、格式化输出 | `demoOS()` `demoStrconv()` `demoFmt()` |
| `games/guess.go` | 随机数、缓冲读取输入、循环+条件判断综合应用 | `guessNumber()` |
| `structs/roster.go` | 结构体 struct、方法 method、切片排序、字符串拼接、格式化输出 | `Add()` `Top()` `Best()` `AverageScore()` |
| `concurrency/worker.go` | goroutine 并发、channel 通道通信、sync.WaitGroup 同步、关闭通道 | `RunConcurrency()` `worker()` |
| `errors/errors.go` | 错误作为返回值、`errors.New` 创建错误、`fmt.Errorf` 带上下文、调用方错误处理 | `divide()` `sqrt()` |
| `text/text.go` | `strings` 切分/拼接/替换/前后缀、`strconv` 字符串与数字互转 | `RunText()` |
| `todo/todo.go` | 结构体 struct、切片 slice、map 思路、JSON 文件持久化、flag 参数解析 | `add()` `list()` `done()` `rm()` `clear()` |

建议阅读顺序：**先读 `basics` → 再读 `stdlib` → 然后挑战 `games` → 接着学 `structs` → 最后研究 `todo`**。

---

## 五之二、待办事项工具（todo）详解

这是一个**能真正日常使用**的小工具，把你学的知识串起来：用 `struct` 表示一条待办，用 `slice` 存很多条，用 `json` 文件保存到磁盘（关掉程序数据也不丢），用 `flag` 解析命令。

**运行示例：**
```powershell
go run ./todo add "学习 Go 语言"     # 添加一条
go run ./todo add "去超市买菜"       # 再添加一条
go run ./todo list                   # 列出全部
go run ./todo done 1                 # 把第 1 条标记完成
go run ./todo rm 2                   # 删除第 2 条
go run ./todo clear                  # 清空所有
go run ./todo help                   # 查看帮助
```

**预期输出（list 时）：**
```
===== 待办列表 =====
[x] 1. 学习 Go 语言
[ ] 2. 去超市买菜
```
（`[x]` 表示已完成，`[ ]` 表示未完成）

> 💡 数据保存在项目目录下的 `.todo.json` 文件中（已被 .gitignore 忽略，不会上传）。

---

## 六、常见问题（小白必看）

**Q1：`go run` 报错 `go: command not found`？**
A：Go 没装好或没加入环境变量 PATH。重新安装 Go 并勾选"Add to PATH"。

**Q2：猜数字游戏运行时卡住不动？**
A：它在等你输入数字并回车。输入一个数字按回车即可，直到猜中。

**Q3：修改代码后运行没变化？**
A：`go run` 每次都会重新编译，通常不会有缓存问题；如果怀疑，先关掉再重跑。

**Q4：想加自己的练习？**
A：在对应目录下新建 `xxx.go`，包名写 `package main` 并写 `func main()`，就能独立运行。

---

祝学习愉快！🎉 把本项目跑通并上传到 GitHub，你就完成了"从零到上线"的第一步。
