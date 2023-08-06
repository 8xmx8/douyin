package conf

type Config struct {
	Address   string       `json:"address"`    // 监听地址
	Port      int          `json:"port"`       // 监听端口
	JwtSecret string       `json:"jwt_secret"` // Jwt密钥，随机生成
	Scheme    confScheme   `json:"scheme"`     // HTTPS配置
	Database  confDatabase `json:"database"`   // 数据库配置
	Redis     confRedis    `json:"redis"`      // Redis 缓存配置
	Oss       confOss      `json:"oss"`        // Oss 配置(阿里云)
	Log       confLog      `json:"log"`        // Log配置
}
type confScheme struct {
	Https    bool   `json:"https"`     // 启用HTTPS
	CertFile string `json:"cert_file"` // 证书路径
	KeyFile  string `json:"key_file"`  // 证书路径
}

type confDatabase struct {
	Type     string `json:"type"`     // 数据库类型，支持 sqlite3，mysql，postgresSql
	Host     string `json:"host"`     // 数据库地址
	Port     int    `json:"port"`     // 数据库端口
	User     string `json:"user"`     // 用户名
	Password string `json:"password"` // 密码
	Name     string `json:"name"`     // 数据库名
	DbFile   string `json:"db_file"`  // sqlite3的数据库文件，当为空时使用内存数据库
}
type confLog struct {
	Enable     bool   `json:"enable"`      // 是否启用日志
	Level      string `json:"level"`       // 日志等级，可用 panic,fatal,error,warn,info,debug,trace
	Name       string `json:"name"`        // 日志文件名
	MaxSize    int    `json:"max_size"`    // 日志最大大小
	MaxBackups int    `json:"max_backups"` // 日志最大备份数
	MaxAge     int    `json:"max_age"`     // 日志最长时间
	Compress   bool   `json:"compress"`    // 日志是否压缩
}

type confRedis struct {
	Host     string `json:"host"`     // 数据库地址
	Port     int    `json:"port"`     // 数据库端口
	Password string `json:"password"` // 密码
	Db       int    `json:"db"`       // 数据库编号
}

type confOss struct {
	// 阿里云配置
	AccessKeyID     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	Endpoint        string `json:"Endpoint"`
	BucketName      string `json:"BucketName"`
}
