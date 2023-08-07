## server 文件夹介绍

**服务器相关内容**

路由，中间件等

### 文件夹结构
2023-07-26 14:58:27
```
server
├── common 			// 统一返回结构
│   └── common.go
├── handlers 			// 路由处理/接口
│   ├── comment.go
│   ├── favorite.go
│   ├── message.go
│   ├── publish.go
│   ├── relation.go
│   └── user.go
├── middleware 			// 中间件
│   └── logger.go
├── README.md
└── router.go			// 路由初始化
```

