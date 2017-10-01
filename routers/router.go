package routers

import (
	"github.com/cuu/select_tags/controllers"
	
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include( &controllers.NurController{} ) //quote router
}
