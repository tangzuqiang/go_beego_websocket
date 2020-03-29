### beego 框架 websocket 推送给全部或指定用户
### 1安装 bee
### 2 执行 go mod tidy
### 3 执行 go mod vendor
### 会自动下载相关的依赖包
###### 目录结构
```
├─conf       //配置文件
├─controllers // 控制器文件
├─models // 模型文件用于操作数据库
├─routers // 路由文件
├─service // websocket服务
├─static  // 前端静态文件
│  ├─css
│  ├─img
│  └─js
├─tests  //测试
├─util  // 自定义工具类
│  ├─data //数据格式方法
│  └─server // websocket 管理
├─vendor // 插件文件夹
│  ├─github.com
│  │  ├─astaxie
│  │  │  └─beego
│  │  │      ├─config
│  │  │      ├─context
│  │  │      │  └─param
│  │  │      ├─grace
│  │  │      ├─logs
│  │  │      ├─session
│  │  │      ├─toolbox
│  │  │      └─utils
│  │  ├─go-redis
│  │  │  └─redis
│  │  │      └─internal
│  │  │          ├─consistenthash
│  │  │          ├─hashtag
│  │  │          ├─pool
│  │  │          ├─proto
│  │  │          ├─singleflight
│  │  │          └─util
│  │  ├─gopherjs
│  │  │  └─gopherjs
│  │  │      └─js
│  │  ├─gorilla
│  │  │  └─websocket
│  │  ├─jtolds
│  │  │  └─gls
│  │  ├─konsorten
│  │  │  └─go-windows-terminal-sequences
│  │  ├─satori
│  │  │  └─go.uuid
│  │  ├─shiena
│  │  │  └─ansicolor
│  │  ├─sirupsen
│  │  │  └─logrus
│  │  └─smartystreets
│  │      ├─assertions
│  │      │  └─internal
│  │      │      ├─go-render
│  │      │      │  └─render
│  │      │      └─oglematchers
│  │      └─goconvey
│  │          └─convey
│  │              ├─gotest
│  │              └─reporting
│  ├─golang.org
│  │  └─x
│  │      ├─crypto
│  │      │  └─acme
│  │      │      └─autocert
│  │      ├─net
│  │      │  └─idna
│  │      ├─sys
│  │      │  └─unix
│  │      └─text
│  │          ├─secure
│  │          │  └─bidirule
│  │          ├─transform
│  │          └─unicode
│  │              ├─bidi
│  │              └─norm
│  └─gopkg.in
│      └─yaml.v2
└─views // 视图文件
```
```
里面去除了token redis mysql之类的操作,方便demo的部署
```
###运行效果图片
![run_image](https://tang-zuqiang.oss-cn-shenzhen.aliyuncs.com/img/websocket.png)
####最后
没有花太多的时间弄，应该还有很多问题,同时希望认识一些热爱技术的朋友一起撸代码
可以加我微信
####扫码加我微信,一起撸代码,或者有项目可以一起做
![Image weixin](https://tang-zuqiang.oss-cn-shenzhen.aliyuncs.com/img/mmqrcode1585448410857.png)

