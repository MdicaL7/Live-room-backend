package handler

import (
	"live_room/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LiveRoom struct {
}

// 获取直播间列表
func (lr *LiveRoom) GetLiveRooms(c *gin.Context) {
	liveRoomService := &service.LiveRoom{}
	liveRooms, err := liveRoomService.GetLiveRooms()
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "直播间列表为空",
			"data": nil,
		})
		return
	} else if err != nil {
		log.Printf("GetLiveRooms error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务端异常",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": liveRooms,
	})
}

// 获取单个直播间详情
func (lr *LiveRoom) GetLRDetail(c *gin.Context) {
	liveRoomId := c.Param("id")
	id, err := strconv.Atoi(liveRoomId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数id格式错误",
			"data": nil,
		})
		return
	}
	liveRoomService := &service.LiveRoom{}
	liveRoom, err := liveRoomService.GetLRDetail(int64(id))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "未找到直播间",
			"data": nil,
		})
		return
	} else if err != nil {
		log.Printf("GetLRDetail error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务端异常",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": liveRoom,
	})
}
