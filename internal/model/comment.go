package model

import (
	"context"
	"github.com/Godvictory/douyin/utils"
	"time"

	"gorm.io/gorm"
)

type (
	// Comment 评论表
	Comment struct {
		Model
		UserID  int64  `json:"-" gorm:"index:idx_uvid;comment:评论用户信息"`
		VideoID int64  `json:"-" gorm:"index:idx_uvid;comment:评论视频信息"`
		User    User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Video   Video  `json:"video" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Content string `json:"content" gorm:"comment:评论内容"`
		// create_date string // 评论发布日期，格式 mm-dd
		// 自建字段
		ReplyID int64 `json:"reply_id" gorm:"index;comment:回复ID"`
	}
)

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == 0 {
		c.ID = utils.GetId(2, 20230724)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getKey(c.VideoID, videoCommentCountKey)
	rdb.Incr(ctx, key)
	return
}

func init() {
	addMigrate(&Comment{})
}
