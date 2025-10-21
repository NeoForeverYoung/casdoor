# Casdoor API Demo - 使用示例

这是一个完整的 Casdoor API 使用演示项目，展示了如何使用 Casdoor SDK 实现各种常见的认证和授权功能。

## 📋 目录

- [功能特性](#功能特性)
- [前置条件](#前置条件)
- [安装配置](#安装配置)
- [使用说明](#使用说明)
- [项目结构](#项目结构)
- [常见问题](#常见问题)

## ✨ 功能特性

本 Demo 包含以下功能演示:

1. **OAuth 2.0 授权码流程** - 完整的 Web 登录演示
2. **用户管理** - 增删改查用户信息
3. **Token 验证** - JWT Token 解析和验证
4. **API 调用** - 组织、应用、角色、权限等 API 调用

## 🔧 前置条件

1. **Go 环境**: Go 1.20 或更高版本
2. **Casdoor 服务**: 需要运行 Casdoor 服务器
   - 默认地址: `http://localhost:8000`
   - 如使用其他地址，请修改 `config.go` 中的配置

3. **Casdoor 应用配置**:
   - 需要在 Casdoor 管理界面创建应用
   - 获取 Client ID 和 Client Secret
   - 配置回调 URL: `http://localhost:9000/callback`

## 📦 安装配置

### 1. 下载依赖

```bash
cd examples
go mod download
```

### 2. 配置应用信息

编辑 `config.go` 文件，填入你的 Casdoor 配置:

```go
type Config struct {
    Endpoint         string // Casdoor 服务地址
    ClientId         string // 应用 Client ID
    ClientSecret     string // 应用 Client Secret
    Certificate      string // JWT 验证证书
    OrganizationName string // 组织名称
    ApplicationName  string // 应用名称
}
```

### 3. 在 Casdoor 中配置应用

1. 登录 Casdoor 管理界面
2. 进入 "应用" 管理页面
3. 创建或编辑应用，配置以下信息:
   - **Redirect URLs**: 添加 `http://localhost:9000/callback`
   - **Grant Types**: 勾选 `Authorization Code`
   - 复制 **Client ID** 和 **Client Secret**
4. 更新 `config.go` 中的配置

## 🚀 使用说明

### 命令格式

```bash
go run . <command> [args]
```

### 可用命令

#### 1. OAuth 2.0 Web 演示服务器

启动一个 Web 服务器，演示完整的 OAuth 2.0 授权码流程:

```bash
go run . server
```

然后在浏览器访问: `http://localhost:9000`

**演示流程:**
1. 点击 "使用 Casdoor 登录"
2. 跳转到 Casdoor 登录页面
3. 输入用户名密码登录
4. 授权后返回并展示用户信息和 Token

#### 2. 用户管理演示

演示用户的增删改查操作:

```bash
go run . user
```

**演示内容:**
- 获取用户列表
- 创建新用户
- 查询用户详情
- 更新用户信息
- 删除用户
- 统计用户数量

#### 3. Token 验证演示

验证并解析 JWT Token:

```bash
go run . token <your-access-token>
```

**示例:**
```bash
go run . token eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJvd25lciI6ImJ1aWx0LWluIiwibmFtZSI6ImFkbWluIiwiY3JlYXRlZFRpbWUiOiIyMDIxLTA1LTI4VDA2OjU1OjA2KzAwOjAwIiwidXBkYXRlZFRpbWUiOiIyMDI0LTEwLTIwVDEyOjAwOjAwKzAwOjAwIiwiaWQiOiJhZG1pbiIsInR5cGUiOiJub3JtYWwtdXNlciIsInBhc3N3b3JkIjoiIiwicGFzc3dvcmRTYWx0IjoiIiwicGFzc3dvcmRUeXBlIjoicGxhaW4iLCJkaXNwbGF5TmFtZSI6IkFkbWluIiwiZmlyc3ROYW1lIjoiIiwibGFzdE5hbWUiOiIiLCJhdmF0YXIiOiJodHRwczovL2Nkbi5jYXNiaW4ub3JnL2ltZy9jYXNkb29yLWxvZ28uYXZpZiIsImF2YXRhclR5cGUiOiIiLCJwZXJtYW5lbnRBdmF0YXIiOiIiLCJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwic...
```

**显示信息:**
- 用户基本信息
- Token 元信息 (签发时间、过期时间等)
- 权限和 Scope
- 完整的 JWT Claims
- Token 结构分析

#### 4. API 综合演示

演示各种 Casdoor API 的调用:

```bash
go run . api
```

**演示内容:**
- 组织管理 API
- 应用管理 API
- 角色管理 API
- 权限管理 API
- Provider 管理 API
- 资源管理 API
- URL 生成工具

## 📁 项目结构

```
examples/
├── go.mod              # Go 模块定义
├── go.sum              # 依赖锁定文件
├── README.md           # 项目说明文档 (本文件)
├── config.go           # Casdoor 配置
├── main.go             # 主程序入口
├── oauth_demo.go       # OAuth 2.0 演示
├── user_demo.go        # 用户管理演示
├── token_demo.go       # Token 验证演示
└── api_demo.go         # API 调用演示
```

## 💡 代码示例

### 初始化 Casdoor SDK

```go
import "github.com/casdoor/casdoor-go-sdk/casdoorsdk"

casdoorsdk.InitConfig(
    "http://localhost:8000",     // endpoint
    "294b09fbc17f95daf2fe",      // clientId
    "dd8982f7046ccba1bbd7851d5c1ece4e52bf039d", // clientSecret
    certificate,                  // certificate
    "built-in",                   // organizationName
    "app-built-in",              // applicationName
)
```

### 生成登录 URL

```go
redirectURI := "http://localhost:9000/callback"
state := "random_state_string"
authURL := casdoorsdk.GetSigninUrl(redirectURI, state)
// 重定向用户到 authURL
```

### 用授权码换取 Token

```go
code := r.URL.Query().Get("code")
state := r.URL.Query().Get("state")

token, err := casdoorsdk.GetOAuthToken(code, state)
if err != nil {
    // 处理错误
}

// token.AccessToken - 访问令牌
// token.IdToken - ID 令牌
// token.RefreshToken - 刷新令牌
// token.ExpiresIn - 过期时间(秒)
```

### 验证和解析 Token

```go
claims, err := casdoorsdk.ParseJwtToken(accessToken)
if err != nil {
    // Token 无效或过期
}

// 使用用户信息
fmt.Println("用户名:", claims.Name)
fmt.Println("邮箱:", claims.Email)
fmt.Println("组织:", claims.Owner)
```

### 用户管理

```go
// 获取所有用户
users, err := casdoorsdk.GetUsers()

// 获取单个用户
user, err := casdoorsdk.GetUser("username")

// 创建用户
newUser := &casdoorsdk.User{
    Owner:       "built-in",
    Name:        "john",
    DisplayName: "John Doe",
    Email:       "john@example.com",
    Password:    "Pass@123",
}
affected, err := casdoorsdk.AddUser(newUser)

// 更新用户
user.DisplayName = "Jane Doe"
affected, err := casdoorsdk.UpdateUser(user)

// 删除用户
affected, err := casdoorsdk.DeleteUser(user)
```

## ❓ 常见问题

### 1. 无法连接到 Casdoor 服务

**问题**: 运行时提示连接失败

**解决方案**:
- 确保 Casdoor 服务正在运行: `http://localhost:8000`
- 检查 `config.go` 中的 `Endpoint` 配置是否正确
- 尝试在浏览器访问 Casdoor 管理界面确认服务可用

### 2. OAuth 回调失败

**问题**: 登录后回调 URL 报错

**解决方案**:
- 在 Casdoor 应用配置中添加回调 URL: `http://localhost:9000/callback`
- 确保 ClientId 和 ClientSecret 配置正确
- 检查应用的 Grant Types 是否包含 "Authorization Code"

### 3. Token 验证失败

**问题**: 解析 Token 时报错

**解决方案**:
- 确保提供的是完整的 JWT Token
- 检查证书配置是否正确
- Token 可能已过期，尝试重新登录获取新 Token

### 4. API 调用返回 401/403

**问题**: 调用 API 时返回未授权错误

**解决方案**:
- 确保使用了正确的 ClientId 和 ClientSecret
- 检查用户是否有相应的权限
- 某些 API 可能需要管理员权限

### 5. 依赖下载失败

**问题**: `go mod download` 失败

**解决方案**:
```bash
# 配置 Go 代理
export GOPROXY=https://goproxy.cn,direct
go mod download
```

## 📚 学习资源

- [Casdoor 官方文档](https://casdoor.org/docs/overview)
- [Casdoor Go SDK](https://github.com/casdoor/casdoor-go-sdk)
- [OAuth 2.0 RFC 6749](https://tools.ietf.org/html/rfc6749)
- [OpenID Connect](https://openid.net/connect/)
- [JWT.io](https://jwt.io) - JWT Token 在线解析工具

## 📝 下一步

学完本 Demo 后，你可以:

1. **集成到自己的应用**
   - 参考 `oauth_demo.go` 实现登录功能
   - 使用 Token 保护你的 API

2. **探索更多功能**
   - SAML 单点登录
   - LDAP 集成
   - 多因素认证 (MFA)
   - Webhook 通知

3. **参与开发**
   - 查看 Casdoor 源代码
   - 贡献代码或提交 Issue
   - 加入社区讨论

## 🤝 贡献

欢迎提交 Issue 和 Pull Request!

## 📄 许可证

Apache License 2.0

---

**祝你使用愉快! 🎉**

如有问题，请参考 [Casdoor 官方文档](https://casdoor.org) 或提交 Issue。

