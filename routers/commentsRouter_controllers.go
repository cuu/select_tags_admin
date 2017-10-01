package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers:NurController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers:NurController"],
		beego.ControllerComments{
			Method: "GetNur",
			Router: `/nur`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
