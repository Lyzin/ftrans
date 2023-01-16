/**
@File   : UTILS
@Date   : 2023/1/16 12:56 下午
@Author : lyzin
@Desc   :
**/

package serve

import (
	"os"
)


func GetFilePathCurrent() string {
	currentPath, _ := os.Getwd()
	return currentPath
}

func ObjIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}