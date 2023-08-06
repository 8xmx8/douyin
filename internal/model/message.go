package model

import (
	"github.com/Godvictory/douyin/utils"
	"gorm.io/gorm"
)

type (
	// Message 消息表
	Message struct {
		Model
		ToUserID   int64  `json:"to_user_id" gorm:"primaryKey;comment:该消息接收者的id"`
		FromUserID int64  `json:"from_user_id" gorm:"primaryKey;comment:该消息发送者的id"`
		ToUser     User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		FromUser   User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Content    string `json:"content" gorm:"comment:消息内容"`
		//CreateTime string `json:"create_time" gorm:"comment:消息创建时间"`
	}
)

func (u *Message) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.ID = utils.GetId(3, 114514)
	}
	return
}

func init() {
	addMigrate(&Message{})
}
