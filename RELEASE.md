# 自动发布流程说明

## 概述

本项目使用GitHub Actions来自动构建多个平台版本的应用并发布release。

## 支持的平台

- **Linux**: amd64, arm64
- **Windows**: amd64, arm64 (带.exe扩展名)
- **macOS**: amd64, arm64

## 使用方法

### 1. 推送代码到GitHub

确保您的代码已经推送到GitHub仓库。

### 2. 创建并推送标签

当您准备发布新版本时，创建一个新的标签：

```bash
# 创建本地标签
git tag v1.0.0

# 推送标签到GitHub
git push origin v1.0.0
```

### 3. 自动触发构建

推送标签后，GitHub Actions会自动：
1. 为所有支持的平台构建应用
2. 创建release并上传构建产物
3. 生成release notes

## 构建产物

每个release包含以下文件：
- `SmsForwarder_tgbot-linux-amd64` - Linux x86_64版本
- `SmsForwarder_tgbot-linux-arm64` - Linux ARM64版本
- `SmsForwarder_tgbot-windows-amd64.exe` - Windows x86_64版本
- `SmsForwarder_tgbot-windows-arm64.exe` - Windows ARM64版本
- `SmsForwarder_tgbot-darwin-amd64` - macOS x86_64版本
- `SmsForwarder_tgbot-darwin-arm64` - macOS ARM64版本

## 注意事项

- 标签必须以 `v` 开头（如：v1.0.0, v2.1.3）
- 构建过程大约需要5-10分钟
- 确保GitHub仓库有足够的存储空间
- 构建产物会自动添加到release中，无需手动操作

## 故障排除

如果构建失败，请检查：
1. 代码是否能正常编译
2. GitHub Actions的权限设置
3. 仓库的存储空间是否充足 