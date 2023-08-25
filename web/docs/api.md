## /基础

```text
暂无描述
```

#### Header 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### Query 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### Body 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

## /基础/获取视频流

```text
原文档是最新投稿`秒级时间戳`,不够精确

在不增加新字段也不麻烦的情况下,选择使用返回id,且id默认排序
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/feed?latest_time=1234567891234567891&token=eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8

#### 请求方式

> GET

#### Content-Type

> none

#### 请求 Query 参数

| 参数名      | 示例值                                                    | 参数类型 | 是否必填 | 参数描述          |
| ----------- | --------------------------------------------------------- | -------- | -------- | ----------------- |
| latest_time | 1234567891234567891                                       | String   | 否       | 返回视频最后的 id |
| token       | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8 | String   | 否       | 用户鉴权 token    |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
	"next_time": 5074202321428952064,
	"status_code": 0,
	"status_msg": "ok!",
	"video_list": [
		{
			"id": 5074206165047358464,
			"author": {
				"id": 3382802237327496192,
				"name": "ocyss",
				"follow_count": 2,
				"follower_count": 2,
				"is_follow": false,
				"avatar": "https://api.multiavatar.com/ocyss.png",
				"background_image": "https://api.paugram.com/wallpaper/",
				"signature": "此人巨懒",
				"work_count": 0,
				"total_favorited": 1,
				"favorite_count": 7
			},
			"play_url": "https://byte-hunters.oss.aliyuncs.com/t/test2.mp4",
			"cover_url": "https://byte-hunters.oss.aliyuncs.com/t/test2.jpg",
			"favorite_count": 1,
			"comment_count": 1,
			"play_count": 0,
			"is_favorite": false,
			"title": "标题~~~~",
			"desc": ""
		}
	]
}
```

#### 错误响应示例

```javascript
{"errmsg":["token is malformed: could not JSON decode header: unexpected end of JSON input"],"status_code":1,"status_msg":"Token 错误,请重新登录"}
```

## /基础/视频投稿

```text
暂无描述
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/douyin/publish/action/

#### 请求方式

> POST

#### Content-Type

> form-data

#### 请求 Body 参数

| 参数名 | 示例值                                                    | 参数类型 | 是否必填 | 参数描述       |
| ------ | --------------------------------------------------------- | -------- | -------- | -------------- |
| data   | -                                                         | String   | 是       | 投稿视频文件   |
| token  | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8 | String   | 是       | 用户鉴权 token |
| title  | 标题                                                      | String   | 是       | 视频标题       |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
    "status_code":0,
    "status_msg":"",
    "vid":1234567899876543210
}
```

#### 错误响应示例

```javascript
{
	"status_code": 1,
	"status_msg": "投稿失败...",
	"errmsg": [
		"Err..."
	]
}
```

## /基础/获取发布列表

```text
暂无描述
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/publish/list/?token=eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8&user_id=3382802237327496156

#### 请求方式

> GET

#### Content-Type

> json

#### 请求 Query 参数

| 参数名  | 示例值                                                    | 参数类型 | 是否必填 | 参数描述       |
| ------- | --------------------------------------------------------- | -------- | -------- | -------------- |
| token   | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8 | String   | 是       | 用户鉴权 token |
| user_id | 3382802237327496156                                       | String   | 是       | 用户 id        |

#### 请求 Body 参数

```javascript

```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
	"status_code": 0,
	"status_msg": "ok!",
	"video_list": [
		{
			"id": 5074206165027444736,
			"author": {
				"id": 3382802237327496192,
				"name": "ocyss",
				"follow_count": 2,
				"follower_count": 2,
				"is_follow": false,
				"avatar": "https://api.multiavatar.com/ocyss.png",
				"background_image": "https://api.paugram.com/wallpaper/",
				"signature": "此人巨懒",
				"work_count": 0,
				"total_favorited": 1,
				"favorite_count": 7
			},
			"play_url": "https://byte-hunters.oss.aliyuncs.com/t/test1.mp4",
			"cover_url": "https://byte-hunters.oss.aliyuncs.com/t/test1.jpg",
			"favorite_count": 0,
			"comment_count": 0,
			"play_count": 0,
			"is_favorite": false,
			"title": "test1",
			"desc": ""
		},
	]
}
```

#### 错误响应示例

```javascript
{"errmsg":["Key: 'userReqs.Token' Error:Field validation for 'Token' failed on the 'required' tag"],"status_code":1,"status_msg":"参数不正确"}
```

## /基础/视频投稿(test)

```text
暂无描述
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/publish/actionUrl

#### 请求方式

> POST

#### Content-Type

> json

#### 请求 Body 参数

```javascript
{
	"id": 8795259707917331,
	"url": "https://byte-hunters.oss.aliyuncs.com/t/test2.mp4",
	"title": "况该位观广",
	"cover_url": "https://byte-hunters.oss.aliyuncs.com/t/test2.jpg",
	"user_creations": "条形江积局不"
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
    "status_code":0,
    "status_msg":"",
    "vid":1234567899876543210
}
```

#### 错误响应示例

```javascript
{
	"status_code": 0,
	"status_msg": "投稿失败..",
	"errmsg": [
		"Err..."
	]
}
```

## /互动

```text
暂无描述
```

#### Header 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### Query 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### Body 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

## /互动/点赞

```text
暂无描述
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/favorite/action/

#### 请求方式

> POST

#### Content-Type

> json

#### 请求 Body 参数

```javascript
{
	"token": "eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8",
	"video_id": "5074202302310375752",
	"action_type": "1"
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{"status_code":0,"status_msg":"ok!"}
```

#### 错误响应示例

```javascript
{"errmsg":["Error"],"status_code":1,"status_msg":"网卡了,再试一次吧"}
```

## /互动/获取喜欢列表

```text
暂无描述
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/favorite/list/?token=eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8&user_id=3382802237327496192

#### 请求方式

> GET

#### Content-Type

> none

#### 请求 Query 参数

| 参数名  | 示例值                                                    | 参数类型 | 是否必填 | 参数描述                          |
| ------- | --------------------------------------------------------- | -------- | -------- | --------------------------------- |
| token   | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8 | String   | 是       | 用户鉴权 token                    |
| user_id | 3382802237327496192                                       | String   | 是       | 用户 id,为空/0 获取自己的喜欢列表 |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
	"status_code": 0,
	"status_msg": "ok!",
	"video_list": [
		{
			"id": 5074202302310375424,
			"author": {
				"id": 3382800949025138688,
				"name": "musicc5",
				"follow_count": 0,
				"follower_count": 0,
				"is_follow": false,
				"avatar": "https://api.multiavatar.com/musicc5.png",
				"background_image": "https://api.paugram.com/wallpaper/",
				"signature": "此人巨懒",
				"work_count": 0,
				"total_favorited": 1,
				"favorite_count": 0
			},
			"play_url": "https://byte-hunters.oss.aliyuncs.com/y2/175.mp4",
			"cover_url": "https://byte-hunters.oss.aliyuncs.com/y2/175.jpg",
			"favorite_count": 1,
			"comment_count": 1,
			"play_count": 0,
			"is_favorite": true,
			"title": "学习简单的吉他弹唱技巧",
			"desc": ""
		},
	]
}
```

#### 错误响应示例

```javascript
{"errmsg":["strconv.ParseInt: parsing \".er\": invalid syntax"],"status_code":1,"status_msg":"参数不正确"}
```

## /互动/评论

```text
暂无描述
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/comment/action/

#### 请求方式

> POST

#### Content-Type

> json

#### 请求 Body 参数

```javascript
{
	"token": "eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8",
	"video_id": 5074202322353033216,
	"action_type": 1,
	"comment_text": "第一",
	"comment_id": "3382937143857520084"
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
	"comment": {
		"id": 3382937143857520128,
		"user": {
			"id": 3382800949287004672,
			"name": "musicc6",
			"follow_count": 3,
			"follower_count": 3,
			"is_follow": false,
			"avatar": "https://api.multiavatar.com/musicc6.png",
			"background_image": "https://api.paugram.com/wallpaper/",
			"signature": "此人巨懒",
			"work_count": 0,
			"total_favorited": 0,
			"favorite_count": 0
		},
		"content": "第一!!",
		"reply_id": 0
	},
	"status_code": 0,
	"status_msg": "ok!"
}
```

#### 错误响应示例

```javascript
{"errmsg":["Key: 'commentReqs.Token' Error:Field validation for 'Token' failed on the 'required' tag\nKey: 'commentReqs.VideoId' Error:Field validation for 'VideoId' failed on the 'required' tag\njson: cannot unmarshal string into Go struct field commentReqs.video_id of type int64"],"status_code":1,"status_msg":"参数不正确"}
```

## /互动/获取评论列表

```text
暂无描述
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/comment/list/?token=eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8&video_id=5074202322353033506

#### 请求方式

> GET

#### Content-Type

> none

#### 请求 Query 参数

| 参数名   | 示例值                                                    | 参数类型 | 是否必填 | 参数描述       |
| -------- | --------------------------------------------------------- | -------- | -------- | -------------- |
| token    | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8 | String   | 是       | 用户鉴权 token |
| video_id | 5074202322353033506                                       | String   | 是       | 视频 id        |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
	"comment_list": [
		{
			"id": 3382935911244216320,
			"user": {
				"id": 3382800949287004672,
				"name": "musicc6",
				"follow_count": 3,
				"follower_count": 3,
				"is_follow": false,
				"avatar": "https://api.multiavatar.com/musicc6.png",
				"background_image": "https://api.paugram.com/wallpaper/",
				"signature": "此人巨懒",
				"work_count": 0,
				"total_favorited": 0,
				"favorite_count": 0
			},
			"content": "哈哈",
			"reply_id": 0
		}
	],
	"status_code": 0,
	"status_msg": "ok!"
}
```

#### 错误响应示例

```javascript
{"errmsg":["Key: 'commentReqs.Token' Error:Field validation for 'Token' failed on the 'required' tag\nEOF"],"status_code":1,"status_msg":"参数不正确"}
```

## /社交

```text
暂无描述
```

#### Header 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### Query 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### Body 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

## /社交/关注操作

```text
根据 `action_type` 来进行关注和取消关注
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/relation/action/

#### 请求方式

> POST

#### Content-Type

> json

#### 请求 Body 参数

```javascript
{
	"token": "eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8",
	"to_user_id": "3382800950603989432",
	"action_type": "1"
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{"status_code":0,"status_msg":"ok!"}
```

#### 错误响应示例

```javascript
{"errmsg":["Key: 'relationReqs.Token' Error:Field validation for 'Token' failed on the 'required' tag\nEOF"],"status_code":1,"status_msg":"参数不正确"}
```

## /社交/获取关注列表

```text
必须提供`Token`,没`user_id`时为查询自己的关注列表

当有`user_id`则查询对方的关注列表
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/relation/follow/list/?token=eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8&user_id=3382802237327496156

#### 请求方式

> GET

#### Content-Type

> none

#### 请求 Query 参数

| 参数名  | 示例值                                                    | 参数类型 | 是否必填 | 参数描述       |
| ------- | --------------------------------------------------------- | -------- | -------- | -------------- |
| token   | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8 | String   | 是       | 用户鉴权 token |
| user_id | 3382802237327496156                                       | String   | 是       | 用户 id        |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
	"status_code": 0,
	"status_msg": "ok!",
	"user_list": [
		{
			"id": 3382800949287004672,
			"name": "musicc6",
			"follow_count": 1,
			"follower_count": 1,
			"is_follow": false,
			"avatar": "https://api.multiavatar.com/musicc6.png",
			"background_image": "https://api.paugram.com/wallpaper/",
			"signature": "此人巨懒",
			"work_count": 0,
			"total_favorited": 0,
			"favorite_count": 0
		},
		{
			"id": 3382800950603989504,
			"name": "musicc11",
			"follow_count": 0,
			"follower_count": 0,
			"is_follow": false,
			"avatar": "https://api.multiavatar.com/musicc11.png",
			"background_image": "https://api.paugram.com/wallpaper/",
			"signature": "此人巨懒",
			"work_count": 0,
			"total_favorited": 0,
			"favorite_count": 0
		}
	]
}
```

#### 错误响应示例

```javascript
{"errmsg":["strconv.ParseInt: parsing \"er\": invalid syntax\nEOF"],"status_code":1,"status_msg":"参数不正确"}
```

## /社交/获取粉丝列表

```text
必须提供`Token`,没`user_id`时为查询自己的粉丝列表

当有`user_id`则查询对方的粉丝列表
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/relation/follower/list/?token=eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8&user_id=3382802237327496156

#### 请求方式

> GET

#### Content-Type

> none

#### 请求 Query 参数

| 参数名  | 示例值                                                    | 参数类型 | 是否必填 | 参数描述       |
| ------- | --------------------------------------------------------- | -------- | -------- | -------------- |
| token   | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8 | String   | 是       | 用户鉴权 token |
| user_id | 3382802237327496156                                       | String   | 是       | 用户 id        |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{"status_code":0,"status_msg":"ok!","user_list":[{"id":3382800949287004698,"name":"musicc6","follow_count":1,"follower_count":1,"is_follow":false,"avatar":"https://api.multiavatar.com/musicc6.png","background_image":"https://api.paugram.com/wallpaper/","signature":"此人巨懒","work_count":0,"total_favorited":0,"favorite_count":0}]}
```

#### 错误响应示例

```javascript
{"errmsg":["strconv.ParseInt: parsing \"err\": invalid syntax\nEOF"],"status_code":1,"status_msg":"参数不正确"}
```

## /社交/获取好友列表

```text
必须提供`Token`,且只能获取自己的好友列表

还会带上和该好友的最后聊天记录,如果没聊天过则显示为"快来和你的新朋友聊天吧"
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/relation/friend/list/?token=eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8&user_id=3382802237327496156

#### 请求方式

> GET

#### Content-Type

> none

#### 请求 Query 参数

| 参数名  | 示例值                                                                                                                                                                                          | 参数类型 | 是否必填 | 参数描述       |
| ------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- | -------- | -------------- |
| token   | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8                                                                                                                                       | String   | 是       | -              |
| user_id | 3382802237327496156                                                                                                                                                                             | String   | 是       | -              |
| token   | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MzM4MjgwMjIzNzMyNzQ5NjE1NiwidXNlcm5hbWUiOiJvY3lzcyIsImlzcyI6IkJ5dGVIdW50ZXJzIiwiZXhwIjoxNjk5MTc3MTE4fQ.Je9tZWxiQxRyJMaCYMO471bd36co9g7wfnPveB6vAAg | String   | 是       | 用户鉴权 token |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
	"status_code": 0,
	"status_msg": "ok!",
	"user_list": [
		{
			"id": 3382800949287004672,
			"name": "musicc6",
			"follow_count": 1,
			"follower_count": 1,
			"is_follow": false,
			"avatar": "https://api.multiavatar.com/musicc6.png",
			"background_image": "https://api.paugram.com/wallpaper/",
			"signature": "此人巨懒",
			"work_count": 0,
			"total_favorited": 0,
			"favorite_count": 0,
            "message":"快来和你的新朋友聊天吧",
            "msg_type":0
		}
	]
}
```

#### 错误响应示例

```javascript
{"errmsg":["strconv.ParseInt: parsing \"e\": invalid syntax\nEOF"],"status_code":1,"status_msg":"参数不正确"}
```

## /社交/消息操作

```text
目前只实现了发送消息
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/message/action/

#### 请求方式

> POST

#### Content-Type

> json

#### 请求 Body 参数

```javascript
{
	"token": "eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8",
	"to_user_id": "3382802237327496156",
	"content": "线么院至",
	"action_type": 1
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{"status_code":0,"status_msg":"ok!"}
```

#### 错误响应示例

```javascript
{"errmsg":["Key: 'messageReqs.Token' Error:Field validation for 'Token' failed on the 'required' tag\nKey: 'messageReqs.ToUserId' Error:Field validation for 'ToUserId' failed on the 'required' tag\njson: cannot unmarshal string into Go struct field messageReqs.to_user_id of type int64"],"status_code":1,"status_msg":"参数不正确"}
```

## /社交/获取聊天记录

```text
通过`pre_msg_time`毫秒级时间戳来轮询获取聊天记录
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/message/chat/?token=eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8&to_user_id=3382802237327496156&pre_msg_time=1691467224691

#### 请求方式

> GET

#### Content-Type

> none

#### 请求 Query 参数

| 参数名       | 示例值                                                    | 参数类型 | 是否必填 | 参数描述       |
| ------------ | --------------------------------------------------------- | -------- | -------- | -------------- |
| token        | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8 | String   | 是       | 用户鉴权 token |
| to_user_id   | 3382802237327496156                                       | String   | 是       | 用户 id        |
| pre_msg_time | 1691467224691                                             | String   | 是       | 毫秒级时间戳   |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
	"message_list": [
		{
			"id": 5074402551676547072,
			"create_time": 1691467517226,
			"to_user_id": 3382802237327496192,
			"from_user_id": 3382800949287004672,
			"content": "你好啊~~~~~~"
		}
	],
	"status_code": 0,
	"status_msg": "ok!"
}
```

#### 错误响应示例

```javascript
{"errmsg":["Key: 'messageReqs.ToUserId' Error:Field validation for 'ToUserId' failed on the 'required' tag\nEOF"],"status_code":1,"status_msg":"参数不正确"}
```

## /用户

```text
暂无描述
```

#### Header 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### Query 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### Body 参数

| 参数名 | 示例值 | 参数描述 |
| ------ | ------ | -------- |

暂无参数

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

## /用户/用户信息

```text
根据`id`和`token`来判断是否合法,然后返回该用户的信息
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/user/?user_id=ocyss&token=eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8

#### 请求方式

> GET

#### Content-Type

> none

#### 请求 Query 参数

| 参数名  | 示例值                                                    | 参数类型 | 是否必填 | 参数描述       |
| ------- | --------------------------------------------------------- | -------- | -------- | -------------- |
| user_id | ocyss                                                     | String   | 是       | 用户 id        |
| token   | eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8 | String   | 是       | 用户鉴权 token |

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{
	"status_code": 0,
	"status_msg": "Success",
	"user": {
		"user_id": 6762478695648966000,
		"name": "Ocyss",
		"follow_count": 8888,
		"follower_count": 8888,
		"is_follow": true,
		"avatar": "https://api.multiavatar.com/ocyss.png",
		"background_image": "https://api.paugram.com/wallpaper/",
		"signature": "此人巨懒",
		"work_count": 666,
		"total_favorited": 888,
		"favorite_count": 888
	}
}
```

#### 错误响应示例

```javascript
{
	"errmsg": [
		"err"
	],
	"status_code": 1,
	"status_msg": "Token 认证失败"
}
```

## /用户/注册

```text
根据用户提供的
- 用户名称*
- 用户密码*
- 用户头像
- 用户个人页顶部大图
- 个人简介

信息创建一个新账户,并返回其id和token
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/user/register/

#### 请求方式

> POST

#### Content-Type

> json

#### 请求 Body 参数

```javascript
{
	"username": "ocyss",
	"password": "230724",
	"avatar": "http://fkrnvcfxq.dz/rwuedjos",
	"background_image": "http://paj.edu/gayenukqbl",
	"signature": "此人巨懒"
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{"status_code":0,"status_msg":"Success","token":"eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8","user_id":6762478695648965464}
```

#### 错误响应示例

```javascript
{
	"status_code": 1,
	"status_msg": "参数不正确",
	"errmsg": [
		"Err..."
	]
}
```

## /用户/登录

```text
根据用户名和密码,进行登录验证
登录成功会返回当前用户id和鉴权token
```

#### 接口状态

> 已完成

#### 接口 URL

> http://localhost:23724/douyin/user/login/

#### 请求方式

> POST

#### Content-Type

> json

#### 请求 Body 参数

```javascript
{
	"username": "需法了电层它热",
	"password": "提对争派价术"
}
```

#### 认证方式

```text
noauth
```

#### 预执行脚本

```javascript
暂无预执行脚本;
```

#### 后执行脚本

```javascript
暂无后执行脚本;
```

#### 成功响应示例

```javascript
{"status_code":0,"status_msg":"Success","token":"eyJhbGciOiJIUJ9.eyJpZCI6Njc2MjQ3ODY2NzQ1fQ.L8wBmGzb8Tl-S8","user_id":6762478695648965464}
```

#### 错误响应示例

```javascript
{
	"status_code": 1,
	"status_msg": "参数不正确",
	"errmsg": [
		"Err..."
	]
}
```
