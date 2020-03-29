package main

import (
	"github.com/astaxie/beego"
	_ "webapi/routers"
	_ "webapi/service"
	"webapi/util/server"
)

func main() {
	//启动客户端管理
	go server.Manager.Start()
	//关闭超时连接客户端
	go server.Manager.CloseTask()
	beego.Run()

}
