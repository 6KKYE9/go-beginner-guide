// 示例 10：用 net/http 写一个迷你 Web 服务
//
// 前面 9 个示例都是"在命令行里跑、看文字输出"。本示例带你走出终端，
// 写一个有真实 HTTP 接口的 Web 服务——这是 Go 在后端开发里最常用、也最拿手的事。
//
// 你会学到三件事：
//   1. 用 net/http 注册路由（把"网址路径"对应到"处理函数"）。
//   2. 处理函数怎么读请求、写响应（http.ResponseWriter / *http.Request）。
//   3. 启动服务器并监听端口（http.ListenAndServe）。
//
// 跑起来后，浏览器或 curl 访问 http://localhost:8080 就能看到返回内容。
// 建议先读懂 basics / structs / concurrency 再看本文件。
package main

import (
	"fmt"
	"net/http"
)

// 一个极简的"访客计数器"：记录总共被访问了多少次。
// 注意：这个计数器在并发下并不严谨（没加锁），本示例重点是 HTTP 流程，
// 严谨写法属于进阶内容，先不展开。
var visitCount int

// homeHandler 处理根路径 "/" 的请求，返回一个简单的欢迎页面。
// 参数 w 用来"写回响应"，r 是"收到的请求"。
func homeHandler(w http.ResponseWriter, r *http.Request) {
	visitCount++
	// 设置响应头的 Content-Type，告诉浏览器这是纯文本（而不是要下载的文件）。
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// Fprintf 把格式化字符串直接写进响应体，浏览器/curl 就能收到。
	fmt.Fprintf(w, "欢迎来到迷你 Web 服务！\n你访问的是：%s\n这是第 %d 次被访问。\n", r.URL.Path, visitCount)
}

// helloHandler 处理 "/hello" 路径，演示怎么读取 URL 里的查询参数（?name=xxx）。
// 例如访问 /hello?name=小明，就返回"你好，小明！"。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Query 返回一个 map，Get 按键取值；没传 name 时给个默认值。
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "陌生人"
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "你好，%s！这是 /hello 接口。\n", name)
}

// apiHandler 处理 "/api/ping"，返回一个 JSON 字符串，模拟最常见的"健康检查"接口。
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, `{"status":"ok","message":"pong"}`)
}

// setupRoutes 把"路径"和"处理函数"绑定起来，并返回配置好的路由器。
// 抽成函数，方便测试时直接复用（测试里不会真的占用 8080 端口）。
func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux() // ServeMux 是 Go 内置的"路由表"
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/api/ping", apiHandler)
	return mux
}

func main() {
	// 把路径和处理函数绑定好。
	mux := setupRoutes()

	addr := ":8080"
	fmt.Println("迷你 Web 服务已启动，打开浏览器访问：")
	fmt.Printf("  首页:        http://localhost%s/\n", addr)
	fmt.Printf("  带名字的问候: http://localhost%s/hello?name=小明\n", addr)
	fmt.Printf("  JSON 接口:    http://localhost%s/api/ping\n", addr)
	fmt.Println("（按 Ctrl+C 停止服务）")

	// ListenAndServe 会"阻塞"在这里，一直监听端口等待请求。
	// 第二个参数是路由器；传 nil 会用默认路由器，这里我们传自己配好的 mux。
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Printf("服务启动失败: %v\n", err)
	}
}
