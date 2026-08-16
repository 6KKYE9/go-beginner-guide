# Go 初学者学习项目

当年我学 Go 的时候，最大的痛苦不是语法，是"每个教程都讲得对，但我合上网页还是写不出一个能跑的东西"。

所以这个仓库不搞花活：每个示例都能直接 `go run`、能看到真实输出，代码里的注释就当讲解，不用再开十几个标签页查资料。你跟着跑完一轮，应该就能自己写出小程序，并且用 Git 把它管起来。

## 环境要求

| 工具 | 版本要求 | 作用 |
|------|----------|------|
| Go   | >= 1.21（本项目用 1.26 测过） | 编译运行 Go 代码 |
| Git  | 任意较新版本 | 把代码传上 GitHub |

只用 Go 自带的"标准库"，不需要 `go get` 装任何东西。

先确认 Go 装好了：

```powershell
go version
```

能显示类似 `go version go1.26.1 windows/amd64` 就成（版本号不同没事，能出来就行）。

## 跑起来（三步）

### 第 1 步：拿到代码

从 GitHub 克隆（后面会讲怎么传）：

```powershell
git clone <你的仓库地址>
cd go-beginner-guide
```

本地已经有这文件夹就直接 `cd` 进来。

### 第 2 步：跑示例

**示例 1 —— 基础语法（变量 / 控制流 / 函数）：**

```powershell
go run ./basics
```

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

**示例 2 —— 标准库（fmt / os / strconv）：**

```powershell
go run ./stdlib 小红 20
```

后面跟了两个参数：名字和年龄。

```
你好， 小红 ！你输入的年龄字符串是： 20
明年你就 19 岁了
姓名：小红，成绩：92.5 分
恭喜 小红 考了 92.5 分
```

**示例 3 —— 猜数字小游戏：**

```powershell
go run ./games
```

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

**示例 5 —— 并发入门（goroutine + channel + WaitGroup）：**

```powershell
go run ./concurrency
```

```
工人 1 开始处理任务 1
工人 2 开始处理任务 2
工人 3 开始处理任务 3
工人 1 开始处理任务 4
工人 2 开始处理任务 5
全部结果（按完成顺序）: [2 4 6 8 10]
```

`go 函数()` 就能开一个并发任务；`chan` 是它们之间的传声筒；`sync.WaitGroup` 像点名册，主程序等所有工人干完再收工。注释里逐行讲了。

**示例 6 —— 错误处理入门：**

```powershell
go run ./errors
```

```
10 / 2 = 5.00
出错了: 除数不能为 0
出错了: 不能对负数开平方，收到的值: -4
sqrt(9) = 3.00
```

Go 没有 try/catch，而是把"可能出错"当成一个普通返回值（类型 `error`）。约定写法：`if err != nil { 处理 } else { 用结果 }`。一开始别扭，习惯了比异常清爽。

**示例 7 —— 文本处理（strings / strconv）：**

```powershell
go run ./text
```

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

八成时间都在跟字符串打交道。`strings.TrimSpace` 去空格、`Split/Join` 切了再拼、`ReplaceAll` 全换；`strconv.Atoi/Itoa` 做字符串和数字互转。坑注释里都标了。

**示例 8 —— 文件读写（os / bufio）：**

```powershell
go run ./files
```

```text
文件读写演示（数据存在 .notes.txt ）
读到了 3 行：
  1. 买牛奶
  2. 写代码
  3. 散步
含"代码"的有 1 行：[写代码]
```

`os.OpenFile` 用 `O_APPEND|O_CREATE` 就能追加写；`bufio.Scanner` 按行读最省事。文件不存在时 `os.Open` 会报 `os.ErrNotExist`，处理一下当空文件就行。

**示例 9 —— 时间处理（time）：**

```powershell
go run ./timepkg
```

```text
现在： 2026-08-06 17:20:00
90 秒后： 17:21:30
差了： 90 秒
睡了 1 秒
解析出来的星期： Thursday
```

`time.Now()` 拿当前时间；格式化要用 Go 的"参考时间" `2006-01-02 15:04:05` 当模板，别用 `YYYY-MM-DD` 那套。`Add`/`Sub`/`Since` 算时间差，`Sleep` 真会停，`Parse` 按同样模板解析回来。

**示例 10 —— 迷你 Web 服务（net/http）：**

```powershell
go run ./web
```

启动后一直监听 8080，另开终端用 curl 试：

```powershell
curl http://localhost:8080/
curl "http://localhost:8080/hello?name=小明"
curl http://localhost:8080/api/ping
```

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

`http.HandleFunc` 把"网址路径"绑到"处理函数"；处理函数里用 `http.ResponseWriter` 写回、`*http.Request` 读请求（`r.URL.Query()` 取参数）。`http.ListenAndServe(":8080", mux)` 启动后阻塞监听，`Ctrl+C` 停。不想开服务器验证？`web/server_test.go` 用 `httptest` 在内存里打这三个路由，`go test ./web` 就成。

### 第 3 步：跑测试

```powershell
go test ./...
```

```
ok  	go-beginner-guide/basics	0.019s
ok  	go-beginner-guide/structs	0.00Xs
?   	go-beginner-guide/games	[no test files]
?   	go-beginner-guide/stdlib	[no test files]
```

`ok` 是过了；`?` 是这个包还没写测试（正常，不用慌）。

## 各示例学什么

| 文件 | 知识点 | 关键函数 |
|------|--------|----------|
| `basics/basics.go` | 变量三种声明、const、if-else、for、switch、函数 | `add()` |
| `stdlib/stdlib.go` | 命令行参数、字符串转数字、格式化输出 | `demoOS()` `demoStrconv()` `demoFmt()` |
| `games/guess.go` | 随机数、缓冲读输入、循环+条件 | `guessNumber()` |
| `structs/roster.go` | struct、方法、切片排序、拼接、格式化 | `Add()` `Top()` `Best()` `AverageScore()` |
| `concurrency/worker.go` | goroutine、channel、WaitGroup、关通道 | `RunConcurrency()` `worker()` |
| `errors/errors.go` | 错误当返回值、errors.New、fmt.Errorf | `divide()` `sqrt()` |
| `text/text.go` | strings 切/拼/换/前后缀、strconv 互转 | `RunText()` |
| `files/files.go` | os 读写、bufio 按行、过滤、持久化 | `appendNote()` `readNotes()` |
| `timepkg/timepkg.go` | 当前时间、参考时间格式化、时间差、Sleep、Parse | `demoNow()` `demoSleep()` `demoParse()` |
| `web/server.go` | 路由注册、读写请求响应、启服务、查参数 | `homeHandler()` `helloHandler()` `apiHandler()` `setupRoutes()` |
| `todo/todo.go` | struct、slice、map 思路、JSON 持久化、flag | `add()` `list()` `done()` `rm()` `clear()` |

建议顺序：先 `basics` → `stdlib` → 挑战 `games` → 再 `structs` → 最后 `todo`。

## 那个能天天用的 todo 小工具

把前面学的串起来：用 `struct` 表示一条待办，用 `slice` 存很多条，用 `json` 文件存盘（关程序也不丢），用 `flag` 解析命令。

```powershell
go run ./todo add "学习 Go 语言"
go run ./todo add "去超市买菜"
go run ./todo list
go run ./todo done 1
go run ./todo rm 2
go run ./todo clear
go run ./todo help
```

```
[x] 1. 学习 Go 语言
[ ] 2. 去超市买菜
```

（`[x]` 完成，`[ ]` 没完成）

数据在项目目录的 `.todo.json`（已被 .gitignore 忽略，不会上传）。

## 常见问题

**Q1：`go run` 报 `go: command not found`？**
A：Go 没装好或没进 PATH。重装时勾"Add to PATH"。

**Q2：猜数字游戏卡住不动？**
A：它在等你输数字回车。输入一个数字敲回车，直到猜中。

**Q3：改了代码重跑没变化？**
A：`go run` 每次重新编译，一般没缓存问题；怀疑就关掉重跑。

**Q4：想加自己的练习？**
A：对应目录下新建 `xxx.go`，包名写 `package main` 并写 `func main()`，就能独立跑。
