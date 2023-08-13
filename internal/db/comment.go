package db

import (
	"github.com/Godvictory/douyin/internal/model"
)

// CommentPush 发送评论
func CommentPush(uid, vid int64, content string) (*model.Comment, error) {
	data := model.Comment{UserID: uid, VideoID: vid, Content: content}
	err := db.Create(&data).Error
	if err != nil {
		return nil, err
	}
	data.User.ID = uid
	db.Find(&data.User)
	video := model.Video{Model: id(vid)}
	video.HIncrByCommentCount(1)
	return &data, nil
}

// CommentDel 删除评论
func CommentDel(cid int64) error {
	var data model.Comment
	data.ID = cid
	err := db.Find(&data).Error
	if err != nil {
		return err
	}
	video := model.Video{Model: id(data.VideoID)}
	video.HIncrByCommentCount(-1)
	return db.Delete(&data).Error
}

// CommentGet 获取评论
func CommentGet(vid int64) ([]*model.Comment, error) {
	var data []*model.Comment
	err := db.Preload("User").Where("video_id = ?", vid).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
