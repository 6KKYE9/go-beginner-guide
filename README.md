# Go 初学者学习项目（go-beginner-guide）

这是给编程新手准备的 Go 入门项目。每个示例都能直接跑、看到真实输出，代码里的注释就当讲解用，不用额外查资料。跟着跑完，你应该能写出可运行的 Go 小程序，并用 Git/GitHub 管理代码。

## 环境要求

| 工具 | 版本要求 | 作用 |
|------|----------|------|
| Go   | >= 1.21（本项目用 1.26 测试） | 编译运行 Go 代码 |
| Git  | 任意较新版本 | 把代码上传到 GitHub |

只用 Go 自带的"标准库"，不需要 `go get` 装任何东西。

检查 Go 是否装好：

```powershell
go version
```

预期输出类似：`go version go1.26.1 windows/amd64`（版本号可能不同，能显示即成功）。

## 安装与运行步骤

### 第 1 步：获取项目代码

如果从 GitHub 克隆（后面会讲怎么上传），用：

```powershell
git clone <你的仓库地址>
cd go-beginner-guide
```

本地已有本文件夹就直接 `cd` 进目录。

### 第 2 步：运行示例程序

**示例 1 —— 基础语法演示（变量/控制流/函数）：**

```powershell
go run ./basics
```

预期输出：

```
姓名： 小明
年龄： 18
身高： 1.75
圆周率： 3.14
成绩 75 分：及格啦！
这是第 1 次循环
这是第 2 次循环
这是第 3 次循环
这是第 4 次循环
这是第 5 次循环
周末：好好休息
3 + 5 = 8
```

**示例 2 —— 标准库演示（fmt / os / strconv）：**

```powershell
go run ./stdlib 小红 20
```

注意后面跟了两个参数：名字和年龄。

预期输出：

```
你好， 小红 ！你输入的年龄字符串是： 20
明年你就 19 岁了
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
  1. 小明        分数：88.5
  2. 小红        分数：95.0
  3. 小刚        分数：76.0
  4. 小美        分数：91.5
  全班平均分：87.75
  第一名：小红（95.0 分）

  1. 小红 95.0
  2. 小美 91.5
  3. 小明 88.5
  4. 小刚 76.0

   小明、小红、小刚、小美
```

**示例 5 —— 并发编程入门（goroutine + channel + WaitGroup）：**

```powershell
go run ./concurrency
```

预期输出（完成顺序可能不同，因为 3 个工人并发工作）：

```
工人 1 开始处理任务 1
工人 2 开始处理任务 2
工人 3 开始处理任务 3
工人 1 开始处理任务 4
工人 2 开始处理任务 5
全部结果（按完成顺序）: [2 4 6 8 10]
```

`go 函数()` 就能开一个并发任务；`chan` 是它们之间的传声筒；`sync.WaitGroup` 像点名册，主程序等所有工人干完再收工。注释里有逐行讲解。

**示例 6 —— 错误处理入门（error / errors.New / fmt.Errorf）：**

```powershell
go run ./errors
```

预期输出：

```
10 / 2 = 5.00
出错了: 除数不能为 0
出错了: 不能对负数开平方，收到的值: -4
sqrt(9) = 3.00
```

Go 没有 try/catch，而是把"可能出错"作为函数的一个普通返回值（类型 `error`）。约定写法：`if err != nil { 处理错误 } else { 用结果 }`。

**示例 7 —— 文本处理入门（strings / strconv）：**

```powershell
go run ./text
```

预期输出：

```text
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

八成时间都在跟字符串打交道。`strings.TrimSpace` 去空格、`Split/Join` 切了再拼、`ReplaceAll` 全量替换；`strconv.Atoi/Itoa` 做字符串和数字互转。注释里都标了坑。

**示例 8 —— 文件读写入门（os / bufio）：**

```powershell
go run ./files
```

预期输出：

```text
文件读写演示（数据存在 .notes.txt ）
读到了 3 行：
  1. 买牛奶
  2. 写代码
  3. 散步
含"代码"的有 1 行：[写代码]
```

`os.OpenFile` 用 `O_APPEND|O_CREATE` 就能追加写；`bufio.Scanner` 按行读最省事。文件不存在时 `os.Open` 会报 `os.ErrNotExist`，处理一下就当空文件。

**示例 9 —— 时间处理入门（time）：**

```powershell
go run ./timepkg
```

预期输出（时间部分会因运行时刻不同而变化）：

```text
现在： 2026-08-06 17:20:00
90 秒后： 17:21:30
差了： 90 秒
睡了 1 秒
解析出来的星期： Thursday
```

`time.Now()` 拿当前时间；格式化要用 Go 的"参考时间" `2006-01-02 15:04:05` 当模板，别用 `YYYY-MM-DD` 那套。`Add`/`Sub`/`Since` 算时间差，`Sleep` 真会停，`Parse` 把字符串按同样模板解析回来。

**示例 10 —— 迷你 Web 服务（net/http）：**

```powershell
go run ./web
```

启动后会一直监听 8080 端口（终端会打印访问地址），另开一个终端用 curl 访问：

```powershell
curl http://localhost:8080/
curl "http://localhost:8080/hello?name=小明"
curl http://localhost:8080/api/ping
```

各接口预期返回：

```text
# GET /
欢迎来到迷你 Web 服务！
你访问的是：/
这是第 1 次被访问。

# GET /hello?name=小明
你好，小明！这是 /hello 接口。

# GET /api/ping
{"status":"ok","message":"pong"}
```

用 `http.HandleFunc` 把"网址路径"绑到"处理函数"，处理函数里用 `http.ResponseWriter` 写回内容、用 `*http.Request` 读请求（比如 `r.URL.Query()` 取参数）。`http.ListenAndServe(":8080", mux)` 启动后阻塞监听，按 `Ctrl+C` 停止。想验证又不想开服务器？`web/server_test.go` 用 `httptest` 在内存里打这三个路由，直接 `go test ./web` 即可。

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

## 各示例程序说明

| 文件 | 学到的知识点 | 关键函数 |
|------|--------------|----------|
| `basics/basics.go` | 变量三种声明方式、const 常量、if-else、for 循环、switch、函数定义 | `add()` |
| `stdlib/stdlib.go` | 命令行参数读取、字符串转数字、格式化输出 | `demoOS()` `demoStrconv()` `demoFmt()` |
| `games/guess.go` | 随机数、缓冲读取输入、循环+条件判断综合应用 | `guessNumber()` |
| `structs/roster.go` | 结构体 struct、方法 method、切片排序、字符串拼接、格式化输出 | `Add()` `Top()` `Best()` `AverageScore()` |
| `concurrency/worker.go` | goroutine 并发、channel 通道通信、sync.WaitGroup 同步、关闭通道 | `RunConcurrency()` `worker()` |
| `errors/errors.go` | 错误作为返回值、`errors.New` 创建错误、`fmt.Errorf` 带上下文、调用方错误处理 | `divide()` `sqrt()` |
| `text/text.go` | `strings` 切分/拼接/替换/前后缀、`strconv` 字符串与数字互转 | `RunText()` |
| `files/files.go` | `os` 读写文件、`bufio` 按行读、字符串过滤、数据持久化 | `appendNote()` `readNotes()` |
| `timepkg/timepkg.go` | `time` 当前时间、格式化（参考时间 2006-01-02 15:04:05）、时间差、Sleep、Parse | `demoNow()` `demoSleep()` `demoParse()` |
| `web/server.go` | `net/http` 注册路由、处理函数读写请求响应、启动 HTTP 服务、URL 查询参数 | `homeHandler()` `helloHandler()` `apiHandler()` `setupRoutes()` |
| `todo/todo.go` | 结构体 struct、切片 slice、map 思路、JSON 文件持久化、flag 参数解析 | `add()` `list()` `done()` `rm()` `clear()` |

建议阅读顺序：先读 `basics` → 再读 `stdlib` → 然后挑战 `games` → 接着学 `structs` → 最后研究 `todo`。

## 待办事项工具（todo）详解

这是一个能真正日常使用的小工具，把你学的知识串起来：用 `struct` 表示一条待办，用 `slice` 存很多条，用 `json` 文件保存到磁盘（关掉程序数据也不丢），用 `flag` 解析命令。

运行示例：

```powershell
go run ./todo add "学习 Go 语言"     # 添加一条
go run ./todo add "去超市买菜"       # 再添加一条
go run ./todo list                   # 列出全部
go run ./todo done 1                 # 把第 1 条标记完成
go run ./todo rm 2                   # 删除第 2 条
go run ./todo clear                  # 清空所有
go run ./todo help                   # 查看帮助
```

预期输出（list 时）：

```
[x] 1. 学习 Go 语言
[ ] 2. 去超市买菜
```

（`[x]` 表示已完成，`[ ]` 表示未完成）

数据保存在项目目录下的 `.todo.json` 文件中（已被 .gitignore 忽略，不会上传）。

## 常见问题

**Q1：`go run` 报错 `go: command not found`？**
A：Go 没装好或没加入 PATH。重新安装 Go 并勾选"Add to PATH"。

**Q2：猜数字游戏运行时卡住不动？**
A：它在等你输入数字并回车。输入一个数字按回车即可，直到猜中。

**Q3：修改代码后运行没变化？**
A：`go run` 每次都会重新编译，通常不会有缓存问题；如果怀疑，先关掉再重跑。

**Q4：想加自己的练习？**
A：在对应目录下新建 `xxx.go`，包名写 `package main` 并写 `func main()`，就能独立运行。
