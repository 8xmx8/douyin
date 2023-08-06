package model

import (
	"github.com/Godvictory/douyin/utils"
	"gorm.io/gorm"
)

type (
	// User 用户信息表
	User struct {
		Model
		Name          string `json:"name" gorm:"index:,unique;size:32;comment:用户名称"`
		Pawd          string `json:"-" gorm:"size:128;comment:用户密码"`
		FollowCount   int64  `json:"follow_count" gorm:"default:0;comment:关注总数"`
		FollowerCount int64  `json:"follower_count" gorm:"default:0;comment:粉丝总数"`
		//IsFollow        bool       `json:"is_follow" gorm:"-"` // 是否关注
		Avatar          string     `json:"avatar" gorm:"comment:用户头像"`
		BackgroundImage string     `json:"background_image" gorm:"comment:用户个人页顶部大图"`
		Signature       string     `json:"signature" gorm:"default:此人巨懒;comment:个人简介"`
		WorkCount       int64      `json:"work_count" gorm:"default:0;comment:作品数量"`
		TotalFavorited  int64      `json:"total_favorited" gorm:"-"` // 获赞数量
		FavoriteCount   int64      `json:"favorite_count" gorm:"-"`  // 点赞数量
		Follow          []*User    `json:"follow,omitempty" gorm:"many2many:UserFollow;comment:关注列表"`
		Follower        []*User    `json:"follower,omitempty" gorm:"many2many:UserFollower;comment:粉丝列表"`
		Friend          []*User    `json:"friend,omitempty" gorm:"many2many:UserFriend;comment:好友列表"`
		Favorite        []*Video   `json:"like_list,omitempty" gorm:"many2many:UserFavorite;comment:喜欢列表"`
		Videos          []*Video   `json:"video_list,omitempty" gorm:"many2many:UserCreation;comment:作品列表"`
		Comment         []*Comment `json:"comment_list,omitempty" gorm:"comment:评论列表"`
	}
	// FriendUser 好友结构体
	FriendUser struct {
		User
		Message string `json:"message" gorm:"comment:和该好友的最新聊天消息"`
		MsgType bool   `json:"msg_type,number" gorm:"comment:消息类型"` // 0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.ID = utils.GetId(2, 114514)
	}
	return
}

func init() {
	addMigrate(&User{})
}
