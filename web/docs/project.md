<p align="center">
  <a href="https://github.com/Ocyss/Douyin">
    <img src="https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/douyin/logo.svg" alt="Logo" width="180" height="180">
  </a>

  <h1 align="center">极简版抖音</h1>
  <p align="center">
    一个字节青训营的实战项目
	<br />
	开始于2023.7.24 结束于2023.8.20
    <br />
     <br />
    <a href="https://github.com/Ocyss/Douyin/issues">报告Bug</a>
    <a href="https://github.com/Ocyss/Douyin/issues">提出新特性</a>
  </p>

## 前言

虽然是一个青训营的项目,开发的时间也不算长,但还是会尽力去写好一个项目,做好各个文档和注释,即使结束了,也能帮助后来者学习.

一开始想着收集些 17,18 年的热门视频来当资料,耗费了很多时间,确实不好收集,就直接 pxx 用了营销号的"无版权资源"~

项目结构借鉴 [Alist](https://github.com/alist-org/alist) 项目

## 技术栈

#### 后端 Golang 1.20

- Gin [(Web 框架)](https://gin-gonic.com/zh-cn/)
- GORM [(ORM)](https://gorm.io/zh_CN/)
- Cobra [(CLI 框架)](https://github.com/spf13/cobra)
- MySQL,SQLite,PostgreSQL [(数据库)]()
- Redis [(缓存)]()

#### 前端 Vue.js 3

- Vite [(构建工具)](https://cn.vitejs.dev/)
- element-plus [(UI 库)](https://element-plus.org/zh-CN/)
- xgplayer [(西瓜播放器)](https://v2.h5player.bytedance.com/gettingStarted/)
- md-editor-v3 [(Markdown 编辑器)](https://www.wangeditor.com/)

## 部署方法

#### clone 项目

```sh
git clone https://github.com/Ocyss/douyin.git && cd douyin
```

#### 编译/运行

```sh
go build && ./douyin
```

> 项目端口默认`:23724`

#### Web 端配置

##### 1.直接下载 releases

https://github.com/Ocyss/douyin-web/releases

解压到 web 文件夹中,结构如下

```
douyin
├── data
│   ├── config.json
│   └── log
├── web
│  ├── assets
│  ├── docs
│  ├── static
│  └── index.html
└── douyin
```

运行`douyin`,打开`http://localhost:23724`,即可看到 Web 端界面

##### 2.自行编译

```sh
git clone https://github.com/Ocyss/douyin-web.git
```

### 文件目录说明

```
douyin
├── cmd				# 启动项/参数配置
│   └── flags
├── data				# 数据目录
│   └── log
├── internal			# 内部服务
│   ├── bootstrap
│   ├── conf
│   ├── db
│   └── model
├── server			# 路由服务
│   ├── common
│   ├── handlers
│   └── middleware
├── test
├── utils			# 通用工具
│   ├── checks
│   ├── tokens
│   └── upload
└── web				# Web 服务
```

## 预览

![推荐](https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/douyin/shot_2023-08-17_01-23-10.png)
![评论](https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/douyin/shot_2023-08-17_01-23-29.png)
![主页](https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/douyin/shot_2023-08-17_01-23-47.png)
![关注](https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Q/douyin/shot_2023-08-17_01-24-09.png)

### 版本控制

该项目使用 Git 进行版本管理。您可以在 repository 参看当前可用版本。

### 联系方式

​ [me@ocyss.icu](mailto:me@ocyss.icu)

### 团队成员

- [daydayw](https://github.com/daydayw)
- [Godvictory](https://github.com/Godvictory)
- [haoer](https://github.com/haoaer)
- [hblovo](https://github.com/hblovo)
- [koutaManaka](https://github.com/koutaManaka)
- [leaveYoung](https://github.com/leaveYoung)
- [lyhlyh03](https://github.com/lyhlyh03)
- [Ocyss_04](https://github.com/ocyss)
