package handlers

import (
	"github.com/Godvictory/douyin/internal/db"
	"github.com/Godvictory/douyin/internal/model"
	"github.com/Godvictory/douyin/utils/tokens"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	videoActionData struct {
		Data          multipart.File        `json:"data" form:"data"`           // 视频数据
		Token         string                `json:"token" form:"token"`         // 用户鉴权token
		Title         string                `json:"title" form:"title"`         // 视频标题
		TypeOf        string                `json:"type" form:"type"`           // 视频类型
		Url           string                `json:"url" form:"url"`             // 视频URL(测试环境)
		CoverUrl      string                `json:"cover_url" form:"cover_url"` // 视频封面URL(测试环境)
		UserID        int64                 `json:"id,string" form:"id"`        // 用户ID(测试环境)
		UserCreations []*model.UserCreation `json:"user_creations"`             // 联合投稿作者(半成品)
	}
)

// VideoGet 视频流获取
func VideoGet(c *gin.Context) (int, any) {
	var err error
	claims := new(tokens.MyClaims)
	token := c.Query("token")
	ty := c.Query("type")
	repeat, _ := strconv.ParseBool(c.Query("repeat"))
	if token != "" {
		claims, err = tokens.CheckToken(token)
		if err != nil {
			return Err("Token 错误,请重新登录", err)
		} // 没办法控制客户端退出登录,就这样好了(反正token 3个月才过期)
	}
	if ty == "" {
		ty = "all"
	}
	data, err := db.Feed(claims.ID, c.ClientIP(), ty, repeat)
	if err != nil {
		return Err("数据获取出错，请稍后再试.", err)
	}

	res := H{
		"video_list": data,
	}
	return Ok(res)
}

// VideoAction 视频投稿
func VideoAction(c *gin.Context) (int, any) {
	var data videoActionData
	file, _, err := c.Request.FormFile("data")
	data.Data = file
	data.Token = c.PostForm("token")
	data.Title = c.PostForm("title")
	data.TypeOf = c.PostForm("type")
	if err != nil || data.Token == "" {
		return ErrParam(err)
	}
	token, err := tokens.CheckToken(data.Token)
	if err != nil {
		return Err("Token 错误", err)
	}
	id, msg, err := db.VideoUpload(token.ID, data.Data, "", "", data.Title, data.TypeOf, data.UserCreations)
	if err != nil {
		return Err(msg, err)
	}
	return Ok(H{"vid": id})
}

// VideoActionUrl 视频投稿
// 测试接口可直接指定URL，或使用ID进行投稿
func VideoActionUrl(c *gin.Context) (int, any) {
	var data videoActionData

	err := c.ShouldBindJSON(&data)
	if err != nil || (data.UserID == 0 && data.Token == "") || (data.Data == nil && data.Url == "") {
		return ErrParam(err)
	}
	if data.Token != "" {
		token, err := tokens.CheckToken(data.Token)
		if err != nil {
			return Err("Token 错误", err)
		}
		data.UserID = token.ID
	}

	id, msg, err := db.VideoUpload(data.UserID, data.Data, data.Url, data.CoverUrl, data.Title, data.TypeOf, data.UserCreations)
	if err != nil {
		return Err(msg, err)
	}

	return Ok(H{"vid": id})
}

// VideoList 发布列表
func VideoList(c *gin.Context) (int, any) {
	var reqs userReqs
	// 参数绑定
	if err := c.ShouldBindQuery(&reqs); err != nil {
		return ErrParam(err)
	}
	if reqs.Token != "" {
		claims, err := tokens.CheckToken(reqs.Token)
		if err != nil {
			return Err("Token 错误", err)
		}
		if reqs.ID == 0 {
			reqs.ID = claims.ID
		}
	} else if reqs.ID == 0 {
		return Err("无参数!!!")
	}
	data, err := db.VideoList(reqs.ID)
	if err != nil {
		return Err("网卡了,再试一次吧", err)
	}

	return Ok(H{"video_list": data})
}

func VideoFollowList(c *gin.Context) (int, any) {
	var reqs userReqs
	// 参数绑定
	if err := c.ShouldBindQuery(&reqs); err != nil {
		return ErrParam(err)
	}
	claims, err := tokens.CheckToken(reqs.Token)
	if err != nil {
		return Err("Token 错误", err)
	}
	data, err := db.VideoFollowList(claims.ID)
	if err != nil {
		return Err("网卡了,再试一次吧", err)
	}
	return Ok(H{"video_list": data})
}
