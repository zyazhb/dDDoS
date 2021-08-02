package main

import (
	"main/server/web"
)

func main() {
	// TODO：暂时只运行web服务，剩下的交互看是否合适放在web端上边
	/*
		1. 监控端与链端的启动
		2. 动态获取client的状态 -> (ATTACKED, RUNNING, RESTARTING DEVICE, BLOCK IP)
		3. 动态分配生成账户
		4. 分布式client与主node的信息交互
		5. ... -> ...
	*/

	web.RunWeb()
}
