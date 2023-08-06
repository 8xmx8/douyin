package common

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func Bind(c *gin.Context, data any) error {
	if err := c.ShouldBindQuery(data); err != nil {
		if err2 := c.ShouldBindJSON(data); err2 != nil {
			return errors.Join(err, err2)
		}
	}
	return nil
}
