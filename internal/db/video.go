package db

import (
	"errors"
	"fmt"
	"github.com/Godvictory/douyin/internal/model"
	"github.com/Godvictory/douyin/utils/upload"
	"mime/multipart"

	"gorm.io/gorm"
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

	return data, nil
}

// VideoUpload 视频投稿
func VideoUpload(uid int64, file multipart.File, PlayUrl, CoverUrl, title string, UserCreations []*model.UserCreation) (int64, string, error) {
	data := model.Video{
		AuthorID: uid,
		PlayUrl:  PlayUrl,
		CoverUrl: CoverUrl,
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
		fname := fmt.Sprintf("t/%d.mp4", data.ID)
		url, err := upload.Aliyun(fname, file)
		if err != nil {
			tx.Rollback()
			return 0, "上传出错...", err
		}
		data.PlayUrl = url + fname
		data.CoverUrl = url + fmt.Sprintf("t/%d.jpg", data.ID)
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
		tx.Model(&model.User{Model: id(uc.UserID)}).UpdateColumn("work_count", gorm.Expr("work_count + ?", 1))
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
