package db

import (
	"github.com/Godvictory/douyin/internal/model"
	"reflect"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDb 初始化数据库服务
func InitDb(d *gorm.DB) {
	db = d
	for _, m := range model.GetMigrate() {
		err := db.AutoMigrate(m)
		if err != nil {
			log.Fatalf("%s 模型自动迁移失败: %s", reflect.TypeOf(m), err.Error())
		}
	}
	err := db.SetupJoinTable(&model.Video{}, "CoAuthor", &model.UserCreation{})
	if err != nil {
		log.Fatalf("自定义连接表设置失败,Video: %s", err)
	}
	err = db.SetupJoinTable(&model.User{}, "Videos", &model.UserCreation{})
	if err != nil {
		log.Fatalf("自定义连接表设置失败,User: %s", err)
	}
}

// id 快捷用法返回一个Model{id:val}
func id(val int64) model.Model {
	return model.Model{ID: val}
}

func GetDb() *gorm.DB {
	return db
}
