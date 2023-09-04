package server

import (
	"bytes"
	"fmt"
	"github.com/Godvictory/douyin/cmd/flags"
	"github.com/Godvictory/douyin/internal/bootstrap"
	"github.com/Godvictory/douyin/internal/model"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	flags.DataDir = "../data"
	bootstrap.InitConf()
	bootstrap.InitDb()
	bootstrap.InitRdb()
	flags.Tst = true
	gin.SetMode(gin.TestMode)
	router = gin.New()
	Init(router)
}

func TestServer(t *testing.T) {
	t.Run("游客视频流获取", func(t *testing.T) {
		v := videoGet(t, nil)
		assert.NotNil(t, v["video_list"], v)
	})
	var token1, token2 string
	var uid1, uid2 string
	var video1 map[string]any
	t.Run("注册,登录,用户信息", func(t *testing.T) {
		user := &Data{Json: H{
			"username": "test111",
			"password": "230724",
		}}
		v1 := userRegister(t, user)
		assert.IsType(t, "", v1["user_id"])
		v2 := userLogin(t, user)
		assert.IsType(t, "", v2["token"])
		assert.IsType(t, "", v2["user_id"])
		token1 = v2["token"].(string)
		uid1 = v2["user_id"].(string)
		v3 := userInfo(t, &Data{Query: S{"user_id": uid1, "token": token1}})
		assert.NotNil(t, v3["user"], uid1, token1)
	})
	t.Run("登录视频流获取", func(t *testing.T) {
		v := videoGet(t, &Data{Query: S{"token": token1}})
		assert.NotNil(t, []model.Video{}, v["video_list"])
		video1 = v["video_list"].([]any)[0].(map[string]any)
	})
	t.Run("点赞,评论,关注", func(t *testing.T) {
		favoriteAction(t, &Data{Json: H{
			"token":       token1,
			"video_id":    video1["id"],
			"action_type": 1,
		}})
		commentAction(t, &Data{Json: H{
			"token":        token1,
			"video_id":     video1["id"],
			"action_type":  1,
			"comment_text": "testtest111",
		}})
		relationAction(t, &Data{Json: H{
			"token":       token1,
			"to_user_id":  video1["author"].(map[string]any)["id"],
			"action_type": 1,
		}})
	})
	t.Run("喜欢/评论/关注列表", func(t *testing.T) {
		v := favoriteList(t, &Data{Query: S{"user_id": uid1, "token": token1}})
		assert.NotNil(t, v["video_list"])
		v = commentList(t, &Data{Query: S{"video_id": uid1, "token": token1}})
		assert.NotNil(t, v["comment_list"])
		v = relationFollowList(t, &Data{Query: S{"user_id": uid1, "token": token1}})
		assert.NotNil(t, v["user_list"])
	})
	t.Run("投稿,发布列表", func(t *testing.T) {
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)
		filePath := "../public/test1.mp4"
		file, err := os.Open(filePath)
		if err != nil {
			t.Error(err)
		}

		part, _ := bodyWriter.CreateFormFile("data", filePath)
		_ = bodyWriter.WriteField("token", token1)
		_ = bodyWriter.WriteField("title", "video_test111")
		var fileData []byte
		file.Read(fileData)
		part.Write(fileData)
		file.Close()
		bodyWriter.Close()
		videoAction(t, &Data{Form: &F{
			r: bodyBuf,
			w: bodyWriter,
		}})
		v := videoList(t, &Data{Query: S{"user_id": uid1, "token": token1}})
		assert.NotNil(t, v["video_list"])
	})
	t.Run("注册2,关注,粉丝/好友列表", func(t *testing.T) {
		user := &Data{Json: H{
			"username": "test222",
			"password": "230724",
		}}
		v := userRegister(t, user)
		assert.IsType(t, "", v["user_id"])
		assert.IsType(t, "", v["token"])
		token2 = v["token"].(string)
		uid2 = v["user_id"].(string)
		relationAction(t, &Data{Json: H{
			"token":       token2,
			"to_user_id":  uid1,
			"action_type": 1,
		}})
		relationAction(t, &Data{Json: H{
			"token":       token1,
			"to_user_id":  uid2,
			"action_type": 1,
		}})
		relationFollowerList(t, &Data{Json: H{"user_id": uid1, "token": token1}})
		relationFriendList(t, &Data{Json: H{"user_id": uid1, "token": token1}})
	})
	t.Run("轮询问候", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				// 1 To 2
				messageAction(t, &Data{Json: H{
					"token":       token1,
					"to_user_id":  uid2,
					"action_type": 1,
					"content":     fmt.Sprint("message Test 1To2", i),
				}})
			} else {
				// 2 To 1
				messageAction(t, &Data{Json: H{
					"token":       token2,
					"to_user_id":  uid1,
					"action_type": 1,
					"content":     fmt.Sprint("message Test 2To1", i),
				}})
			}
		}
	})
	t.Run("聊天记录", func(t *testing.T) {
		messageChat(t, &Data{Query: S{
			"token":        token2,
			"to_user_id":   uid1,
			"pre_msg_time": "0",
		}})
		messageChat(t, &Data{Query: S{
			"token":        token1,
			"to_user_id":   uid2,
			"pre_msg_time": "0",
		}})
	})
}
