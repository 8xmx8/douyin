package db

import "github.com/Godvictory/douyin/internal/model"

func MessagePush(fid, tid int64, content string) error {
	data := model.Message{
		ToUserID:   tid,
		FromUserID: fid,
		Content:    content,
	}
	return db.Create(&data).Error
}

func MessageGet(fid, tid, preTime int64) ([]*model.Message, error) {
	var data []*model.Message
	err := db.Where(model.Message{ToUserID: tid, FromUserID: fid}).Order("created_at DESC").Find(&data).Error
	return data, err
}
