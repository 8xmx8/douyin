# 公共

## model

> 不使用自增 id,改为使用雪花自增 19 位 id (雪花个毛线,就是纳秒时间戳简单乘加法运算<-🤡)

```go
type Model struct {
	ID        int64          `json:"id" gorm:"primarykey;comment:主键"`
	CreatedAt time.Time      `json:"-" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"-" gorm:"comment:修改时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"comment:删除时间"`
}

func (u *Model) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = utils.GetId()
	return
}
```

# User 类

## User

> 关注总数,粉丝总数 走 Redis,获赞,点赞走 db
>
> 粉丝,好友列表使用原生 SQL 走 UserFollow 连接表

```go
User struct {
	Model
	Name            string     `json:"name" gorm:"index:,unique;size:32;comment:用户名称"`
	Pawd            string     `json:"-" gorm:"size:128;comment:用户密码"`
	Avatar          string     `json:"avatar" gorm:"comment:用户头像"`
	BackgroundImage string     `json:"background_image" gorm:"comment:用户个人页顶部大图"`
	Signature       string     `json:"signature" gorm:"default:此人巨懒;comment:个人简介"`
	WorkCount       int64      `json:"work_count" gorm:"default:0;comment:作品数量"`
	Follow          []*User    `json:"follow,omitempty" gorm:"many2many:UserFollow;comment:关注列表"`
	Favorite        []*Video   `json:"like_list,omitempty" gorm:"many2many:UserFavorite;comment:喜欢列表"`
	Videos          []*Video   `json:"video_list,omitempty" gorm:"many2many:UserCreation;comment:作品列表"`
	Comment         []*Comment `json:"comment_list,omitempty" gorm:"comment:评论列表"`
	FollowCount     int64      `json:"follow_count" gorm:"-"`       // 关注总数
	FollowerCount   int64      `json:"follower_count" gorm:"-"`     // 粉丝总数
	TotalFavorited  int64      `json:"total_favorited" gorm:"-"`    // 获赞数量
	FavoriteCount   int64      `json:"favorite_count" gorm:"-"`     // 点赞数量
	IsFollow        bool       `json:"is_follow" gorm:"-"`          // 是否关注
	Follower        []*User    `json:"follower,omitempty" gorm:"-"` // 粉丝列表
	Friend          []*User    `json:"friend,omitempty" gorm:"-"`   // 好友列表
}
```

## FriendUser

> 这好像也不是模型,怎么在这?

```go
FriendUser struct {
	User
	Message string `json:"message"`         // 和该好友的最新聊天消息
	MsgType bool   `json:"msg_type,number"` // 0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}
```

## UserCreation

> 联合投稿~ 好像没时间写了....

```go
UserCreation struct {
		VideoID   int64  `json:"video_id,omitempty" gorm:"primaryKey"`
		UserID    int64  `json:"author_id" gorm:"primaryKey"`
		Type      string `json:"type" gorm:"comment:创作者类型"` //Up主,参演，剪辑，录像，道具，编剧，打酱油
		CreatedAt time.Time
		DeletedAt gorm.DeletedAt `json:"-"`
	}
```

# Video 类

## Video

> 点赞,点赞数量,评论数量,播放量都走 Redis
> 播放量是根据 ip 然后 HyperLogLog,并简单实现不重复推送
> 封面走阿里云 OSS 视频裁剪便宜好用,不像隔壁某山
> 联合投稿?不知道设计的时候咋写出来的捏

```go
Video struct {
	Model
	AuthorID      int64      `json:"-" gorm:"index;notNull;comment:视频作者信"`
	Author        User       `json:"author"`
	PlayUrl       string     `json:"play_url" gorm:"comment:视频播放地址"`
	CoverUrl      string     `json:"cover_url" gorm:"comment:视频封面地址"`
	Title         string     `json:"title" gorm:"comment:视频标题"`
	Desc          string     `json:"desc" gorm:"comment:简介"`
	Comment       []*Comment `json:"comment,omitempty" gorm:"comment:评论列表"`
	FavoriteUser  []*User    `json:"-" gorm:"many2many:UserFavorite;comment:欢用户列表"`
	IsFavorite    bool       `json:"is_favorite" gorm:"-"`    // 是否点赞
	PlayCount     int64      `json:"play_count" gorm:"-"`     // 视频播放量
	FavoriteCount int64      `json:"favorite_count" gorm:"-"` // 视频的点赞总数
	CommentCount  int64      `json:"comment_count" gorm:"-"`  // 视频的评论总数
	// 自建字段
	CoAuthor []*User `json:"authors,omitempty" gorm:"many2many:UserCreation;"` // 联合投稿
}
```

# Message 类

## Message

> **大坑!!!**
>
> api 文档 create_time 解释是`消息发送时间 yyyy-MM-dd HH:MM:ss`,项目文档也是 string 类型,最后才发现是毫秒级时间戳 🙃

```go
Message struct {
	ID         int64  `json:"id" gorm:"primarykey;comment:主键"`
	CreatedAt  int64  `json:"create_time" gorm:"autoUpdateTime:milli"`
	ToUserID   int64  `json:"to_user_id" gorm:"primaryKey;comment:该消息接收者的id"`
	FromUserID int64  `json:"from_user_id" gorm:"primaryKey;comment:该消息发送者的id"`
	ToUser     User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FromUser   User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content    string `json:"content" gorm:"comment:消息内容"`
}
```

# Comment 类

## Comment

> CreateDate 的生成使用`BeforeCreate`钩子实现

```go
Comment struct {
	Model
	UserID     int64  `json:"-" gorm:"index:idx_uvid;comment:评论用户信息"`
	VideoID    int64  `json:"-" gorm:"index:idx_uvid;comment:评论视频信息"`
	User       User   `json:"user" gorm:"constraint:OnUpdate:CASCADEOnDelete:CASCADE;"`
	Video      *Video `json:"video,omitempty"gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content    string `json:"content" gorm:"comment:评论内容"`
	CreateDate string `json:"create_date" gorm:"comment:评论发布日期"` // 格式mm-dd
	// 自建字段
	ReplyID int64 `json:"reply_id" gorm:"index;comment:回复ID"`
}

c.CreateDate = time.Now().Format("01-02")
```
