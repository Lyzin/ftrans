/**
@File   : router
@Date   : 2023/1/16 11:23 上午
@Author : lyzin
@Desc   :
**/

package serve

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	FileServeHandler := FileServe{}

	// 添加中间件
	r.Use(func(c *gin.Context) {
		// 设置eTag
		eTag := strconv.FormatInt(time.Now().UnixMilli(), 10)
		// 检查If-None-Match
		match := c.GetHeader("If-None-Match")
		if match == eTag {
			c.Header("ETag", eTag)
			c.Status(http.StatusNotModified)
			return
		}
		c.Header("Cache-Control", "no-cache")
		c.Header("ETag", eTag)
		c.Next()
	})

	r.GET("/", FileServeHandler.FileList)
	r.GET("/:file/*action", FileServeHandler.FileDownLoadFile)
	return r
}
