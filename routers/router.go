package routers

import (
	"github.com/cuu/select_tags/controllers"
	"github.com/cuu/select_tags/controllers/nur"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include( &nur.NurController{} ) //quote router
}
