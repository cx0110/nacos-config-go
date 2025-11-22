# Nacos 配置管理 Go SDK

[![构建状态](https://github.com/cx0110/nacos-config-go/actions/workflows/badge.yml/badge.svg)](https://github.com/cx0110/nacos-config-go/actions)
[![Go 参考文档](https://pkg.go.dev/badge/github.com/cx0110/nacos-config-go)](https://pkg.go.dev/github.com/cx0110/nacos-config-go?tab=doc)
[![Go 报告卡](https://goreportcard.com/badge/github.com/cx0110/nacos-config-go)](https://goreportcard.com/report/github.com/cx0110/nacos-config-go)
[![许可证](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

> 一个功能完善的 Nacos 配置管理和权限控制 Go SDK，提供简洁易用的 API 来操作 Nacos 配置中心。

## 文档 Documentation

- [简体中文](./README_zh.md)
- [English](./README.md)

## 资源 Resources

- [Go 参考文档](https://pkg.go.dev/github.com/cx0110/nacos-config-go)
- [示例代码](https://pkg.go.dev/github.com/cx0110/nacos-config-go#pkg-examples)

## 功能特性

### 🔧 配置管理
- [x] 获取配置
- [x] 发布配置
- [x] 删除配置
- [x] 监听配置变化
- [x] 查询配置历史
- [x] 查询历史详情
- [x] 查询上一版本

### 🛡️ 权限管理
- [x] 用户管理
  - 创建用户
  - 查询用户列表
  - 修改用户密码
  - 删除用户
- [x] 角色管理
  - 创建角色
  - 查询角色列表
  - 删除角色
- [x] 权限管理
  - 创建权限
  - 查询权限列表
  - 删除权限

### 🌐 命名空间管理
- [x] 查询命名空间
- [x] 创建命名空间
- [x] 修改命名空间
- [x] 删除命名空间

## 安装

nacos-client 支持最新的两个 Go 版本，需要 Go 1.20 或更高版本，并支持 Go modules。

```shell
go mod init github.com/my/repo
go get github.com/cx0110/nacos-config-go
```

## 快速开始

### 基础用法

```go
package main

import (
	"fmt"
	"log"
	nacos "github.com/cx0110/nacos-config-go"
)

func main() {
	// 创建客户端
	client := nacos.New(&nacos.Config{
		Addr:     "http://localhost:8848",
		Username: "nacos",
		Password: "nacos",
	})

	// 检查服务健康状态
	if err := client.Health(); err != nil {
		log.Fatal("健康检查失败:", err)
	}

	// 登录获取访问令牌
	if err := client.Login(); err != nil {
		log.Fatal("登录失败:", err)
	}

	// 发布配置
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
		log.Fatal("发布配置失败:", err)
	}

	// 获取配置
	content, err := client.GetConfig(&nacos.ConfigBase{
		DataId: "app.properties",
		Group:  "DEFAULT_GROUP",
		Tenant: "public",
	})
	if err != nil {
		log.Fatal("获取配置失败:", err)
	}

	fmt.Println("配置内容:", content)
}
```

### 用户管理示例

```go
// 创建用户
err := client.CreateUser(&nacos.User{
	Username: "testuser",
	Password: "password123",
})

// 查询用户列表
users, err := client.GetUsers(&nacos.Page{
	PageNo:   1,
	PageSize: 50,
})

// 删除用户
err := client.DeleteUser(&nacos.DeleteUserRequest{
	Username: "testuser",
})
```

### 权限管理示例

```go
// 创建权限
err := client.CreatePermission(&nacos.CreatePermissionRequest{
	Role:        "ADMIN",
	NamespaceId: "test-namespace",
	Action:      "rw", // r=read, w=write, rw=read/write
})

// 查询权限列表
permissions, err := client.GetPermissions(&nacos.Page{
	PageNo:   1,
	PageSize: 50,
})
```

### 高级用法

### 配置监听

```go
// 注意：此功能正在开发中
err := client.ListenConfig(&nacos.ListeningConfigs{
	ConfigBase: nacos.ConfigBase{
		DataId: "app.properties",
		Group:  "DEFAULT_GROUP",
		Tenant: "public",
	},
	ContentMD5: "5d41402abc4b2a76b9719d911017c592",
})
```

## API 参考

### 配置管理 APIs

| 方法 | 描述 |
|------|------|
| `GetConfig()` | 获取配置 |
| `PublishConfig()` | 发布配置 |
| `DeleteConfig()` | 删除配置 |
| `ListenConfig()` | 监听配置变化 |
| `GetConfigHistory()` | 获取配置历史 |
| `GetConfigHistoryDetail()` | 获取历史详情 |
| `GetConfigHistoryPrevious()` | 获取上一版本 |

### 用户管理 APIs

| 方法 | 描述 |
|------|------|
| `CreateUser()` | 创建用户 |
| `GetUsers()` | 查询用户列表 |
| `PutUser()` | 修改用户 |
| `DeleteUser()` | 删除用户 |

### 角色管理 APIs

| 方法 | 描述 |
|------|------|
| `CreateRoles()` | 创建角色 |
| `GetRoles()` | 查询角色列表 |
| `DeleteRoles()` | 删除角色 |

### 权限管理 APIs

| 方法 | 描述 |
|------|------|
| `CreatePermission()` | 创建权限 |
| `GetPermissions()` | 查询权限列表 |
| `DeletePermission()` | 删除权限 |

## 测试

```shell
# 运行所有测试
go test ./...

# 运行测试并显示覆盖率
go test -v -coverprofile=coverage.out ./...

# 生成覆盖率报告
go tool cover -html=coverage.out -o coverage.html
```

## 贡献

我们欢迎所有形式的贡献！包括但不限于：

- 🐛 报告 Bug
- 💡 新功能建议
- 📝 文档改进
- 🔧 代码贡献
- 🧪 测试用例

请查看 [CONTRIBUTING.md](CONTRIBUTING.md) 了解详细的贡献指南。

## 许可证

本项目采用 Apache 2.0 许可证。详情请查看 [LICENSE](LICENSE) 文件。

## 联系我们

- 📧 邮箱: dev@nacos-config-go.com
- 🐛 问题反馈: [GitHub Issues](https://github.com/cx0110/nacos-config-go/issues)
- 💬 讨论区: [GitHub Discussions](https://github.com/cx0110/nacos-config-go/discussions)

## 致谢

- [Nacos](https://nacos.io/) - 一个易于使用的动态服务发现、配置管理和服务管理平台
- [Go](https://golang.org/) - Go 编程语言
- [Resty](https://github.com/go-resty/resty) - 简单的 HTTP 和 REST 客户端库

