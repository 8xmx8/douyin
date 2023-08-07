package db

import (
	"errors"
	"fmt"
	"github.com/Godvictory/douyin/internal/model"
	"github.com/Godvictory/douyin/utils/upload"
	"mime/multipart"
	"sync"
)

// Feed 获取视频流
func Feed(uid int64, ip string, latestTime string) ([]model.Video, error) {
	var data []model.Video
	if len(latestTime) != 19 {
		latestTime = "9223372036854775806"
	}
	// t := time.Unix(0, latestTime*int64(time.Millisecond))
	err := db.Set("user_id", uid).Preload("Author").Where("id < ?", latestTime).Order("id DESC").Limit(5).Find(&data).Error
	if err != nil {
		return nil, err
	}

	for i := range data {
		var wg sync.WaitGroup
		if uid != 0 {
			wg.Add(1)
			go getIsFavorite(&wg, uid, data[i].ID, &data[i].IsFavorite) // 是否点赞
		}
		wg.Add(3)
		go getFavoriteCount(&wg, data[i].ID, &data[i].FavoriteCount) // 喜欢总数
		go getCommentCount(&wg, data[i].ID, &data[i].CommentCount)   // 评论总数
		go setPlayCount(&wg, ip, data[i].ID, &data[i].PlayCount)     // 播放量
		wg.Wait()
	}
	return data, nil
}

// VideoUpload 视频投稿
func VideoUpload(uid int64, file multipart.File, url, title string, UserCreations []*model.UserCreation) (int64, string, error) {
	data := model.Video{
		AuthorID: uid,
		PlayUrl:  url,
		Title:    title,
	}
	// 开启事务,上传失败不添加数据
	tx := db.Begin()
	err := tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}
	if file != nil {
		// reader := bytes.NewReader(file)
		url, err := upload.Aliyun(fmt.Sprintf("t/%d.mp4", data.ID), file)
		if err != nil {
			tx.Rollback()
			return 0, "上传出错...", err
		}
		data.PlayUrl = url
		err = tx.Save(&data).Error
		if err != nil {
			tx.Rollback()
			return 0, "更新出错...", err
		}
	}
	UserCreations = append([]*model.UserCreation{{VideoID: data.ID, UserID: uid, Type: "Up主"}}, UserCreations...)
	for _, uc := range UserCreations {
		uc.VideoID = data.ID
		err := tx.Create(uc).Error
		if err != nil {
			tx.Rollback()
			return 0, "创建出错...", err
		}
	}
	tx.Commit()
	return data.ID, "", nil
}

// VideoLike 视频点赞操作
func VideoLike(uid, vid int64, _type int) error {
	var err error
	association := db.Model(&model.User{Model: id(uid)}).Association("Favorite")
	switch _type {
	case 1:
		err = association.Append(&model.Video{Model: id(vid)})
	case 2:
		err = association.Delete(&model.Video{Model: id(vid)})
	default:
		err = errors.New("你看看你传的什么东西吧")
	}
	if err != nil {
		return err
	}
	return nil
}

// VideoLikeList 获取喜欢列表
func VideoLikeList(uid int64) ([]*model.Video, error) {
	var data []*model.Video
	err := db.Set("user_id", uid).Model(&model.User{Model: id(uid)}).Association("Favorite").Find(&data)
	if err != nil {
		return nil, err
	}
	for i := range data {
		err = db.Preload("Author").Find(data[i]).Error
		if err != nil {
			return nil, err
		}
		data[i].IsFavorite = true
	}
	return data, nil
}

// VideoList 获取作品列表
func VideoList(uid int64) ([]*model.Video, error) {
	var data []*model.Video
	err := db.Set("user_id", uid).Model(&model.User{Model: id(uid)}).Association("Videos").Find(&data)
	if err != nil {
		return nil, err
	}
	for i := range data {
		err = db.Preload("Author").Find(data[i]).Error
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}