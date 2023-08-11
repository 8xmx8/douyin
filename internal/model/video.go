package model

import (
	"context"
	"github.com/Godvictory/douyin/utils"
	"gorm.io/gorm"
	"time"
)

type (
	// Video 视频表
	Video struct {
		Model
		AuthorID      int64      `json:"-" gorm:"index;notNull;comment:视频作者信息"`
		Author        User       `json:"author"`
		PlayUrl       string     `json:"play_url" gorm:"comment:视频播放地址"`
		CoverUrl      string     `json:"cover_url" gorm:"comment:视频封面地址"`
		Title         string     `json:"title" gorm:"comment:视频标题"`
		Desc          string     `json:"desc" gorm:"comment:简介"`
		Comment       []*Comment `json:"comment,omitempty" gorm:"comment:评论列表"`
		FavoriteUser  []*User    `json:"-" gorm:"many2many:UserFavorite;comment:喜欢用户列表"`
		IsFavorite    bool       `json:"is_favorite" gorm:"-"`    // 是否点赞
		PlayCount     int64      `json:"play_count" gorm:"-"`     // 视频播放量
		FavoriteCount int64      `json:"favorite_count" gorm:"-"` // 视频的点赞总数
		CommentCount  int64      `json:"comment_count" gorm:"-"`  // 视频的评论总数
		// 自建字段
		CoAuthor []*User `json:"authors,omitempty" gorm:"many2many:UserCreation;"` // 联合投稿
	}
)

var (
	videoCountKey     = make([]byte, 0, 50) // video_count:
	videoPlayCountKey = make([]byte, 0, 50) // video_play_count:
)

func (v *Video) AfterFind(tx *gorm.DB) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getKey(v.ID, videoCountKey)
	playKey := getKey(v.ID, videoPlayCountKey)
	if uidA, ok := tx.Get("user_id"); ok {
		if uid, ok := uidA.(int64); ok {
			v.IsFavorite = getVideoIsFavorite(tx, uid, v.ID)
		}
	}
	// v.PlayCount, _ = rdb.HIncrBy(ctx, key, "play_count", 1).Result()
	v.PlayCount, _ = rdb.PFCount(ctx, playKey).Result()
	v.FavoriteCount, _ = rdb.HGet(ctx, key, "favorite_count").Int64()
	v.CommentCount, _ = rdb.HGet(ctx, key, "comment_count").Int64()
	tx.Find(&v.Author, v.AuthorID)
	return
}

func (v *Video) ViewedFilter(ip string) bool {
	playKey := getKey(v.ID, videoPlayCountKey)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	val, _ := rdb.PFAdd(ctx, playKey, ip).Result()
	// 1:未看 0:已看
	return val == 1
}

func (v *Video) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == 0 {
		v.ID = utils.GetId(3, 20230724)
	}
	return
}

func (v *Video) HIncrByFavoriteCount(incr int64) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	key := getKey(v.ID, videoCountKey)
	v.FavoriteCount, _ = rdb.HIncrBy(ctx, key, "favorite_count", incr).Result()
	return v.FavoriteCount
}

func (v *Video) HIncrByCommentCount(incr int64) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	key := getKey(v.ID, videoCountKey)
	v.FavoriteCount, _ = rdb.HIncrBy(ctx, key, "comment_count", incr).Result()
	return v.FavoriteCount
}

//// getVideoFavoriteCount 视频的点赞总数
//func getVideoFavoriteCount(tx *gorm.DB, vid int64) int64 {
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//	key := getKey(vid, videoFavoriteCountKey)
//	favoriteCount, err := rdb.Get(ctx, key).Int64()
//	if err == redis.Nil {
//		tx.Table("user_favorite").Where("video_id = ?", vid).Count(&favoriteCount)
//		_ = rdb.Set(ctx, key, favoriteCount, 3*time.Second)
//	}
//	return favoriteCount
//}

//// getVideoCommentCount 视频的评论总数
//func getVideoCommentCount(tx *gorm.DB, vid int64) int64 {
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//	key := getKey(vid, videoCommentCountKey)
//	CommentCount, err := rdb.Get(ctx, key).Int64()
//	if err == redis.Nil {
//		tx.Model(&Comment{}).Where("video_id = ?", vid).Count(&CommentCount)
//		_ = rdb.Set(ctx, key, CommentCount, 3*time.Second)
//	}
//	return CommentCount
//}

//// getVideoPlayCount 视频的播放量
//func getVideoPlayCount(ip string, vid int64) int64 {
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//	key := getKey(vid, videoPlayCountKey)
//	rdb.PFAdd(ctx, key, ip)
//	val, _ := rdb.PFCount(ctx, key).Result()
//	return val
//}

// getVideoIsFavorite 视频是否点赞
func getVideoIsFavorite(tx *gorm.DB, uid, vid int64) bool {
	result := map[string]any{}
	return tx.Table("user_favorite").Where("user_id = ? AND video_id = ?", uid, vid).Scan(&result).RowsAffected == 1
	// data[i].IsFavorite = db.Raw("SELECT * FROM user_favorite WHERE user_id = ? AND video_id = ?", uid, data[i].ID).Scan(&result).RowsAffected == 1
}

func init() {
	addMigrate(&Video{})
	videoCountKey = append(videoCountKey, "video_count:"...)
	videoPlayCountKey = append(videoPlayCountKey, "video_play_count:"...)
}
