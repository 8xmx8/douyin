package handlers

import (
	"errors"
	"github.com/Godvictory/douyin/internal/db"
	"github.com/Godvictory/douyin/internal/model"
	"github.com/Godvictory/douyin/server/common"
	"github.com/Godvictory/douyin/utils/tokens"

	"github.com/gin-gonic/gin"
)

type commentReqs struct {
	Token       string `form:"token" json:"token" binding:"required"`       // 用户鉴权token
	VideoId     int64  `form:"video_id" json:"video_id" binding:"required"` // 视频id
	ActionType  int    `form:"action_type" json:"action_type"`              // 1-发布评论，2-删除评论
	CommentText string `form:"comment_text" json:"comment_text"`            // 用户填写的评论内容
	CommentId   int64  `form:"comment_id" json:"comment_id"`                // 要删除的评论id
}

// CommentAction 评论操作
func CommentAction(c *gin.Context) (int, any) {
	var (
		reqs commentReqs
		resp *model.Comment
		err  error
	)
	// 参数绑定
	if err := common.Bind(c, &reqs); err != nil {
		return ErrParam(err)
	}
	claims, err := tokens.CheckToken(reqs.Token)
	if err != nil {
		return Err("Token 错误", err)
	}
	switch reqs.ActionType {
	case 1:
		resp, err = db.CommentPush(claims.ID, reqs.VideoId, reqs.CommentText)
	case 2:
		err = db.CommentDel(reqs.CommentId)
	default:
		return ErrParam(errors.New("不合法的 ActionType"))
	}
	if err != nil {
		return Err("请再试一次吧", err)
	}
	return Ok(H{"comment": resp})
}

// CommentList 评论列表
func CommentList(c *gin.Context) (int, any) {
	var reqs commentReqs
	// 参数绑定
	if err := common.Bind(c, &reqs); err != nil {
		return ErrParam(err)
	}
	_, err := tokens.CheckToken(reqs.Token)
	if err != nil {
		return Err("Token 错误", err)
	}
	data, err := db.CommentGet(reqs.VideoId)
	if err != nil {
		return Err("稍后试试.", err)
	}
	return Ok(H{"comment_list": data})
}
