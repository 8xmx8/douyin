package db

import (
	"github.com/Godvictory/douyin/internal/model"
	"reflect"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	videoAll map[string][]int64
)

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
	{
		var data []map[string]any
		db.Model(&model.Video{}).Select("id", "`type_of`").Find(&data)
		videoAll = make(map[string][]int64, 10)
		videoAll["all"] = make([]int64, 0, len(data))
		for i := range data {
			id := data[i]["id"].(int64)
			videoAll["all"] = append(videoAll["all"], id)
			if data[i]["type_of"] != nil {
				ty := data[i]["type_of"].(string)
				videoAll[ty] = append(videoAll[ty], id)
			}
		}
	}
}

// id 快捷用法返回一个Model{id:val}
func id(val int64) model.Model {
	return model.Model{ID: val}
}

func GetDb() *gorm.DB {
	return db
}
