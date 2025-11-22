# Nacos Config Management SDK for Go

[![build workflow](https://github.com/cx0110/nacos-config-go/actions/workflows/main.yml/badge.svg)](https://github.com/cx0110/nacos-config-go/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/cx0110/nacos-config-go)](https://pkg.go.dev/github.com/cx0110/nacos-config-go?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/cx0110/nacos-config-go)](https://goreportcard.com/report/github.com/cx0110/nacos-config-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

> 一个功能完善的 Nacos 配置管理和权限控制 Go SDK，提供简洁易用的 API 来操作 Nacos 配置中心。
> A comprehensive Nacos configuration management and permission control Go SDK that provides simple and easy-to-use APIs for operating Nacos Configuration Center.

## 文档 Documentation

- [简体中文](./README_zh.md)
- [English](./README.md)

## 资源 Resources

- [Go Reference](https://pkg.go.dev/github.com/cx0110/nacos-config-go)
- [Examples](https://pkg.go.dev/github.com/cx0110/nacos-config-go#pkg-examples)

## 特性 Features

### 🔧 配置管理 Configuration Management
- [x] 获取配置 Get Configuration
- [x] 发布配置 Publish Configuration
- [x] 删除配置 Delete Configuration
- [x] 监听配置变化 Listen Configuration Changes
- [x] 查询配置历史 Query Configuration History
- [x] 查询历史详情 Get History Details
- [x] 查询上一版本 Get Previous Version

### 🛡️ 权限管理 Permission Management
- [x] 用户管理 User Management
  - 创建用户 Create User
  - 查询用户列表 List Users
  - 修改用户密码 Update User Password
  - 删除用户 Delete User
- [x] 角色管理 Role Management
  - 创建角色 Create Role
  - 查询角色列表 List Roles
  - 删除角色 Delete Role
- [x] 权限管理 Permissions Management
  - 创建权限 Create Permission
  - 查询权限列表 List Permissions
  - 删除权限 Delete Permission

### 🌐 命名空间管理 Namespace Management
- [x] 查询命名空间 List Namespaces
- [x] 创建命名空间 Create Namespace
- [x] 修改命名空间 Update Namespace
- [x] 删除命名空间 Delete Namespace

## 安装 Installation

nacos-client 支持最新的两个 Go 版本，需要 Go 1.20 或更高版本，并支持 Go modules。
nacos-client supports the 2 latest Go versions and requires Go 1.20+ with modules support.

```shell
go mod init github.com/my/repo
go get github.com/cx0110/nacos-config-go
```

## 快速开始 Quick Start

### 基础用法 Basic Usage

```go
package main

import (
	"fmt"
	"log"
	nacos "github.com/cx0110/nacos-config-go"
)

func main() {
	// 创建客户端 Create client
	client := nacos.New(&nacos.Config{
		Addr:     "http://localhost:8848",
		Username: "nacos",
		Password: "nacos",
	})

	// 检查服务健康状态 Check service health
	if err := client.Health(); err != nil {
		log.Fatal("Health check failed:", err)
	}

	// 登录获取访问令牌 Login to get access token
	if err := client.Login(); err != nil {
		log.Fatal("Login failed:", err)
	}

	// 发布配置 Publish configuration
	err := client.PublishConfig(&nacos.PublishConfigRequest{
		ConfigBase: nacos.ConfigBase{
			DataId: "app.properties",
			Group:  "DEFAULT_GROUP",
			Tenant: "public",
		},
		Content:     "app.name=demo\napp.version=1.0.0",
		ContentType: "properties",
	})
	if err != nil {
		log.Fatal("Publish config failed:", err)
	}

	// 获取配置 Get configuration
	content, err := client.GetConfig(&nacos.ConfigBase{
		DataId: "app.properties",
		Group:  "DEFAULT_GROUP",
		Tenant: "public",
	})
	if err != nil {
		log.Fatal("Get config failed:", err)
	}

	fmt.Println("Configuration content:", content)
}
```

### 用户管理示例 User Management Example

```go
// 创建用户 Create user
err := client.CreateUser(&nacos.User{
	Username: "testuser",
	Password: "password123",
})

// 查询用户列表 List users
	users, err := client.GetUsers(&nacos.Page{
	PageNo:   1,
	PageSize: 50,
})

// 删除用户 Delete user
err := client.DeleteUser(&nacos.DeleteUserRequest{
	Username: "testuser",
})
```

### 权限管理示例 Permission Management Example

```go
// 创建权限 Create permission
err := client.CreatePermission(&nacos.CreatePermissionRequest{
	Role:        "ADMIN",
	NamespaceId: "test-namespace",
	Action:      "rw", // r=read, w=write, rw=read/write
})

// 查询权限列表 List permissions
permissions, err := client.GetPermissions(&nacos.Page{
	PageNo:   1,
	PageSize: 50,
})
```

### 高级用法 Advanced Usage

### 配置监听 Configuration Listener

```go
// 注意：此功能正在开发中
// Note: This feature is under development
err := client.ListenConfig(&nacos.ListeningConfigs{
	ConfigBase: nacos.ConfigBase{
		DataId: "app.properties",
		Group:  "DEFAULT_GROUP",
		Tenant: "public",
	},
	ContentMD5: "5d41402abc4b2a76b9719d911017c592",
})
```

## API 参考 API Reference

### 配置管理 APIs

| 方法 Method | 描述 Description |
|-----------|----------------|
| `GetConfig()` | 获取配置 Get configuration |
| `PublishConfig()` | 发布配置 Publish configuration |
| `DeleteConfig()` | 删除配置 Delete configuration |
| `ListenConfig()` | 监听配置变化 Listen to configuration changes |
| `GetConfigHistory()` | 获取配置历史 Get configuration history |
| `GetConfigHistoryDetail()` | 获取历史详情 Get history details |
| `GetConfigHistoryPrevious()` | 获取上一版本 Get previous version |

### 用户管理 APIs

| 方法 Method | 描述 Description |
|-----------|----------------|
| `CreateUser()` | 创建用户 Create user |
| `GetUsers()` | 查询用户列表 List users |
| `PutUser()` | 修改用户 Update user |
| `DeleteUser()` | 删除用户 Delete user |

### 角色管理 APIs

| 方法 Method | 描述 Description |
|-----------|----------------|
| `CreateRoles()` | 创建角色 Create role |
| `GetRoles()` | 查询角色列表 List roles |
| `DeleteRoles()` | 删除角色 Delete role |

### 权限管理 APIs

| 方法 Method | 描述 Description |
|-----------|----------------|
| `CreatePermission()` | 创建权限 Create permission |
| `GetPermissions()` | 查询权限列表 List permissions |
| `DeletePermission()` | 删除权限 Delete permission |

## 测试 Testing

```shell
# 运行所有测试 Run all tests
go test ./...

# 运行测试并显示覆盖率 Run tests with coverage
go test -v -coverprofile=coverage.out ./...

# 生成覆盖率报告 Generate coverage report
go tool cover -html=coverage.out -o coverage.html
```

## 贡献 Contributing

我们欢迎所有形式的贡献！我们欢迎任何形式的贡献，包括但不限于：
We welcome all forms of contribution! This includes but is not limited to:

- 🐛 报告 Bug (Bug Reports)
- 💡 新功能建议 (Feature Requests)
- 📝 文档改进 (Documentation Improvements)
- 🔧 代码贡献 (Code Contributions)
- 🧪 测试用例 (Test Cases)

请查看 [CONTRIBUTING.md](CONTRIBUTING.md) 了解详细的贡献指南。
Please check [CONTRIBUTING.md](CONTRIBUTING.md) for detailed contribution guidelines.

## 许可证 License

本项目采用 Apache 2.0 许可证。详情请查看 [LICENSE](LICENSE) 文件。
This project is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for details.

## 联系我们 Contact Us

- 📧 Email: githun.cx0110@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/cx0110/nacos-config-go/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/cx0110/nacos-config-go/discussions)

## 致谢 Acknowledgments

- [Nacos](https://nacos.io/) - 一个易于使用的动态服务发现、配置管理和服务管理平台
  An easy-to-use dynamic service discovery, configuration management and service management platform
- [Go](https://golang.org/) - Go 编程语言
  The Go Programming Language
- [Resty](https://github.com/go-resty/resty) - 简单的 HTTP 和 REST 客户端库
  Simple HTTP and REST client library for Go

