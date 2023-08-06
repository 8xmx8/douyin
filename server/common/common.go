package common

import (
	"github.com/Godvictory/douyin/cmd/flags"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 2023-08-02 20:28:15 决定废弃此方案,重新封装路由
const (
	statusOk  = 0
	statusErr = 1
)

// OK 成功的请求
func OK(c *gin.Context, data ...map[string]any) {
	res := gin.H{
		"status_code": statusOk,
		"status_msg":  "Success",
	}
	for d := range data {
		for k, v := range data[d] {
			res[k] = v
		}
	}
	c.JSON(http.StatusOK, res)
}

// Err 失败的请求
// 推荐所有错误返回都带上err，包括参数错误，在debug和dev模式下能快速定位bug
func Err(c *gin.Context, msg string, err ...error) {
	res := gin.H{
		"status_code": statusErr,
		"status_msg":  msg,
	}
	// 调试与开发模式，返回错误消息。
	if (flags.Dev || flags.Debug) && len(err) > 0 {
		errs := make([]string, len(err))
		for i, e := range err {
			if e != nil {
				errs[i] = e.Error()
			}
		}
		res["errmsg"] = errs
	}
	c.JSON(http.StatusOK, res)
}

// ErrParam 参数错误封装
func ErrParam(c *gin.Context, err ...error) {
	Err(c, "参数不正确", err...)
}
