package service

import (
	"live_room/internal/model"
	"live_room/pkg/db"
	"log"
)

type LiveRoom struct {
	liveRoom model.LiveRoom
}

// 查询直播间列表
func (lr *LiveRoom) GetLiveRooms() ([]*model.LiveRoom, error) {
	liveRooms, err := lr.liveRoom.GetLiveRooms(db.DB)
	if err != nil {
		log.Fatalf("liveRooms, err := lr.liveRoom.GetLiveRooms(db.DB) err:%v", err)
		return nil, err
	}
	return liveRooms, err
}

// 查询单个直播间详情
func (lr *LiveRoom) GetLRDetail(id int64) (*model.LiveRoom, error) {
	err := lr.liveRoom.GetLiveRoomByID(db.DB, id)
	if err != nil {
		return nil, err
	}
	return &lr.liveRoom, err
}
