package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"],
		beego.ControllerComments{
			Method: "Menu",
			Router: `/menu`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"],
		beego.ControllerComments{
			Method: "MenuAdd",
			Router: `/menu/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"],
		beego.ControllerComments{
			Method: "MenuAddPost",
			Router: `/menu/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"],
		beego.ControllerComments{
			Method: "MenuDelete",
			Router: `/menu/delete/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"],
		beego.ControllerComments{
			Method: "MenuEdit",
			Router: `/menu/edit/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/menu:MenuController"],
		beego.ControllerComments{
			Method: "MenuEditPost",
			Router: `/menu/edit/?:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
