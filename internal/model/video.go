package model

import (
	"context"
	"github.com/Godvictory/douyin/utils"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type (
	// Video 视频表
	Video struct {
		Model
		AuthorID      int64      `json:"-" gorm:"index;notNull;comment:视频作者信息"`
		Author        User       `json:"author"`
		PlayUrl       string     `json:"play_url" gorm:"comment:视频播放地址"`
		CoverUrl      string     `json:"cover_url" gorm:"comment:视频封面地址"`
		FavoriteCount int64      `json:"favorite_count" gorm:"-"`          // 视频的点赞总数
		FavoriteUser  []*User    `json:"-" gorm:"many2many:UserFavorite:"` // 喜欢用户列表
		CommentCount  int64      `json:"comment_count" gorm:"-"`           // 视频的评论总数
		PlayCount     int64      `json:"play_count" gorm:"comment:视频的播放量"`
		IsFavorite    bool       `json:"is_favorite" gorm:"-"` // 是否点赞
		Title         string     `json:"title" gorm:"comment:视频标题"`
		Desc          string     `json:"desc" gorm:"comment:简介"`
		Comment       []*Comment `json:"comment,omitempty"` // 评论列表
		// 自建字段
		CoAuthor []*User `json:"authors,omitempty" gorm:"many2many:UserCreation;"` // 联合投稿
	}
)

var (
	videoFavoriteCountKey = make([]byte, 0, 50)
	videoCommentCountKey  = make([]byte, 0, 50)
	videoPlayCountKey     = make([]byte, 0, 50)
)

func (v *Video) AfterFind(tx *gorm.DB) (err error) {
	if uidA, ok := tx.Get("user_id"); ok {
		if uid, ok := uidA.(int64); ok {
			v.IsFavorite = getVideoIsFavorite(tx, uid, v.ID)
		}
	}
	if ipA, ok := tx.Get("ip"); ok {
		if ip, ok := ipA.(string); ok {
			v.PlayCount = getVideoPlayCount(ip, v.ID)
		}
	}
	v.FavoriteCount = getVideoFavoriteCount(tx, v.ID)
	v.CommentCount = getVideoCommentCount(tx, v.ID)
	return
}

func (v *Video) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == 0 {
		v.ID = utils.GetId(3, 20230724)
	}
	return
}

// getVideoFavoriteCount 视频的点赞总数
func getVideoFavoriteCount(tx *gorm.DB, vid int64) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getKey(vid, videoFavoriteCountKey)
	favoriteCount, err := rdb.Get(ctx, key).Int64()
	if err == redis.Nil {
		tx.Table("user_favorite").Where("video_id = ?", vid).Count(&favoriteCount)
		_ = rdb.Set(ctx, key, favoriteCount, 3*time.Second)
	}
	return favoriteCount
}

// getVideoCommentCount 视频的评论总数
func getVideoCommentCount(tx *gorm.DB, vid int64) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getKey(vid, videoCommentCountKey)
	CommentCount, err := rdb.Get(ctx, key).Int64()
	if err == redis.Nil {
		tx.Model(&Comment{}).Where("video_id = ?", vid).Count(&CommentCount)
		_ = rdb.Set(ctx, key, CommentCount, 3*time.Second)
	}
	return CommentCount
}

// getVideoPlayCount 视频的播放量
func getVideoPlayCount(ip string, vid int64) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := getKey(vid, videoPlayCountKey)
	rdb.PFAdd(ctx, key, ip)
	val, _ := rdb.PFCount(ctx, key).Result()
	return val
}

// getVideoIsFavorite 视频是否点赞
func getVideoIsFavorite(tx *gorm.DB, uid, vid int64) bool {
	result := map[string]any{}
	return tx.Table("user_favorite").Where("user_id = ? AND video_id = ?", uid, vid).Take(&result).RowsAffected == 1
	// data[i].IsFavorite = db.Raw("SELECT * FROM user_favorite WHERE user_id = ? AND video_id = ?", uid, data[i].ID).Scan(&result).RowsAffected == 1
}

func init() {
	addMigrate(&Video{})
	videoFavoriteCountKey = append(videoFavoriteCountKey, "video:favorite_count/"...)
	videoCommentCountKey = append(videoCommentCountKey, "video:comment_count/"...)
	videoPlayCountKey = append(videoPlayCountKey, "video:play_count/"...)
}
