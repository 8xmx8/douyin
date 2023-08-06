package db

import (
	"errors"
	"github.com/Godvictory/douyin/internal/model"

	"gorm.io/gorm"
)

func RelationAction(fid, tid int64, ActionType int) error {
	var (
		associationA, associationB *gorm.Association
		errA, errB                 error
	)
	tx := db.Begin()
	associationA = tx.Model(&model.User{Model: id(fid)}).Association("Follow")
	associationB = tx.Model(&model.User{Model: id(tid)}).Association("Follower")
	switch ActionType {
	case 1:
		errA = associationA.Append(&model.User{Model: id(tid)})
		errB = associationB.Append(&model.User{Model: id(fid)})
	case 2:
		errA = associationA.Delete(&model.User{Model: id(tid)})
		errB = associationB.Delete(&model.User{Model: id(fid)})
	default:
		return errors.New("不合法的 ActionType")
	}
	if errA != nil || errB != nil {
		tx.Rollback()
		return errors.Join(errA, errB)
	}
	tx.Commit()
	return nil
}

func RelationFollowGet(uid int64) ([]*model.User, error) {
	var data []*model.User
	err := db.Set("user_id", uid).Model(&model.User{Model: id(uid)}).Association("Follow").Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func RelationFollowerGet(uid int64) ([]*model.User, error) {
	var data []*model.User
	err := db.Set("user_id", uid).Model(&model.User{Model: id(uid)}).Association("Follower").Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func RelationFriendGet(uid int64) ([]*model.User, error) {
	var data []*model.User
	err := db.Set("user_id", uid).Model(&model.User{Model: id(uid)}).Association("Friend").Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
