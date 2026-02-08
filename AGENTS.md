# 商品管理系统demo

## 项目概览（Project Overview）

### 简介
这是一个用于管理商品信息的demo，包含商品的增删改查等功能。
需要有前端页面，后端管理API。

### 技术栈
前端使用 React + Vite 框架，后端使用 Go 1.25，Web 框架使用 Gin，数据库暂时不用，所有东西都暂存在内存里。

### 功能
1. 需要对商品进行增删改查的管理。并且能看到商品列表信息。只需要最简单的实现即可。
2. 注意，由于没有数据库，每次启动时你都需要写入同样的静态的mock数据。

## 运行与环境（Environment & Commands）

后端项目运行在Kubernetes，因此你需要制备Dockerfile和Deployment、Service资源文件。

## 代码结构（Codebase Layout）

前端遵循标准的 React + Vite 技术栈结构即可。
后端遵循 golang-standards/project-layout 规范即可。

## 编码规范（Coding Guidelines）

前端和后端的所有注释必须使用中文。

## 测试与质量（Testing & QA）

后端的所有Web接口必须有单元测试。

## 文档

你需要将所有设计的API接口写到 ./backend/docs 中，便于后续编写文档和自动化测试。
