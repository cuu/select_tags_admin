package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/menu:MenuController"],
		beego.ControllerComments{
			Method: "Menu",
			Router: `/menu`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/menu:MenuController"],
		beego.ControllerComments{
			Method: "MenuAdd",
			Router: `/menu/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}