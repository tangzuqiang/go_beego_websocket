package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	utils2 "webapi/util"
	data2 "webapi/util/data"
)

type Client struct {
	ID           string
	Socket       *websocket.Conn
	Send         chan []byte
	UserId       int64
	Token        string
	RegisterTime int64
}

//客户端连接消息读取
func (c *Client) Read() {
	for {
		_, _, err := c.Socket.ReadMessage()
		if err != nil {
			c.Close()
			break
		}
	}
}

//客户端消息写入
func (c *Client) Write() {
	defer c.Close()

	for {
		select {
		case message, ok := <-c.Send: //发送数据
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			var res data2.ResMessage
			json.Unmarshal(message, &res)

			go c.SaveUserMsgLog(res.Data, 2)
			//if !c.CheckToken() {
			//	c.CloseAndRes(200, "token过期", "messaage")
			//	break
			//}

			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				go c.SaveUserMsgLog(res.Data, 2)
				c.Close()
				break
			}

			go c.SaveUserMsgLog(res.Data, 1)
		}
	}
}

//保存用户消息日志
func (c *Client) SaveUserMsgLog(data data2.MessageData, status int) {
	//var pushMsgLogModel models2.PushMessageLogModel
	//if data.MsgLogId == 0 {
	//	pushMsgLogModel.Create(data.MsgId, data.MsgType, c.UserId, c.ID, status)
	//} else {
	//	pushMsgLogModel.Save(data.MsgLogId, c.ID, status)
	//}
}

//发送给用户最近未读的必达消息
func (c *Client) SendMustReadMsg() {
	//var pushMsgLogModel models2.PushMessageLogModel
	//msgData := pushMsgLogModel.GetMustReadMsgByUserId(c.UserId, time.Now().Unix()-data2.LAST_MSG_TIME_LIMIT)
	//for _, row := range msgData {
	//	message, _ := json.Marshal(&data2.ResMessage{Error: 0, Msg: "ok", Event: "message", Data: row})
	//
	//	//业务更新消息
	//	if row.MsgType == 3 {
	//		c.Send <- message
	//		continue
	//	}
	//	popType, options, msg := Manager.GetMsgPopType(message)
	//	//发送element通知
	//	if popType == "ele" || popType == "all" {
	//		c.Send <- Manager.UpdateMsgPopType("ele", options, msg)
	//	}
	//
	//	//发送浏览器通知
	//	if popType == "browser" || popType == "all" {
	//		c.Send <- Manager.UpdateMsgPopType("browser", options, msg)
	//	}
	//
	//}
}

//发送消息
func (c *Client) SendRes(res data2.ResMessage) interface{} {
	resJson, _ := json.Marshal(res)
	err := c.Socket.WriteMessage(websocket.TextMessage, resJson)
	if err != nil {
		log.Error("Send Res Error:" + c.ID)
		c.Close()
		return "发送错误:" + c.ID + err.Error()
	}
	return true
}

//关闭连接
func (c *Client) Close() {
	Manager.Unregister <- c
	c.Socket.Close()
}

//发送错误消息并关闭连接
func (c *Client) CloseAndRes(errCode int, msg string, event string) {
	c.SendRes(data2.ResMessage{Error: errCode, Msg: msg, Event: event})
	c.Close()
}

//check user login token
func (c *Client) CheckToken() bool {
	return true
	redisClient := utils2.BaseRedis.Connect("default")
	defer redisClient.Close()

	if len(c.Token) < 8 {
		log.Debug("token length error")
		return false
	}

	key := fmt.Sprintf(data2.XHX_SESSION_NAME+"%d", c.UserId)
	log.Debug("redis_key" + key)
	token, err := redisClient.Get(key).Result()
	fmt.Println("错误信息")
	fmt.Println("token" + token)
	if err == redis.Nil {
		log.Debug("register token key not exist")
		return false
	}

	if token != c.Token {
		log.Debug("token expired")
		return false
	}

	return true
}

//获取客户端链接数量
func (c *Client) GetClientNumByUserId(userId int64) int {
	count := 0
	Manager.Clients.Range(func(k, v interface{}) bool {
		conn := k.(*Client)
		if conn.UserId == userId {
			count++
		}
		return true
	})

	return count
}
