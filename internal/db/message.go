package db

import "github.com/Godvictory/douyin/internal/model"

// MessagePush 发送消息
func MessagePush(fid, tid int64, content string) error {
	data := model.Message{
		ToUserID:   tid,
		FromUserID: fid,
		Content:    content,
	}
	return db.Create(&data).Error
}

// MessageGet 获取消息列表
func MessageGet(fid, tid, preTime int64) ([]*model.Message, error) {
	var data []*model.Message
	err := db.Where("created_at > ? AND (to_user_id = ? AND from_user_id = ? OR to_user_id = ? AND from_user_id = ?)", preTime, fid, tid, tid, fid).
		Order("created_at ASC").Find(&data).Error
	return data, err
}
