<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <!-- import CSS -->
<link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css">
</head>
<body>
<div id="app">
  <el-form ref="form" :model="form" label-width="80px">
    <el-form-item label="消息列表">
      <el-input type="textarea" v-model="form.desc"
                :autosize="{ minRows: 2, maxRows: 20}"

      ></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="refresh">刷新</el-button>

    </el-form-item>
    <el-form-item label="目前在线">
      <el-input type="textarea" v-model="form.at_line"
                :autosize="{ minRows: 2, maxRows: 10}"
      ></el-input>
    </el-form-item>
    <el-form-item label="发送内容">
      <el-input v-model="form.sendName"></el-input>
    </el-form-item>
    <el-form-item label="发送类型">
      <el-select v-model="form.type" placeholder="请选择类型">
         <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
        </el-option>>
      </el-select>
    </el-form-item>
    <el-form-item label="发送给指定用户id" style="width: 20%;">
      <el-input v-model="form.sendUserId"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click.stop="onSubmit">推送</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
 <div >
 </div>
</div>
</body>
<!-- import Vue before Element -->
<script src="https://cdn.jsdelivr.net/npm/vue@2.6.11"></script>
<!-- import JavaScript -->
<!-- 引入样式 -->
<!-- 引入组件库 -->
<script src="https://unpkg.com/element-ui/lib/index.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
<script>
  new Vue({
    el: '#app',
    data() {
      return {
        visible: false,
        input:'',
        value:'',
        options:[
                //1发送在线用户即时消息 2登录后必达消息 3业务内容更新消息
          {value:'1',label:'发送在线用户即时消息'},
          {value:'2',label:'登录后必达消息'},
          {value:'3',label:'业务内容更新消息'},
        ],
        ws:'',
        interval:'',
        retryConnect:false,
        form: {
          sendName: '',
          sendUserId:'',
          type: '',
          at_line: '',
          date1: '',
          date2: '',
          delivery: false,

          resource: '',
          desc: ''
        }
      }
    },
    created() {
      this.init()
    },
    methods: {
      init() {
        if (!window["WebSocket"]) {
          console.log('not support websocket')
          return
        }

        var that = this;
        let uid = new Date().getTime()
        uid = uid.toString().substr(8)
        this.ws = new WebSocket("ws://127.0.0.1:889/ws?uid="+uid+"&token=00000063_d10f2dd30c087a0573d54e4767640253279");
        console.log(this.ws)
        this.ws.onerror = function () {
          console.log("WebSocket error observed");
          setTimeout(function() {
            that.init()
          }, 15000);
        };
        this.ws.onclose = function(e) {
          clearInterval(that.interval)
          if(!that.retryConnect) {
            return
          }
          console.log('push connection is close, retry connect after 5 seconds')
          setTimeout(function() {
            that.init()
          }, 5000);
        }
        this.ws.addEventListener('open', function (e,a) {
          //登录
          that.ws.send('{"event":"register", "token":"00000063_d10f2dd30c087a0573d54e4767640253279","uid":'+uid+'}');
        });
        this.ws.addEventListener('register',function (e) {
          console.log(e)
        })
        this.ws.addEventListener("message", function(e) {

          let res = JSON.parse(e.data)
          console.log(e)
          //token过期
          if(res.error == 100) {
            console.log(res)
            that.retryConnect = false
            return
          }

          if(res.error != 0) {
            console.log(res.msg)
            return
          }

          //client注册消息
          if(res.event == 'register') {
            console.log('ws connection register success ')
            that.interval = setInterval(function() {
              //保此常连接心跳
              that.ws.send('{}')
            }, 60000)
            that.retryConnect = true
            return;
          }

          if(res.event == 'message') {
            let options = JSON.parse(res.data.options);
            that.form.desc += "收到消息"+ res.data.content+"\r\n"
            that.$notify.info({
              title: res.data.title != '' ? res.data.title : '通知',
              message: res.data.content,
              duration: options.duration,
              position: options.position
            });
          }
        })
      },
      onSubmit(){
        let param = {
          sendName: this.form.sendName,
          sendUserId: this.form.sendUserId,
          msgType: parseInt(this.form.type),
          title: "测试标题",
          content: this.form.sendName,
          userIds: this.form.sendUserId,
          options: 'browser',////弹窗选项目前支持 duration(毫秒), position, type参数（对应elementUi通知组件参数）
                             //popType 弹出类型 ele为elementUi通知，browser为浏览器通知，all为elementUi通知 + 浏览器通知
          timestamp: new Date().getTime().toString()
        }
        let self = this
        axios.post('/push', param)
                .then(function (response) {
                  let {data} = response
                  self.form.desc += "您发送消息给".data.userIds+"\r\n";
                })
                .catch(function (error) {
                  console.log(error);
                });
        return false
      },
      refresh(){
        let self = this
        axios.get('/get_user_cont')
                .then(function (response) {
                  let {data} = response.data

                  self.form.at_line = data.count+"\r\n"
                  for (let k in data.list){
                    self.form.at_line += data.list[k]+"\r\n"
                  }
                })
                .catch(function (error) {
                  console.log(error);
                });
      }
    }
  })
</script>
</html>