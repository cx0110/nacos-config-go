# Contributing to Nacos Config Management SDK for Go

感谢您对 Nacos Config Management SDK for Go 项目的关注！我们欢迎所有形式的贡献。

## 如何贡献

### 报告问题

- 使用 [GitHub Issues](https://github.com/cx0110/nacos-config-go/issues) 报告 bug
- 提供详细的重现步骤和环境信息
- 搜索现有 issues 以避免重复报告

### 提交代码

1. **Fork 项目**
   ```bash
   git clone https://github.com/your-username/nacos-config-go.git
   cd nacos-config-go
   ```

2. **创建分支**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **编写代码**
   - 遵循 Go 编码规范
   - 添加必要的测试
   - 更新相关文档

4. **运行测试**
   ```bash
   go test ./...
   go vet ./...
   go fmt ./...
   ```

5. **提交更改**
   ```bash
   git commit -m "feat: add your feature description"
   ```

6. **推送分支**
   ```bash
   git push origin feature/your-feature-name
   ```

7. **创建 Pull Request**

## 开发指南

### 代码风格

- 遵循 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- 使用 `gofmt` 格式化代码
- 添加必要的注释和文档

### 测试

```bash
# 运行所有测试
go test ./...

# 运行测试并生成覆盖率报告
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# 运行基准测试
go test -bench=. ./...
```

### 提交信息格式

使用 [Conventional Commits](https://www.conventionalcommits.org/) 格式：

- `feat:` 新功能
- `fix:` 修复 bug
- `docs:` 文档更新
- `style:` 代码格式化
- `refactor:` 代码重构
- `test:` 添加或修改测试
- `chore:` 其他不修改代码的更改

## 开发环境设置

### 前置要求

- Go 1.20 或更高版本
- Git
- Docker (用于运行 Nacos 测试环境)

### 设置开发环境

1. 安装依赖：
   ```bash
   go mod download
   ```

2. 启动 Nacos 测试环境：
   ```bash
   docker run -d --name nacos-standalone -e MODE=standalone -p 8848:8848 nacos/nacos-server:v2.2.3
   ```

3. 运行测试：
   ```bash
   go test ./...
   ```

## 发布流程

1. 更新版本号
2. 更新 CHANGELOG.md
3. 创建 Git tag
4. 发布 GitHub Release

## 许可证

通过贡献代码，您同意您的贡献将在 [Apache 2.0](LICENSE) 许可证下发布。

## 联系方式

如有疑问，请通过以下方式联系我们：

- [GitHub Discussions](https://github.com/cx0110/nacos-config-go/discussions)
- [Email](mailto:dev@nacos-config-go.com)

再次感谢您的贡献！🎉