package main

import (
	"fmt"
	"github.com/Chabuduo04/User_management_system/internal/model"
	"github.com/Chabuduo04/User_management_system/internal/service"
	"github.com/Chabuduo04/User_management_system/internal/storage"
)

func main() {
	// 初始化依赖
	memStorage := storage.NewMemoryStorage()
	userService := service.NewUserService(memStorage)

	// 模拟用户注册
	user, err := userService.Register(&model.CreateUserRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "securepassword123",
		Roles:    []string{"admin"},
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Registered user: %+v\n", user)
}