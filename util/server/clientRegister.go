package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	data2 "webapi/util/data"
)

var upgraded = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//客户端注册
func ClientRegister(res http.ResponseWriter, req *http.Request) {
	//res.Write([]byte("链接调试"))
	//return
	//解析一个连接
	fmt.Println("握手")
	conn, error := upgraded.Upgrade(res, req, nil)
	if error != nil {
		fmt.Println(error.Error())
		io.WriteString(res, "这是一个websocket.")
		return
	}

	uid := uuid.NewV4()
	sha1 := uid.String()

	//初始化一个客户端对象
	client := &Client{ID: sha1, Socket: conn, Send: make(chan []byte), RegisterTime: time.Now().Unix()}
	//注册一个对象到channel
	Manager.Register <- client
	go client.Read()
	go client.Write()

	uriParse := strings.Split(req.RequestURI, "?")
	if len(uriParse) < 2 {
		client.CloseAndRes(100, "参数错误(1).", "register")
		return
	}

	params, _ := url.ParseQuery(uriParse[1])
	fmt.Println(params)
	if len(params["token"]) == 0 || len(params["uid"]) == 0 {
		client.CloseAndRes(101, "参数错误(2).", "register")
		return
	}

	client.Token = params["token"][0]
	client.UserId, _ = strconv.ParseInt(params["uid"][0], 10, 64)
	//检查token参数
	if !client.CheckToken() {
		client.CloseAndRes(102, "token 验证失败.", "register")
		return
	}

	//检查是否超过最大连接数
	if client.GetClientNumByUserId(client.UserId) > data2.USER_MAX_CLIENT_NUM {
		client.CloseAndRes(103, "超过最大客户端连接数", "register")
		return
	}

	jsonMessage, _ := json.Marshal(&data2.ResMessage{Error: 0, Msg: "ok", Event: "register"})
	client.Send <- jsonMessage
	//发送必读消息
	go client.SendMustReadMsg()
}
