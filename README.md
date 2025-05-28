# 仿小鹅通直播间项目

## 项目简介

本项目为仿小鹅通的“直播间+课程回放”系统，实现了从用户登录、直播间信息展示、课程回放视频播放、聊天室评论等功能，前后端分离开发，适合用于学习和功能拓展。

---

## 技术栈

- 前端：Vue 3 + Element Plus + Axios + Vue Router + Vite
- 后端：Go + Gin + GORM + MySQL
- 实时通信：WebSocket

---

## 已完成功能

### ✅ 前端（live-vue）

- 首页展示推荐直播课
- 登录弹窗（支持用户名密码登录）
- 登录状态展示及退出登录
- 直播详情页（包含播放器、讲解、聊天室）
- 弹幕评论聊天区（与后端 WebSocket 通信）
- 未登录状态下发送评论自动弹出登录框
- 页面结构组件化（LiveRoomCard / RecommendList / AnchorCard 等）

### ✅ 后端（backend）

- 用户登录接口 `/v1/api/login`
- 获取直播列表接口 `/v1/api/liveRoom`
- 获取直播详情接口 `/v1/api/liveRoom/:id`
- WebSocket 实时评论 `/v1/api/ws/:room_id`
- GORM 模型及数据库初始化
- 数据库模型设计（User、LiveRoom、Comment）

---

## 项目目录结构（简略）

```
project-root/
├── backend/                # Go 后端
│   ├── internal/
│   ├── model/
│   ├── handler/
│   └── main.go
├── live-vue/               # Vue3 前端
│   ├── src/
│   │   ├── pages/
│   │   ├── components/
│   │   ├── api/
│   │   └── App.vue
│   └── public/
│       ├── covers/
│       └── videos/
└── README.md
```

---

## 启动方法（简要）

### 后端

```bash
cd backend
go mod tidy
go run main.go
```

### 前端

```bash
cd live-vue
npm install
npm run dev
```

---

> 项目仍在开发中，后续将支持课程商品推荐、直播推流等高级功能。