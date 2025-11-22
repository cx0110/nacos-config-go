# 变更日志

本文档记录了 Nacos Config Management SDK for Go 项目的所有重要变更。

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
并且本项目遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

## [未发布]

### 新增
- 初始版本发布
- 完整的配置管理功能
- 权限管理系统
- 命名空间管理

## [1.0.0] - 2024-11-22

### 新增
- 🔧 **配置管理**
  - 获取配置 (GetConfig)
  - 发布配置 (PublishConfig)
  - 删除配置 (DeleteConfig)
  - 监听配置变化 (ListenConfig)
  - 查询配置历史 (GetConfigHistory)
  - 查询历史详情 (GetConfigHistoryDetail)
  - 查询上一版本 (GetConfigHistoryPrevious)

- 🛡️ **权限管理**
  - 用户管理
    - 创建用户 (CreateUser)
    - 查询用户列表 (GetUsers)
    - 修改用户密码 (PutUser)
    - 删除用户 (DeleteUser)
  - 角色管理
    - 创建角色 (CreateRoles)
    - 查询角色列表 (GetRoles)
    - 删除角色 (DeleteRoles)
  - 权限管理
    - 创建权限 (CreatePermission)
    - 查询权限列表 (GetPermissions)
    - 删除权限 (DeletePermission)

- 🌐 **命名空间管理**
  - 查询命名空间 (GetNamespaces)
  - 创建命名空间 (CreateNamespace)
  - 修改命名空间 (PutNamespace)
  - 删除命名空间 (DeleteNamespace)

### 技术特性
- 基于 Go 1.20+ 开发
- 使用 Resty 作为 HTTP 客户端
- 支持自动令牌刷新
- 完整的参数验证
- 错误处理和重试机制
- 完整的单元测试覆盖

### 文档
- 完整的中英文 README 文档
- API 参考文档
- 贡献指南
- 变更日志

### 许可证
- 采用 Apache 2.0 许可证

---

## 版本说明

### 版本号格式
本项目使用语义化版本号：`MAJOR.MINOR.PATCH`

- **MAJOR**：不兼容的 API 修改
- **MINOR**：向下兼容的功能性新增
- **PATCH**：向下兼容的问题修正

### 变更类型
- `新增` - 新功能
- `更改` - 对现有功能的变更
- `弃用` - 即将移除的功能
- `移除` - 已移除的功能
- `修复` - 问题修复
- `安全` - 安全相关的修复