package handlers

import (
	"errors"
	"github.com/Godvictory/douyin/internal/db"
	"github.com/Godvictory/douyin/server/common"
	"github.com/Godvictory/douyin/utils/tokens"

	"github.com/gin-gonic/gin"
)

type messageReqs struct {
	Token      string `json:"token" form:"token" binding:"required"`           // 用户鉴权token
	ToUserId   int64  `json:"to_user_id" form:"to_user_id" binding:"required"` // 对方用户id
	ActionType int32  `json:"action_type" form:"action_type"`                  // 1-发送消息
	Content    string `json:"content" form:"content"`                          // 消息内容
	PreMsgTime int64  `json:"pre_msg_time" form:"pre_msg_time"`                // 上次最新消息的时间
}

// MessageChat 聊天记录
func MessageChat(c *gin.Context) (int, any) {
	var reqs messageReqs
	// 参数绑定
	if err := common.Bind(c, &reqs); err != nil {
		return ErrParam(err)
	}
	claims, err := tokens.CheckToken(reqs.Token)
	if err != nil {
		return Err("Token 错误", err)
	}

	data, err := db.MessageGet(claims.ID, reqs.ToUserId, reqs.PreMsgTime)
	if err != nil {
		return Err("聊天记录获取失败", err)
	}
	return Ok(H{"message_list": data})
}

// MessageAction 消息操作
func MessageAction(c *gin.Context) (int, any) {
	var (
		reqs messageReqs
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
		err = db.MessagePush(claims.ID, reqs.ToUserId, reqs.Content)
	default:
		return ErrParam(errors.New("不合法的 ActionType"))
	}
	if err != nil {
		return Err("发送失败.", err)
	}
	return Ok(H{})
}
