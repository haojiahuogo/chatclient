package routers

import (
	"chatclient/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ClientController{})
	beego.Router("/login", &controllers.ClientController{}, "post:Login")

}
