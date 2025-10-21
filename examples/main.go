package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 初始化配置
	config := GetDefaultConfig()
	InitCasdoorSDK(config)

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "server":
		// 启动 OAuth 演示服务器
		runOAuthServer()
	case "user":
		// 用户管理演示
		runUserDemo()
	case "token":
		// Token 验证演示
		if len(os.Args) < 3 {
			fmt.Println("请提供 token: go run . token <your-token>")
			return
		}
		runTokenDemo(os.Args[2])
	case "api":
		// API 调用演示
		runAPIDemo()
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Print(`
Casdoor API Demo - 使用示例

用法:
  go run . <command> [args]

命令:
  server    - 启动 OAuth 2.0 演示服务器 (监听 http://localhost:9000)
  user      - 演示用户管理操作 (创建、查询、更新、删除用户)
  token     - 验证并解析 JWT token
  api       - 演示各种 API 调用

示例:
  go run . server                           # 启动 Web 服务器演示 OAuth 流程
  go run . user                             # 演示用户管理功能
  go run . token eyJhbGc...                 # 验证 token
  go run . api                              # 演示 API 调用

说明:
  1. 启动前请确保 Casdoor 服务运行在 http://localhost:8000
  2. 在 config.go 中配置正确的 ClientId 和 ClientSecret
  3. OAuth 演示需要在浏览器中访问 http://localhost:9000
`)
}

func runOAuthServer() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/logout", handleLogout)

	fmt.Println("========================================")
	fmt.Println("OAuth 2.0 演示服务器已启动")
	fmt.Println("访问: http://localhost:9000")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("演示流程:")
	fmt.Println("1. 访问首页")
	fmt.Println("2. 点击 '使用 Casdoor 登录'")
	fmt.Println("3. 在 Casdoor 登录页面输入用户名密码")
	fmt.Println("4. 授权后返回查看用户信息")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":9000", nil))
}
