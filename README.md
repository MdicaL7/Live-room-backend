# 仿小鹅通直播间项目

## 项目简介

本项目为仿小鹅通的“直播间+课程回放”系统，实现了从用户登录、直播间信息展示、课程回放视频播放、聊天室评论等功能，前后端分离开发。

---

## 技术栈

- 后端：Go + Gin + GORM + MySQL
- 实时通信：WebSocket

---

## 已完成功能

- 用户登录接口 `/v1/api/login`
- 获取直播列表接口 `/v1/api/liveRoom`
- 获取直播详情接口 `/v1/api/liveRoom/:id`
- WebSocket 实时评论 `/v1/api/ws/:room_id`
- GORM 模型及数据库初始化
- 数据库模型设计（User、LiveRoom）
- 基于Gorilla WebSocket实现的实时聊天系统
- 支持并发连接的直播间房间管理
- 使用sync.Mutex保障连接池线程安全
- 数据库模型包含ID/标题/描述/封面/主播/回放地址等字段

---

## 前端功能测试路径
1. **查看回放**
进入主页 --> 点击查看回放 --> 跳转LiveRoom --> 点击播放（可测试暂停、倍速、后退）
2. **登录功能**
点击导航栏的登录按钮 --> 填写信息 --> 登录成功
测试帐号：
- bob 123123
- mike 123123
- alice 123123
- eva 123123
3. **聊天功能**
跳转LiveRoom --> 点击进入讨论区 --> 发送信息即可
可通过四个网页登录四个帐号，在推荐栏中选择两个不同的直播间，分别进入LiveRoom的讨论区来测试直播间隔离聊天

---

## 项目目录结构（完整）

```
project-root/
├── backend/                # Go 后端
│   ├── cmd/
│   │   └── main.go         # 程序入口
│   ├── config/             # 配置文件
│   │   ├── config.go       # 配置加载逻辑
│   │   └── config.yaml     # 配置文件
│   ├── internal/
│   │   ├── handler/        # HTTP路由处理
│   │   │   ├── liveRoom.go # 直播间相关路由
│   │   │   └── user.go     # 用户相关路由
│   │   ├── model/          # 数据库模型
│   │   │   ├── comment.go  # 评论模型
│   │   │   ├── liveRoom.go # 直播间模型
│   │   │   └── user.go     # 用户模型
│   │   ├── service/        # 业务逻辑
│   │   │   ├── liveroom.go # 直播间服务
│   │   │   └── user.go     # 用户服务
│   │   └── webSocket/     # 实时通信
│   │       └── webSocket.go# WebSocket连接管理
│   └── go.mod              # Go模块配置
├── live_room_db.sql        # 数据库初始化脚本
├── docker-compose.yml      # 容器化部署配置
└── README.md
```

---

## 启动方法（完整）
### 启动MySQL服务
```bash
# 初始化数据库（需提前安装MySQL）
mysql -u root -p < live_room_db.sql
# 或者使用docker容器化启动
docker-compose up -d
```
### 安装依赖并启动
```
go mod tidy
go run main.go
```

---
