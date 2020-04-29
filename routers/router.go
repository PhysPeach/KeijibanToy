package routers

import (
	"KeijibanToy/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/post", &controllers.MainController{}, "post:Post")
}
