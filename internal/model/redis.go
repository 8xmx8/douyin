package model

import (
	"context"
	"github.com/Godvictory/douyin/cmd/flags"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

// InitRdb 初始化 Redis
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

func GetRdb() *redis.Client {
	return rdb
}

// getKey 字符串快速拼接
func getKey(id int64, prefix []byte) string {
	s := make([]byte, 0, 50)
	copy(s, prefix)
	s = append(s, strconv.FormatInt(id, 10)...)
	return string(s)
}
