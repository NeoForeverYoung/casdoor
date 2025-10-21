package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/golang-jwt/jwt/v4"
)

// runTokenDemo 演示 Token 验证和解析
func runTokenDemo(tokenString string) {
	fmt.Println("========================================")
	fmt.Println("JWT Token 验证与解析演示")
	fmt.Println("========================================")
	fmt.Println()

	// 1. 使用 Casdoor SDK 解析 Token
	fmt.Println("🔐 1. 使用 Casdoor SDK 解析 Token")
	fmt.Println("---")
	claims, err := casdoorsdk.ParseJwtToken(tokenString)
	if err != nil {
		fmt.Printf("❌ Token 解析失败: %v\n", err)
		fmt.Println()
		fmt.Println("提示: 请确保提供的是有效的 Casdoor JWT Token")
		return
	}

	fmt.Println("✅ Token 验证成功!")
	fmt.Println()

	// 2. 显示用户信息
	fmt.Println("👤 2. 用户信息")
	fmt.Println("---")
	fmt.Printf("用户名: %s\n", claims.Name)
	fmt.Printf("显示名称: %s\n", claims.DisplayName)
	fmt.Printf("所属组织: %s\n", claims.Owner)
	fmt.Printf("邮箱: %s\n", claims.Email)
	fmt.Printf("Subject: %s\n", claims.Subject)
	fmt.Println()

	// 3. 显示 Token 元信息
	fmt.Println("🎫 3. Token 元信息")
	fmt.Println("---")
	fmt.Printf("签发者 (Issuer): %s\n", claims.Issuer)
	fmt.Printf("受众 (Audience): %v\n", claims.Audience)

	if claims.IssuedAt != nil {
		fmt.Printf("签发时间: %s\n", claims.IssuedAt.Time.Format(time.RFC3339))
	}
	if claims.ExpiresAt != nil {
		fmt.Printf("过期时间: %s\n", claims.ExpiresAt.Time.Format(time.RFC3339))

		// 计算剩余有效时间
		now := time.Now()
		if claims.ExpiresAt.Time.After(now) {
			remaining := claims.ExpiresAt.Time.Sub(now)
			fmt.Printf("剩余有效期: %.0f 秒 (约 %.0f 分钟)\n", remaining.Seconds(), remaining.Minutes())
		} else {
			fmt.Println("⚠️  Token 已过期!")
		}
	}
	if claims.NotBefore != nil {
		fmt.Printf("生效时间: %s\n", claims.NotBefore.Time.Format(time.RFC3339))
	}
	fmt.Println()

	// 4. 显示权限信息
	fmt.Println("🔑 4. 权限信息")
	fmt.Println("---")
	fmt.Printf("Token 类型: %s\n", claims.TokenType)
	fmt.Printf("用户类型: %s\n", claims.Type)
	fmt.Println()

	// 5. 显示完整的 Claims
	fmt.Println("📋 5. 完整的 JWT Claims")
	fmt.Println("---")
	claimsJSON, _ := json.MarshalIndent(claims, "", "  ")
	fmt.Println(string(claimsJSON))
	fmt.Println()

	// 6. 手动解析 Token (不验证签名，仅用于学习)
	fmt.Println("🔧 6. Token 结构分析")
	fmt.Println("---")
	parser := jwt.NewParser()
	token, _, err := parser.ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("❌ 解析失败: %v\n", err)
	} else {
		fmt.Printf("算法: %s\n", token.Method.Alg())
		fmt.Printf("Token 类型: %s\n", token.Header["typ"])

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println("\nClaims 字段:")
			for key, value := range claims {
				fmt.Printf("  - %s: %v\n", key, value)
			}
		}
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("Token 分析完成!")
	fmt.Println("========================================")
}
