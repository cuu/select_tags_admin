package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers:MainController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers:MainController"],
		beego.ControllerComments{
			Method: "IndexPost",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
