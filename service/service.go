package service

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"net/http"
	s2 "webapi/util/server"
)

func init() {
	//设置路由
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/ws", s2.ClientRegister)
		port := beego.AppConfig.String("websocketport")
		log.Infof("Websocket Server started %s ...", port)
		log.Fatal(http.ListenAndServe(port, mux))
	}()
}
