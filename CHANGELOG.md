# Changelog

本项目遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/)。

目前尚未建立版本 tag，变更记录在 Unreleased 下。

## [Unreleased]

### Added

- 支持解析 owl-admin 的 PHP 渲染器类，生成对应的 Go 渲染器和 `amis.go` 便捷函数。

### Changed

- `php-renderers/` 更新到 amis 6.13.0（230 个文件）。
- 生成代码的 import 路径改为 `github.com/wcz0/gamis/v6/renderers`。
- 依赖升级到 `github.com/wcz0/gamis/v6 v6.13.0`。
- 清理生成目录后重新输出 228 个渲染器文件，避免残留旧文件。

### Fixed

- 修复 `old-renderers/` 与 `dist/` 引用 6.13 新组件、但 `go.mod` 仍锁在 gamis v0.6.13 导致的编译失败。

## 历史

| 日期 | 变更 |
| --- | --- |
| 2025-09-26 | 将解析逻辑迁移为 PHP 渲染器转换，重构类型安全与错误处理 |
| 2024-09-30 | 补充 Go 版本实现 |
| 2024-08-15 | 增加 `Set` 方法，更新 amis 映射 |
| 2024-07-05 | 更新 amis 版本 |
| 2024-05-10 | 增加 `Component` 组件并优化代码 |
| 2024-03-23 | 首个提交 |
