package model

import (
	"time"

	"gorm.io/gorm"
)

type LiveRoom struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	Cover       string    `json:"cover"`
	Anchor      string    `json:"anchor"`
	ReplayURL   string    `json:"replay_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 构造函数
func NewLiveRoom(title, description, replayURL, anchor, cover string, status int) *LiveRoom {
	now := time.Now()
	return &LiveRoom{
		Title:       title,
		Description: description,
		Status:      status,
		Cover:       cover,
		Anchor:      anchor,
		ReplayURL:   replayURL,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (lr *LiveRoom) CreateLiveRoom(db *gorm.DB, room *LiveRoom) error {
	return db.Create(room).Error
}

func (lr *LiveRoom) GetLiveRoomByID(db *gorm.DB, id int64) error {
	if err := db.First(&lr, id).Error; err != nil {
		return err
	}
	return nil
}

func (lr *LiveRoom) GetLiveRooms(db *gorm.DB) ([]*LiveRoom, error) {
	var live_rooms []*LiveRoom
	// if err := db.Find(&live_rooms).Error; err != nil {
	// 	return nil, err
	// }
	if err := db.Select("id,title,description,cover").Find(&live_rooms).Error; err != nil {
		return nil, err
	}
	return live_rooms, nil
}

func (lr *LiveRoom) UpdateLiveRoom(db *gorm.DB, id int64, title, description string) error {
	return db.Model(&LiveRoom{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":       title,
			"description": description,
			"updated_at":  time.Now(),
		}).Error
}

func (lr *LiveRoom) DeleteLiveRoom(db *gorm.DB, id int64) error {
	return db.Delete(&LiveRoom{}, id).Error
}
