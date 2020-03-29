package routers

import (
	"github.com/astaxie/beego"
	"webapi/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/get_user_cont", &controllers.PushController{}, "*:Getatlinenum")
	beego.Router("/push", &controllers.PushController{}, "*:Push")
}
