package bootstrap

import (
	"encoding/json"
	"github.com/Godvictory/douyin/cmd/flags"
	"github.com/Godvictory/douyin/internal/conf"
	"github.com/Godvictory/douyin/utils"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var configPath string

func InitConf() int {
	configPath = filepath.Join(flags.DataDir, "config.json")
	if !utils.Exists(configPath) {
		// 配置文件不存在，创建默认配置
		log.Info("没检测到配置文件，将进行初始化 config.json.")
		basePath := filepath.Dir(configPath)
		err := os.MkdirAll(basePath, 0o766)
		if err != nil {
			log.Fatalf("无法创建文件夹, %s", err)
		}
		conf.Conf = conf.DefaultConfig()
		conf.Conf.JwtSecret = utils.RandString(17)
		defaultData, _ := json.MarshalIndent(conf.Conf, "", "  ")
		err = os.WriteFile(configPath, defaultData, 0o666)
		if err != nil {
			log.Fatalf("配置文件写入错误，请检查,{%s}", err)
		}
		return 1
	}
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("配置读取错误，请检查,{%s}", err)
	}
	data := os.ExpandEnv(string(file))
	err = json.Unmarshal([]byte(data), &conf.Conf)
	if err != nil {
		log.Fatalf("配置文件解析错误，请检查,{%s}", err)
	}
	// 解析完在回写一次,保证配置文件格式最新
	//fileData, _ := json.MarshalIndent(conf.Conf, "", "  ")
	//err = os.WriteFile(configPath, fileData, 0o666)
	//if err != nil {
	//	log.Error("配置文件更新错误，请检查,{%s}", err)
	//}
	return 0
}
