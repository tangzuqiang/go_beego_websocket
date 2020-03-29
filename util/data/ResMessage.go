package data

const XHX_SESSION_NAME = "php:authtoken:1:"      //token session name
const CLIENT_TIMEOUT = 43200                     //客户端连接最大时间
const MAX_SEND_USER_NUM = 1000                   //必达消息最大发送用户量
const LAST_MSG_TIME_LIMIT = 3 * 24 * 3600        //读取最近必达消息时间限制
const LAST_MSG_NUM_LIMIT = 20                    //读取最近必达消息数量限制
const TIMESTAMP_FORMAT = "2006-01-02 15:04:05"   //日期格式
const CLIENT_REGISTER_TIMEOUT = 5                //客户端注册登录超时时间 5秒
const DOWNLOAD_LIST_KEY = "xhx_download_list_%d" //下载队列的key
const USER_MAX_CLIENT_NUM = 10                   //每个用户最大客户端连接时间
//客户端注册消息
type RegisterMessage struct {
	Token string
	Event string
}

//服务端返回消息
type ResMessage struct {
	Error int         `json:"error"`
	Msg   string      `json:"msg"`
	Data  MessageData `json:"data"`
	Event string      `json:"event"`
}

//推送数据结构
type PushMessage struct {
	SenderId   int64  //发送者id
	SenderName string //发送者姓名
	MsgType    int    //消息类型 1发送在线用户即时消息 2登录后必达消息 3业务内容更新消息
	Title      string //消息标题
	Content    string //消息内容
	UserIds    string //用户id以,号分隔 msgType为2时userIds必传
	Options    string //弹窗选项目前支持 duration(毫秒), position, type参数（对应elementUi通知组件参数）
	Timestamp  string //时间戳
}

//发送给客端data消息数据结构
type MessageData struct {
	SenderId   int64  `json:"senderId"`
	SenderName string `json:"senderName"`
	MsgTime    string `json:"msgTime"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Options    string `json:"options"`
	MsgId      int64  `json:"msgId,omitempty"`    //推送消息数据库记录id
	MsgType    int    `json:"msgType,omitempty"`  //消息类型
	MsgLogId   int64  `json:"msgLogId,omitempty"` //用户消息数据库记录id
}

//消息选项数据结构
type MessageOptions struct {
	Duration int    `json:"duration"`
	Position string `json:"position"`
	Type     string `json:"type"`
	PopType  string `json:"popType"`
}
