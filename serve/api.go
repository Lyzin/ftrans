/**
@File   : api
@Date   : 2023/1/16 11:24 上午
@Author : lyzin
@Desc   :
**/

package serve

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

type FileServe struct {}

var currentPath = GetFilePathCurrent()

func (f FileServe)FileList(c *gin.Context) {
	c.File(currentPath)
}

func (f FileServe) FileDownLoadFile(c *gin.Context) {
	file := c.Param("file")
	
	action := c.Param("action")
	
	// 下载链接路径
	newDownLoadPath := currentPath + "/" + file + action
	
	if !ObjIsDir(newDownLoadPath) {
		downLoadFileName := filepath.Base(newDownLoadPath)
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		c.Writer.Header().Add("Content-Disposition", "attachment; filename=" + downLoadFileName)
	}
	c.File(newDownLoadPath)
}