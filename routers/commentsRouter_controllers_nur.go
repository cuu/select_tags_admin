package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "GetNur",
			Router: `/nur`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "AddNur",
			Router: `/nur/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "AddNurPost",
			Router: `/nur/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "EditNur",
			Router: `/nur/edit/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/nur:NurController"],
		beego.ControllerComments{
			Method: "EditNurPost",
			Router: `/nur/edit/?:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
