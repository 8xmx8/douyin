package conf

var Conf Config

func DefaultConfig() Config {
	return Config{
		Address: "0.0.0.0",
		Port:    23724, // 2023-07-24
		Database: confDatabase{
			Type:     "mysql",
			Host:     "localhost",
			Port:     3306,
			User:     "root",
			Password: "root",
			Name:     "douyin",
			DbFile:   "data/data.db",
		},
		Redis: confRedis{
			Host:     "127.0.0.1",
			Port:     6379,
			Password: "123456",
			Db:       3,
		},
		Log: confLog{
			Enable:     true,
			Level:      "info",
			Name:       "data/log/log.log",
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     28,
			Compress:   false,
		},
	}
}
