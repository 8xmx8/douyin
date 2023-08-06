package bootstrap

import (
	"fmt"
	"github.com/Godvictory/douyin/cmd/flags"
	"github.com/Godvictory/douyin/internal/conf"
	"github.com/Godvictory/douyin/internal/db"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	stdlog "log"
	"strings"
	"time"
)

func InitDb() {
	var dialector gorm.Dialector
	var dB *gorm.DB
	var err error
	log.Info("开始初始化 Database 服务!")
	if flags.Memory {
		log.Info("当前处于 Memory模式,将使用内存数据库,并清空 Redis!")
		dialector = sqlite.Open("file::memory:?cache=shared")
	} else {
		database := conf.Conf.Database
		switch strings.ToUpper(database.Type) {
		case "SQLITE3":
			sqliteUrl := fmt.Sprintf("%s?_journal=WAL&_vacuum=incremental", database.DbFile)
			if database.DbFile == "" {
				sqliteUrl = "file::memory:?cache=shared"
			}
			dialector = sqlite.Open(sqliteUrl)
		case "MYSQL":
			dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				database.User, database.Password, database.Host, database.Port, database.Name))
		case "POSTGRES":
			dialector = postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
				database.Host, database.User, database.Password, database.Name, database.Port))
		default:
			log.Fatalf("not supported database type: %s,supported:[sqlite3,mysql,postgres]", database.Type)
		}
	}
	logLevel := logger.Silent
	if flags.Debug || flags.Dev {
		logLevel = logger.Info
	}
	dB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(
			stdlog.New(log.StandardLogger().Out, "\r\n", stdlog.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logLevel,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		TranslateError: true,
	})
	if err != nil {
		log.Fatalf("无法连接到数据库:%s", err.Error())
	}
	db.InitDb(dB)
	log.Info("初始化 Database 成功!", dB)
}

func InitRdb() {
	log.Info("开始初始化 Redis 服务!")
	rconf := conf.Conf.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rconf.Host, rconf.Port),
		Password: rconf.Password,
		DB:       rconf.Db,
	})
	db.InitRdb(rdb)
	log.Info("初始化 Redis 成功! ", rdb)
}
