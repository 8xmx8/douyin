package handlers

import (
	"errors"
	"fmt"
	"github.com/Godvictory/douyin/internal/db"
	"github.com/Godvictory/douyin/server/common"
	"github.com/Godvictory/douyin/utils/tokens"

	"github.com/gin-gonic/gin"
)

type relationReqs struct {
	Token      string `json:"token" form:"token" binding:"required"` // 用户鉴权token
	ToUserId   int64  `json:"to_user_id" form:"to_user_id"`          // 对方用户id
	ActionType int    `json:"action_type" form:"action_type"`        // 1-关注，2-取消关注
	UserId     int64  `json:"user_id" form:"user_id"`                // 用户id
}

// RelationAction 关系操作
func RelationAction(c *gin.Context) (int, any) {
	var (
		reqs relationReqs
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
	case 1, 2:
		err = db.RelationAction(claims.ID, reqs.ToUserId, reqs.ActionType)
	default:
		return ErrParam(errors.New("不合法的 ActionType"))
	}
	if err != nil {
		return Err("失败了.", err)
	}
	return Ok(H{})
}

// RelationFollowList 用户关注列表
func RelationFollowList(c *gin.Context) (int, any) {
	var (
		reqs relationReqs
		err  error
	)
	// 参数绑定
	if err := common.Bind(c, &reqs); err != nil {
		return ErrParam(err)
	}
	claims, err := tokens.CheckToken(reqs.Token)
	fmt.Println(reqs, claims, err)
	if err != nil {
		return Err("Token 错误", err)
	}
	if reqs.UserId == 0 {
		reqs.UserId = claims.ID
	}
	data, err := db.RelationFollowGet(reqs.UserId)
	if err != nil {
		return Err("再试试吧", err)
	}
	return Ok(H{"user_list": data})
}

// RelationFollowerList 用户粉丝列表
func RelationFollowerList(c *gin.Context) (int, any) {
	var (
		reqs relationReqs
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
	if reqs.UserId == 0 {
		reqs.UserId = claims.ID
	}
	data, err := db.RelationFollowerGet(reqs.UserId)
	if err != nil {
		return Err("再试试吧", err)
	}
	return Ok(H{"user_list": data})
}

// RelationFriendList 用户好友列表
func RelationFriendList(c *gin.Context) (int, any) {
	var (
		reqs relationReqs
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
	if reqs.UserId == 0 {
		reqs.UserId = claims.ID
	}
	data, err := db.RelationFriendGet(reqs.UserId)
	if err != nil {
		return Err("再试试吧", err)
	}
	return Ok(H{"user_list": data})
}
