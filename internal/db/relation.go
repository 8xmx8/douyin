package db

import (
	"errors"
	"github.com/Godvictory/douyin/internal/model"

	"gorm.io/gorm"
)

// RelationAction 关注/取关
func RelationAction(fid, tid int64, ActionType int) error {
	var (
		association *gorm.Association
		err         error
	)
	tx := db.Begin()
	fval := &model.User{Model: id(fid)}
	tval := &model.User{Model: id(tid)}
	association = tx.Model(fval).Association("Follow")
	switch ActionType {
	case 1:
		err = association.Append(tval)
		fval.HIncrByFollowCount(1)
		tval.HIncrByFollowerCount(1)
	case 2:
		err = association.Delete(tval)
		fval.HIncrByFollowCount(-1)
		tval.HIncrByFollowerCount(-1)
	default:
		return errors.New("不合法的 ActionType")
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// RelationFollowGet 获取关注列表 uid:本人id tid:待查id
func RelationFollowGet(uid, tid int64) ([]*model.User, error) {
	var data []*model.User
	err := db.Set("user_id", uid).Model(&model.User{Model: id(tid)}).Association("Follow").Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// RelationFollowerGet 获取粉丝列表 uid:本人id tid:待查id
func RelationFollowerGet(uid, tid int64) ([]*model.User, error) {
	var data []*model.User
	err := db.Set("user_id", uid).Table("user").
		Joins("JOIN user_follow ON `user`.`id` = `user_follow`.`user_id` AND `user_follow`.`follow_id` = ?", tid).
		Select("`user`.*").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// RelationFriendGet 获取好友列表 uid:本人id tid:待查id
func RelationFriendGet(uid int64) ([]*model.FriendUser, error) {
	var (
		data []*model.User
		res  []*model.FriendUser
	)
	err := db.Set("user_id", uid).
		Table("(SELECT `user`.* FROM `user` JOIN `user_follow` ON `user`.`id` = `user_follow`.`follow_id` AND `user_follow`.`user_id` = ?) as t", uid).
		Joins("JOIN `user_follow` ON `t`.`id` = `user_follow`.`user_id`").
		Where(" `user_follow`.`follow_id` = ?", uid).
		Select("`t`.*").Find(&data).Error
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		val := model.FriendUser{User: *d}
		tmsg := model.Message{ToUserID: d.ID, FromUserID: uid}
		fmsg := model.Message{ToUserID: uid, FromUserID: d.ID}

		db.Order("created_at DESC").Take(&tmsg)
		db.Order("created_at DESC").Take(&fmsg)

		if tmsg.CreatedAt > fmsg.CreatedAt {
			val.Message = tmsg.Content
			val.MsgType = 0
		} else if tmsg.CreatedAt < fmsg.CreatedAt {
			val.Message = fmsg.Content
			val.MsgType = 1
		} else {
			val.Message = "快来和你的新朋友聊天吧"
		}
		res = append(res, &val)
	}
	return res, nil
}
