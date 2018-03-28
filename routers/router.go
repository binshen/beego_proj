package routers

import (
	"beego_proj/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/account/", &controllers.AccountController{})
}
