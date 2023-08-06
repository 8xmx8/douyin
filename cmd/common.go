package cmd

import (
	"github.com/Godvictory/douyin/cmd/flags"
	"github.com/Godvictory/douyin/utils"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
)

var pid = -1
var pidFile string

// initServer 初始化服务
func initServer() {
	formatter := logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2023-7-23 18:09:23",
		FullTimestamp:             true,
		DisableQuote:              true,
	}
	logrus.SetFormatter(&formatter)
	// 服务初始化
}

// initDaemon 守护进程初始化
func initDaemon() {
	pidFile = filepath.Join(flags.DataDir, "pid")
	if utils.Exists(pidFile) {
		bytes, err := os.ReadFile(pidFile)
		if err != nil {
			logrus.Fatal("无法读取pid文件，", err)
		}
		id, err := strconv.Atoi(string(bytes))
		if err != nil {
			logrus.Fatal("无法转换pid，", err)
		}
		pid = id
	}
}
