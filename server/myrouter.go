package server

import (
	"github.com/Godvictory/douyin/cmd/flags"
	"github.com/Godvictory/douyin/server/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyHandler func(*gin.Context) (int, any)

// decorator 装饰器
func decorator() func(h MyHandler) gin.HandlerFunc {
	return func(h MyHandler) gin.HandlerFunc {
		return func(c *gin.Context) {
			code, data := h(c)
			req := gin.H{
				"status_code": code,
				"status_msg":  "",
			}
			c.Set("code", code)
			if code == 0 {
				// 判断数据类型
				if val, ok := data.(handlers.H); ok {
					for k, v := range val {
						req[k] = v
					}
				}

				req["status_msg"] = "ok!"
				c.JSON(200, req)
			} else {
				switch data.(type) {
				case string:
					c.Set("msg", data)
					req["status_msg"] = data
				case error:
					// 判断是否debug模式，是的话返回错误信息
					if flags.Dev || flags.Debug || flags.Tst {
						req["errmsg"] = data.(error).Error()
					}
				case handlers.MyErr:
					e := data.(handlers.MyErr)
					req["status_msg"] = e.Msg
					c.Set("msg", e.Msg)
					// 判断是否debug模式，是的话返回错误信息
					if flags.Dev || flags.Debug || flags.Tst {
						errs := make([]string, 0, 10)
						for i := range e.Errs {
							if e.Errs[i] == nil {
								continue
							}
							errs = append(errs, e.Errs[i].Error())
						}
						req["errmsg"] = errs
					}
				}

				c.JSON(http.StatusOK, req)
			}
		}
	}
}

func newRouter(group *gin.RouterGroup, method string, path string, handler MyHandler, handlers ...gin.HandlerFunc) {
	if handler != nil {
		// 未开发的路由传nil,不挂载
		group.Handle(method, path, append(handlers, decorator()(handler))...)
	}
}
