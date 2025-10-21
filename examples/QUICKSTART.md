# 🚀 快速开始指南

5分钟上手 Casdoor API Demo!

## 第一步: 配置 Casdoor 应用

### 1. 启动 Casdoor 服务

确保 Casdoor 服务正在运行:
```bash
# 在 casdoor 主目录
go run main.go
```

访问 `http://localhost:8000` 确认服务正常

### 2. 创建或配置应用

1. 登录 Casdoor 管理界面 (默认账号: admin / 123)
2. 进入 **应用** (Applications) 页面
3. 编辑 `app-built-in` 应用 (或创建新应用)
4. 配置以下信息:

**重要配置:**
- **Redirect URLs**: 添加 `http://localhost:9000/callback`
- **Grant types**: 勾选 `Authorization code`

5. 记录以下信息:
   - **Client ID** (例如: 294b09fbc17f95daf2fe)
   - **Client secret** (点击眼睛图标查看)

### 3. 获取证书

1. 进入 **证书** (Certs) 页面
2. 找到 `cert-built-in` 证书
3. 复制证书内容 (PEM 格式)

## 第二步: 配置 Demo

编辑 `examples/config.go` 文件:

```go
func GetDefaultConfig() *Config {
	return &Config{
		Endpoint:         "http://localhost:8000",
		ClientId:         "你的-Client-ID",          // 替换这里
		ClientSecret:     "你的-Client-Secret",      // 替换这里
		OrganizationName: "built-in",
		ApplicationName:  "app-built-in",
		Certificate: `-----BEGIN CERTIFICATE-----
你的证书内容
-----END CERTIFICATE-----`,  // 替换这里
	}
}
```

## 第三步: 运行演示

### 演示 1: OAuth 2.0 登录流程 (推荐先试这个!)

```bash
cd examples
go run . server
```

然后在浏览器打开: `http://localhost:9000`

点击"使用 Casdoor 登录"按钮，体验完整的 OAuth 流程!

### 演示 2: 用户管理

```bash
go run . user
```

查看用户的增删改查操作

### 演示 3: Token 验证

首先通过演示1获取一个 token，然后:

```bash
go run . token <你的-access-token>
```

### 演示 4: API 调用

```bash
go run . api
```

查看各种 Casdoor API 的调用示例

## 💡 常见问题

### Q: 提示"无法连接到 Casdoor 服务"

**A:** 检查 Casdoor 服务是否运行在 `http://localhost:8000`

### Q: OAuth 回调报错

**A:** 确保在应用配置中添加了回调 URL: `http://localhost:9000/callback`

### Q: Token 解析失败

**A:** 检查 `config.go` 中的证书配置是否正确

## 📚 下一步

- 阅读 `README.md` 了解更多详情
- 查看源代码学习实现原理
- 将 Demo 代码集成到你的项目中

---

**祝你使用愉快! 🎉**

