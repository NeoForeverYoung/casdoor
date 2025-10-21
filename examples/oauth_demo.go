package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// handleHome 首页处理
func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Casdoor OAuth Demo</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            max-width: 800px;
            margin: 50px auto;
            padding: 20px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
        }
        .container {
            background: white;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 10px 40px rgba(0,0,0,0.2);
        }
        h1 {
            color: #333;
            text-align: center;
            margin-bottom: 30px;
        }
        .description {
            color: #666;
            line-height: 1.8;
            margin-bottom: 30px;
            padding: 20px;
            background: #f8f9fa;
            border-left: 4px solid #667eea;
            border-radius: 4px;
        }
        .login-btn {
            display: block;
            width: 100%;
            padding: 15px 30px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            text-decoration: none;
            border-radius: 5px;
            text-align: center;
            font-size: 18px;
            font-weight: bold;
            transition: transform 0.2s, box-shadow 0.2s;
        }
        .login-btn:hover {
            transform: translateY(-2px);
            box-shadow: 0 5px 20px rgba(102, 126, 234, 0.4);
        }
        .features {
            margin-top: 30px;
            padding: 20px;
            background: #f8f9fa;
            border-radius: 5px;
        }
        .features h3 {
            color: #333;
            margin-bottom: 15px;
        }
        .features ul {
            list-style: none;
            padding: 0;
        }
        .features li {
            padding: 8px 0;
            color: #555;
        }
        .features li:before {
            content: "✓ ";
            color: #667eea;
            font-weight: bold;
            margin-right: 10px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🔐 Casdoor OAuth 2.0 演示</h1>
        
        <div class="description">
            <p><strong>欢迎使用 Casdoor API Demo!</strong></p>
            <p>这个演示程序展示了如何使用 Casdoor 实现 OAuth 2.0 授权码模式登录。</p>
            <p>点击下方按钮开始体验完整的登录流程。</p>
        </div>
        
        <a href="/login" class="login-btn">🚀 使用 Casdoor 登录</a>
        
        <div class="features">
            <h3>本 Demo 演示的功能:</h3>
            <ul>
                <li>OAuth 2.0 授权码模式</li>
                <li>获取 Access Token</li>
                <li>获取用户信息</li>
                <li>JWT Token 解析</li>
                <li>用户信息展示</li>
            </ul>
        </div>
    </div>
</body>
</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

// handleLogin 处理登录请求，重定向到 Casdoor
func handleLogin(w http.ResponseWriter, r *http.Request) {
	// 生成登录 URL
	redirectURI := "http://localhost:9000/callback"

	// 使用 SDK 生成授权 URL
	authURL := casdoorsdk.GetSigninUrl(redirectURI)

	fmt.Printf("重定向到 Casdoor: %s\n", authURL)
	http.Redirect(w, r, authURL, http.StatusFound)
}

// handleCallback 处理 OAuth 回调
func handleCallback(w http.ResponseWriter, r *http.Request) {
	// 获取授权码
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code == "" {
		http.Error(w, "未获取到授权码", http.StatusBadRequest)
		return
	}

	fmt.Printf("收到授权码: %s, state: %s\n", code, state)

	// 使用授权码换取 token
	token, err := casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		http.Error(w, fmt.Sprintf("获取 token 失败: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("获取到 Access Token: %s\n", token.AccessToken)

	// 解析 token 获取用户信息
	claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("解析 token 失败: %v", err), http.StatusInternalServerError)
		return
	}

	// 显示用户信息
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>登录成功</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            max-width: 900px;
            margin: 50px auto;
            padding: 20px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
        }
        .container {
            background: white;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 10px 40px rgba(0,0,0,0.2);
        }
        h1 {
            color: #28a745;
            text-align: center;
            margin-bottom: 30px;
        }
        .success-icon {
            text-align: center;
            font-size: 80px;
            margin-bottom: 20px;
        }
        .info-section {
            background: #f8f9fa;
            padding: 20px;
            margin: 20px 0;
            border-radius: 5px;
            border-left: 4px solid #667eea;
        }
        .info-section h3 {
            color: #333;
            margin-top: 0;
            margin-bottom: 15px;
        }
        .info-item {
            margin: 10px 0;
            padding: 8px 0;
            border-bottom: 1px solid #e9ecef;
        }
        .info-item:last-child {
            border-bottom: none;
        }
        .info-label {
            font-weight: bold;
            color: #555;
            display: inline-block;
            min-width: 150px;
        }
        .info-value {
            color: #333;
            word-break: break-all;
        }
        .token-box {
            background: #f8f9fa;
            padding: 15px;
            border-radius: 5px;
            margin: 10px 0;
            font-family: monospace;
            font-size: 12px;
            word-break: break-all;
            max-height: 150px;
            overflow-y: auto;
            border: 1px solid #dee2e6;
        }
        .btn {
            display: inline-block;
            padding: 10px 20px;
            margin: 10px 5px;
            border-radius: 5px;
            text-decoration: none;
            transition: all 0.2s;
        }
        .btn-primary {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
        }
        .btn-secondary {
            background: #6c757d;
            color: white;
        }
        .btn:hover {
            transform: translateY(-2px);
            box-shadow: 0 5px 15px rgba(0,0,0,0.2);
        }
        .actions {
            text-align: center;
            margin-top: 30px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="success-icon">✅</div>
        <h1>登录成功!</h1>
        
        <div class="info-section">
            <h3>👤 用户信息</h3>
            <div class="info-item">
                <span class="info-label">用户名:</span>
                <span class="info-value">%s</span>
            </div>
            <div class="info-item">
                <span class="info-label">显示名称:</span>
                <span class="info-value">%s</span>
            </div>
            <div class="info-item">
                <span class="info-label">所属组织:</span>
                <span class="info-value">%s</span>
            </div>
            <div class="info-item">
                <span class="info-label">邮箱:</span>
                <span class="info-value">%s</span>
            </div>
        </div>
        
        <div class="info-section">
            <h3>🎫 Token 信息</h3>
            <div class="info-item">
                <span class="info-label">Token 类型:</span>
                <span class="info-value">%s</span>
            </div>
            <div class="info-item">
                <span class="info-label">过期时间:</span>
                <span class="info-value">%s</span>
            </div>
            <div class="info-item">
                <span class="info-label">Access Token:</span>
                <div class="token-box">%s</div>
            </div>
        </div>
        
        <div class="info-section">
            <h3>📋 JWT Claims</h3>
            <div class="info-item">
                <span class="info-label">Subject (sub):</span>
                <span class="info-value">%s</span>
            </div>
            <div class="info-item">
                <span class="info-label">Issuer (iss):</span>
                <span class="info-value">%s</span>
            </div>
            <div class="info-item">
                <span class="info-label">用户类型:</span>
                <span class="info-value">%s</span>
            </div>
        </div>
        
        <div class="actions">
            <a href="/" class="btn btn-primary">返回首页</a>
            <a href="/logout" class="btn btn-secondary">退出登录</a>
        </div>
    </div>
</body>
</html>
	`,
		claims.Name,
		claims.DisplayName,
		claims.Owner,
		getEmail(claims),
		token.TokenType,
		token.Expiry.Format(time.RFC3339),
		token.AccessToken,
		claims.Subject,
		claims.Issuer,
		claims.Type,
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

// handleLogout 处理退出登录
func handleLogout(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>已退出</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            max-width: 600px;
            margin: 100px auto;
            padding: 20px;
            text-align: center;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
        }
        .container {
            background: white;
            padding: 60px 40px;
            border-radius: 10px;
            box-shadow: 0 10px 40px rgba(0,0,0,0.2);
        }
        h1 {
            color: #333;
            margin-bottom: 30px;
        }
        .icon {
            font-size: 80px;
            margin-bottom: 20px;
        }
        .btn {
            display: inline-block;
            padding: 12px 30px;
            margin-top: 20px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            text-decoration: none;
            border-radius: 5px;
            font-size: 16px;
            transition: all 0.2s;
        }
        .btn:hover {
            transform: translateY(-2px);
            box-shadow: 0 5px 15px rgba(102, 126, 234, 0.4);
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="icon">👋</div>
        <h1>您已成功退出</h1>
        <p>感谢使用 Casdoor OAuth Demo</p>
        <a href="/" class="btn">重新登录</a>
    </div>
</body>
</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

// 辅助函数
func getEmail(claims *casdoorsdk.Claims) string {
	if claims.Email != "" {
		return claims.Email
	}
	return "未设置"
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[i%len(charset)]
	}
	return string(result)
}
