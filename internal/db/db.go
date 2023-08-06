package db

import (
	"context"
	"github.com/Godvictory/douyin/cmd/flags"
	"github.com/Godvictory/douyin/internal/model"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"reflect"
	"time"
)

var db *gorm.DB
var rdb *redis.Client

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

func InitRdb(r *redis.Client) {
	rdb = r
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接redis出错，错误信息：%v", err)
	}
	// 内存模式下清空 Redis
	if flags.Memory {
		rdb.FlushAll(ctx)
	}
}

func id(val int64) model.Model {
	return model.Model{ID: val}
}

func GetDb() *gorm.DB {
	return db
}

func GetRdb() *redis.Client {
	return rdb
}
