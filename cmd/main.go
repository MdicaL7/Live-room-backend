package main

import (
	"live_room/config"
	"live_room/internal/handler"
	websocket "live_room/internal/webSocket"
	"live_room/pkg/db"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化读取配置文件
	config.InitConfig()
	dsn := config.Conf.GetString("mysql.dsn")
	//初始化拿到db连接
	db.InitMySQL(dsn)
	defer db.CloseMySQL()
	r := gin.Default()
	// 添加 CORS 支持
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Static("/videos", "./videos")
	liveRoom := &handler.LiveRoom{}
	//获取直播列表路由
	r.GET("/v1/api/liveRoom", liveRoom.GetLiveRooms)
	//获取单个直播
	r.GET("/v1/api/liveRoom/:id", liveRoom.GetLRDetail)
	user := &handler.User{}
	//登录
	r.POST("/v1/api/login", user.Login)
	//webSocket
	r.GET("/v1/api/ws/:roomID", websocket.WebsHandler)
	r.Run("0.0.0.0:8080")
}
