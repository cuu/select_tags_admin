package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/dish:DishController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/dish:DishController"],
		beego.ControllerComments{
			Method: "Dish",
			Router: `/dish`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/dish:DishController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/dish:DishController"],
		beego.ControllerComments{
			Method: "DishAdd",
			Router: `/dish/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
