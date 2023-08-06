package model

import (
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
		CoverUrl      string     `json:"cover_url" gorm:"-"`               // 视频封面地址
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
	// UserCreation 联合作者
	UserCreation struct {
		VideoID   int64  `json:"video_id,omitempty" gorm:"primaryKey"`
		UserID    int64  `json:"author_id" gorm:"primaryKey"`
		Type      string `json:"type" gorm:"comment:创作者类型"` //Up主,参演，剪辑，录像，道具，编剧，打酱油
		CreatedAt time.Time
		DeletedAt gorm.DeletedAt `json:"-"`
	}
)

func (u *Video) AfterFind(tx *gorm.DB) (err error) {
	//if u.CoverUrl == "" {
	//	u.CoverUrl = ""
	//}
	return
}

func (u *Video) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.ID = utils.GetId(3, 20230724)
	}
	return
}
func (uc *UserCreation) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func init() {
	addMigrate(&Video{})
}
