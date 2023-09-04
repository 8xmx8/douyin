package server

import (
	"github.com/Godvictory/douyin/cmd/flags"
	"github.com/Godvictory/douyin/server/handlers"
	"github.com/Godvictory/douyin/server/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

func Init(r *gin.Engine) {
	r.MaxMultipartMemory = 16 << 20 // 16 MiB
	if !flags.Tst {
		r.Use(middleware.Logger(log.StandardLogger())) // 使用logrus记录日志
		r.Use(gin.Recovery())                          // 恐慌恢复
		r.Use(cors.Default())                          // 跨域处理
	}
	r.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router := r.Group("douyin")
	tester := r.Group("douyin", middleware.Test())
	// 视频类接口
	{
		newRouter(router, "GET", "feed", handlers.VideoGet)                      // 获取视频流
		newRouter(router, "POST", "publish/action/", handlers.VideoAction)       // 视频投稿
		newRouter(tester, "POST", "publish/actionUrl/", handlers.VideoActionUrl) // 视频投稿(测试接口)
		newRouter(router, "GET", "publish/list/", handlers.VideoList)            // 获取发布列表
		newRouter(router, "GET", "publish/follow/", handlers.VideoFollowList)    // 获取关注视频列表
	}
	// 用户类接口
	{
		newRouter(router, "POST", "user/register/", handlers.UserRegister) // 用户注册
		newRouter(router, "POST", "user/login/", handlers.UserLogin)       // 用户登录
		newRouter(router, "GET", "user/", handlers.UserInfo)               // 获取用户信息
	}
	// 互动类接口,
	{
		newRouter(router, "POST", "favorite/action/", handlers.FavoriteAction) // 点赞操作
		newRouter(router, "GET", "favorite/list/", handlers.FavoriteList)      // 获取喜欢列表
		newRouter(router, "POST", "comment/action/", handlers.CommentAction)   // 评论操作
		newRouter(router, "GET", "comment/list/", handlers.CommentList)        // 获取评论列表
	}
	// 社交类接口
	{
		newRouter(router, "POST", "relation/action/", handlers.RelationAction)             // 关注/取关 操作
		newRouter(router, "GET", "relation/follow/list/", handlers.RelationFollowList)     // 获取用户关注列表
		newRouter(router, "GET", "relation/follower/list/", handlers.RelationFollowerList) // 获取用户粉丝列表
		newRouter(router, "GET", "relation/friend/list/", handlers.RelationFriendList)     // 获取用户好友列表
		// 消息类接口
		{
			newRouter(router, "GET", "message/chat/", handlers.MessageChat)      // 获取消息
			newRouter(router, "POST", "message/action/", handlers.MessageAction) // 发送消息
		}
	}
	// 挂载 web 服务
	r.Use(static.Serve("/", static.LocalFile("web", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile("web/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write(content)
			c.Writer.Flush()
		}
	})
}
