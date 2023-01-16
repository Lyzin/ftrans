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
	err := r.Run(":8000")
	if err != nil {
		panic(fmt.Sprintf("start file server err:%v\n", err))
	}
}

