package handlers

import (
	"github.com/Godvictory/douyin/internal/db"
	"github.com/Godvictory/douyin/internal/model"
	"github.com/Godvictory/douyin/server/common"
	"github.com/Godvictory/douyin/utils/tokens"
	"github.com/gin-gonic/gin"
)

type actionReqs struct {
	Token      string `form:"token"   json:"token" binding:"required"`            // 用户鉴权token
	VideoId    int64  `form:"video_id"   json:"video_id" binding:"required"`      // 视频id
	ActionType int    `form:"action_type"  json:"action_type" binding:"required"` // 1-点赞，2-取消点赞
}

// FavoriteAction 点赞
func FavoriteAction(c *gin.Context) (int, any) {
	var reqs actionReqs
	// 参数绑定
	if err := common.Bind(c, &reqs); err != nil {
		return ErrParam(err)
	}
	claims, err := tokens.CheckToken(reqs.Token)
	if err != nil {
		return Err("Token 错误", err)
	}
	err = db.VideoLike(claims.ID, reqs.VideoId, reqs.ActionType)
	if err != nil {
		return Err("网卡了,再试一次吧", err)
	}
	return Ok(nil)
}

// FavoriteList 点赞列表
func FavoriteList(c *gin.Context) (int, any) {
	var (
		data []*model.Video
		reqs userReqs
	)
	// 参数绑定
	if err := c.ShouldBindQuery(&reqs); err != nil {
		return ErrParam(err)
	}
	claims, err := tokens.CheckToken(reqs.Token)
	if err != nil {
		return Err("Token 错误", err)
	}
	if reqs.ID == 0 {
		reqs.ID = claims.ID
	}
	data, err = db.VideoLikeList(reqs.ID)
	if err != nil {
		return Err("网卡了,再试一次吧", err)
	}
	return Ok(H{"video_list": data})
}
