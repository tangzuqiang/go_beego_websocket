package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	data2 "webapi/util/data"
	server2 "webapi/util/server"
)

// PushController operations for Push
type PushController struct {
	BaseController
}

func (c *PushController) Push() {
	// read request
	var pm data2.PushMessage
	decoder := json.NewDecoder(c.Ctx.Request.Body)

	if err := decoder.Decode(&pm); err != nil {
		c.error("Push message error: " + err.Error())
		return
	}

	//数据写入到数据库
	//var pushMsgModel models2.PushMessageModel
	//msgId := pushMsgModel.Create(models2.PushMessageModel{SenderId: pm.SenderId, SenderName: pm.SenderName, Title: pm.Title, Content: pm.Content,
	//	Options: pm.Options, MsgType: pm.MsgType, UserIds: pm.UserIds})

	data := data2.MessageData{SenderId: pm.SenderId, MsgTime: time.Now().Format(data2.TIMESTAMP_FORMAT), SenderName: pm.SenderName, Title: pm.Title, Content: pm.Content,
		Options: pm.Options, MsgId: 1, MsgType: pm.MsgType}
	message, _ := json.Marshal(&data2.ResMessage{Error: 0, Msg: "ok", Event: "message", Data: data})

	if pm.UserIds == "0" { //发全部
		server2.Manager.Broadcast <- message
	} else {
		userIdsArr := strings.Split(pm.UserIds, ",")
		var userIds = make([]interface{}, 0)
		for _, userId := range userIdsArr {
			userId = strings.Trim(userId, " ")
			userId, _ := strconv.ParseInt(userId, 10, 64)
			userIds = append(userIds, userId)
		}

		if len(userIds) > data2.MAX_SEND_USER_NUM && pm.MsgType == 2 {
			c.error(200, fmt.Sprintf("必读消指定用户时最多发送用户量不可超过%d", data2.MAX_SEND_USER_NUM))
			return
		}
		//发送消息到指定用户
		//sendUserIds :=
		server2.Manager.SendMsgToUsers(message, userIds)
		if pm.MsgType == 2 { //如果是必达消息并发模式写入数据库状态为待发送
			go func() {
				//waitSendUserIds := utils2.SliceDiff(userIds, sendUserIds)
				//var pushMsgLogModel models2.PushMessageLogModel
				//pushMsgLogModel.CreateWaiteMessageLogs(waitSendUserIds, msgId, pm.MsgType, time.Now().Format(data2.TIMESTAMP_FORMAT))
			}()
		}
		c.success()
	}
}

func (c *PushController) Getatlinenum() {
	//w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	//w.Header().Set("content-type", "application/json") //返回数据格式是json
	//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//// read request
	//if  r.Method == "OPTIONS"{
	//	return
	//}
	count := 0
	server2.Manager.Clients.Range(func(k, v interface{}) bool {
		count++
		return true
	})

	info := make(map[string]interface{})
	info["count"] = fmt.Sprintf("当前client连接总数：%d\n\n", count)
	var list []string
	server2.Manager.Clients.Range(func(k, v interface{}) bool {
		conn := k.(*server2.Client)
		list = append(list, fmt.Sprintf("ClientID:%s,  UserID:%d\n", conn.ID, conn.UserId))
		return true
	})
	info["list"] = list
	c.success("ok", info)
}
