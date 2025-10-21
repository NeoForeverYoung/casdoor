package main

import (
	"fmt"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// runAPIDemo 演示各种 API 调用
func runAPIDemo() {
	fmt.Println("========================================")
	fmt.Println("Casdoor API 综合演示")
	fmt.Println("========================================")
	fmt.Println()

	// 1. 组织管理
	fmt.Println("🏢 1. 组织管理")
	fmt.Println("---")
	organizations, err := casdoorsdk.GetOrganizations()
	if err != nil {
		fmt.Printf("❌ 获取组织列表失败: %v\n", err)
	} else {
		fmt.Printf("✅ 获取到 %d 个组织:\n", len(organizations))
		for _, org := range organizations {
			fmt.Printf("   - %s (%s)\n", org.Name, org.DisplayName)
		}
	}
	fmt.Println()

	// 2. 应用管理
	fmt.Println("📱 2. 应用管理")
	fmt.Println("---")
	applications, err := casdoorsdk.GetApplications()
	if err != nil {
		fmt.Printf("❌ 获取应用列表失败: %v\n", err)
	} else {
		fmt.Printf("✅ 获取到 %d 个应用:\n", len(applications))
		for _, app := range applications {
			fmt.Printf("   - %s (%s)\n", app.Name, app.DisplayName)
			fmt.Printf("     Client ID: %s\n", app.ClientId)
			fmt.Printf("     Homepage: %s\n", app.HomepageUrl)
		}
	}
	fmt.Println()

	// 3. 角色管理
	fmt.Println("👥 3. 角色管理")
	fmt.Println("---")
	roles, err := casdoorsdk.GetRoles()
	if err != nil {
		fmt.Printf("❌ 获取角色列表失败: %v\n", err)
	} else {
		fmt.Printf("✅ 获取到 %d 个角色:\n", len(roles))
		for _, role := range roles {
			fmt.Printf("   - %s (%s)\n", role.Name, role.DisplayName)
		}
	}
	fmt.Println()

	// 4. 权限管理
	fmt.Println("🔐 4. 权限管理")
	fmt.Println("---")
	permissions, err := casdoorsdk.GetPermissions()
	if err != nil {
		fmt.Printf("❌ 获取权限列表失败: %v\n", err)
	} else {
		fmt.Printf("✅ 获取到 %d 个权限:\n", len(permissions))
		for i, perm := range permissions {
			if i < 5 { // 只显示前5个
				fmt.Printf("   - %s: %s\n", perm.Name, perm.DisplayName)
				fmt.Printf("     Resources: %v\n", perm.Resources)
				fmt.Printf("     Actions: %v\n", perm.Actions)
			}
		}
		if len(permissions) > 5 {
			fmt.Printf("   ... 还有 %d 个权限\n", len(permissions)-5)
		}
	}
	fmt.Println()

	// 5. Provider 管理 (第三方登录提供商)
	fmt.Println("🔗 5. 第三方登录提供商")
	fmt.Println("---")
	providers, err := casdoorsdk.GetProviders()
	if err != nil {
		fmt.Printf("❌ 获取提供商列表失败: %v\n", err)
	} else {
		fmt.Printf("✅ 获取到 %d 个提供商:\n", len(providers))
		for _, provider := range providers {
			fmt.Printf("   - %s (%s) - 类型: %s\n",
				provider.Name, provider.DisplayName, provider.Type)
		}
	}
	fmt.Println()

	// 6. 资源管理
	fmt.Println("📦 6. 资源管理")
	fmt.Println("---")
	// GetResources 需要参数，这里获取 built-in 组织的资源
	resources, err := casdoorsdk.GetResources("built-in", "", "", "", "", "")
	if err != nil {
		fmt.Printf("❌ 获取资源列表失败: %v\n", err)
	} else {
		fmt.Printf("✅ 获取到 %d 个资源\n", len(resources))
		for i, res := range resources {
			if i < 3 {
				fmt.Printf("   - %s\n", res.Name)
			}
		}
		if len(resources) > 3 {
			fmt.Printf("   ... 还有 %d 个资源\n", len(resources)-3)
		}
	}
	fmt.Println()

	// 7. 获取当前用户信息 (需要 token)
	fmt.Println("👤 7. 用户认证演示")
	fmt.Println("---")
	fmt.Println("💡 提示: 可以使用 'server' 命令启动 OAuth 服务器进行完整的登录流程演示")
	fmt.Println()

	// 8. URL 生成演示
	fmt.Println("🔗 8. 重要 URL 生成")
	fmt.Println("---")
	signupUrl := casdoorsdk.GetSignupUrl(true, "http://localhost:9000/callback")
	fmt.Printf("注册 URL: %s\n", signupUrl)

	signinUrl := casdoorsdk.GetSigninUrl("http://localhost:9000/callback")
	fmt.Printf("登录 URL: %s\n", signinUrl)

	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("API 调用演示完成!")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("💡 提示:")
	fmt.Println("   - 使用 'go run . server' 启动 OAuth Web 演示")
	fmt.Println("   - 使用 'go run . user' 演示用户管理功能")
	fmt.Println("   - 使用 'go run . token <token>' 验证 JWT token")
}
