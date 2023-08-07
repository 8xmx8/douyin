package db

import (
	"github.com/Godvictory/douyin/internal/model"
	"gopkg.in/hlandau/passlib.v1"
)

// Register 注册
func Register(data *model.User) (msg string, err error) {
	hash, err := passlib.Hash(data.Pawd)
	if err != nil {
		msg = "请更换您的密码再试一次"
		return
	}
	data.Pawd = hash
	err = db.Create(&data).Error
	if err != nil {
		msg = "抱歉，请稍后再试..."
		return
	}
	return "", nil
}

// Login 登录
func Login(user, pawd string) (data *model.User, msg string, err error) {
	data = new(model.User)
	// 根据用户名获取对应的全部数据
	err = db.Where("name = ?", user).Find(&data).Error
	if err != nil {
		msg = "没有此用户名~"
		return
	}
	// 进行哈希值效验密码是否正确
	newHash, err := passlib.Verify(pawd, data.Pawd)
	if err != nil {
		msg = "用户名或者密码不正确!"
		return
	}
	if newHash != "" {
		// 登陆成功，判断是否需要更换哈希值
		db.Where(data).Update("pawd", newHash)
	}
	return
}

// UserInfo 获取用户信息
func UserInfo(id int64) (data model.User, msg string, err error) {
	data.ID = id
	err = db.Set("user_id", id).Find(&data).Error
	if err != nil {
		msg = "抱歉,请稍后再试"
	}
	return
}
