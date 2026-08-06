package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// 用 httptest 在内存里启动一个"假服务器"打路由，不用真占 8080 端口，
// 这样测试随时可跑、不互相干扰。
func TestRoutes(t *testing.T) {
	// 重置计数器，保证测试可重复。
	visitCount = 0
	mux := setupRoutes()

	// 1) 首页 "/"
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("首页状态码应为 200，实际 %d", rec.Code)
	}
	body, _ := io.ReadAll(rec.Body)
	if !strings.Contains(string(body), "欢迎来到迷你 Web 服务") {
		t.Errorf("首页内容缺失，实际: %s", body)
	}

	// 2) /hello?name=小明
	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/hello?name=小明", nil)
	mux.ServeHTTP(rec, req)
	body, _ = io.ReadAll(rec.Body)
	if !strings.Contains(string(body), "你好，小明！") {
		t.Errorf("/hello 内容缺失，实际: %s", body)
	}

	// 3) /api/ping 返回 JSON
	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/ping", nil)
	mux.ServeHTTP(rec, req)
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Errorf("/api/ping 应为 JSON 响应，实际 Content-Type: %s", ct)
	}
	body, _ = io.ReadAll(rec.Body)
	if !strings.Contains(string(body), `"status":"ok"`) {
		t.Errorf("/api/ping 内容缺失，实际: %s", body)
	}
}
