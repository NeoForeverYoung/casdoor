package main

import (
	"fmt"
	"time"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// runUserDemo 演示用户管理操作
func runUserDemo() {
	fmt.Println("========================================")
	fmt.Println("Casdoor 用户管理 API 演示")
	fmt.Println("========================================")
	fmt.Println()

	// 1. 获取所有用户
	fmt.Println("📋 1. 获取用户列表")
	fmt.Println("---")
	users, err := casdoorsdk.GetUsers()
	if err != nil {
		fmt.Printf("❌ 获取用户失败: %v\n", err)
		return
	}
	fmt.Printf("✅ 成功获取 %d 个用户\n", len(users))
	for i, user := range users {
		if i < 5 { // 只显示前5个
			fmt.Printf("   - %s (%s) - %s\n", user.Name, user.DisplayName, user.Email)
		}
	}
	if len(users) > 5 {
		fmt.Printf("   ... 还有 %d 个用户\n", len(users)-5)
	}
	fmt.Println()

	// 2. 创建新用户
	fmt.Println("➕ 2. 创建新用户")
	fmt.Println("---")

	newUser := &casdoorsdk.User{
		Owner:       "built-in",
		Name:        fmt.Sprintf("demo-user-%d", time.Now().Unix()),
		CreatedTime: time.Now().Format(time.RFC3339),
		DisplayName: "Demo User",
		Email:       fmt.Sprintf("demo-%d@example.com", time.Now().Unix()),
		Phone:       "13812345678",
		Type:        "normal-user",
		Password:    "Demo@123456",
		IsAdmin:     false,
	}

	affected, err := casdoorsdk.AddUser(newUser)
	if err != nil {
		fmt.Printf("❌ 创建用户失败: %v\n", err)
	} else if affected {
		fmt.Printf("✅ 成功创建用户: %s\n", newUser.Name)
		fmt.Printf("   - 显示名称: %s\n", newUser.DisplayName)
		fmt.Printf("   - 邮箱: %s\n", newUser.Email)
		fmt.Printf("   - 手机: %s\n", newUser.Phone)
	} else {
		fmt.Println("⚠️  用户可能已存在或创建失败")
	}
	fmt.Println()

	// 3. 获取单个用户信息
	fmt.Println("🔍 3. 查询用户详情")
	fmt.Println("---")
	userName := newUser.Name
	user, err := casdoorsdk.GetUser(userName)
	if err != nil {
		fmt.Printf("❌ 获取用户失败: %v\n", err)
	} else if user != nil {
		fmt.Printf("✅ 成功获取用户: %s\n", user.Name)
		fmt.Printf("   - ID: %s\n", user.Id)
		fmt.Printf("   - 所属组织: %s\n", user.Owner)
		fmt.Printf("   - 创建时间: %s\n", user.CreatedTime)
		fmt.Printf("   - 显示名称: %s\n", user.DisplayName)
		fmt.Printf("   - 邮箱: %s\n", user.Email)
		fmt.Printf("   - 是否管理员: %v\n", user.IsAdmin)
	}
	fmt.Println()

	// 4. 更新用户
	fmt.Println("✏️  4. 更新用户信息")
	fmt.Println("---")
	if user != nil {
		user.DisplayName = "Updated Demo User"
		user.Phone = "13987654321"

		affected, err := casdoorsdk.UpdateUser(user)
		if err != nil {
			fmt.Printf("❌ 更新用户失败: %v\n", err)
		} else if affected {
			fmt.Printf("✅ 成功更新用户: %s\n", user.Name)
			fmt.Printf("   - 新显示名称: %s\n", user.DisplayName)
			fmt.Printf("   - 新手机号: %s\n", user.Phone)
		}
	}
	fmt.Println()

	// 5. 统计信息
	fmt.Println("📊 5. 统计信息")
	fmt.Println("---")
	count, err := casdoorsdk.GetUserCount("built-in")
	if err != nil {
		fmt.Printf("❌ 获取用户数量失败: %v\n", err)
	} else {
		fmt.Printf("✅ 组织 'built-in' 共有 %d 个用户\n", count)
	}
	fmt.Println()

	// 6. 删除用户（可选，注释掉以避免误删）
	fmt.Println("🗑️  6. 删除演示用户")
	fmt.Println("---")
	if user != nil {
		affected, err := casdoorsdk.DeleteUser(user)
		if err != nil {
			fmt.Printf("❌ 删除用户失败: %v\n", err)
		} else if affected {
			fmt.Printf("✅ 成功删除用户: %s\n", user.Name)
		}
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("用户管理演示完成!")
	fmt.Println("========================================")
}
