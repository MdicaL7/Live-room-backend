package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64     `json:"id"`
	RoomID    int64     `json:"room_id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// 构造函数
func NewComment(roomID, userID int64, content string) *Comment {
	return &Comment{
		RoomID:    roomID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (c *Comment) CreateComment(db *gorm.DB, comment *Comment) error {
	return db.Create(comment).Error
}

func (c *Comment) GetCommentByID(db *gorm.DB, id int64) (*Comment, error) {
	var comment Comment
	if err := db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *Comment) ListCommentsByRoomID(db *gorm.DB, roomID int64) ([]Comment, error) {
	var comments []Comment
	if err := db.Where("room_id = ?", roomID).Order("created_at asc").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *Comment) DeleteComment(db *gorm.DB, id int64) error {
	return db.Delete(&Comment{}, id).Error
}
