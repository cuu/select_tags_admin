package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "Nur",
			Router: `/nur`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "NurAdd",
			Router: `/nur/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "NurAddPost",
			Router: `/nur/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "NurDelete",
			Router: `/nur/delete/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "NurEdit",
			Router: `/nur/edit/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "NurEditPost",
			Router: `/nur/edit/?:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
