/**
@File   : server
@Date   : 2023/1/16 12:59 下午
@Author : lyzin
@Desc   :
**/

package serve

import "fmt"

func StartServe() {
	// 路由
	r := InitRouter()
	localHost := GetHostIp()
	if len(localHost) == 0 {
		localHost = "0.0.0.0"
	}
	fmt.Printf("Service started successfully, Please visit: http://%v:8000\n\n", localHost)
	err := r.Run(":8000")
	if err != nil {
		panic(fmt.Sprintf("Start server err:%v\n", err))
	}
}
