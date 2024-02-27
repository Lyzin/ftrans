/**
@File   : UTILS
@Date   : 2023/1/16 12:56 下午
@Author : lyzin
@Desc   :
**/

package serve

import (
	"net"
	"os"
	"strings"
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

func GetHostIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return ""
	}
	addr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(addr.String(), ":")[0]
	return ip
}
