/**
@File   : router
@Date   : 2023/1/16 11:23 上午
@Author : lyzin
@Desc   :
**/

package serve

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine{
	r := gin.Default()
	FileServeHandler := FileServe{}
	r.GET("/", FileServeHandler.FileList)
	r.GET("/:file/*action", FileServeHandler.FileDownLoadFile)
	return r
}
