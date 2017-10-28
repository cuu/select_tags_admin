package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"],
		beego.ControllerComments{
			Method: "Dish",
			Router: `/dish`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"],
		beego.ControllerComments{
			Method: "DishAdd",
			Router: `/dish/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"],
		beego.ControllerComments{
			Method: "DishAddPost",
			Router: `/dish/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"],
		beego.ControllerComments{
			Method: "DishDelete",
			Router: `/dish/delete/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"],
		beego.ControllerComments{
			Method: "DishEdit",
			Router: `/dish/edit/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags_admin/controllers/dish:DishController"],
		beego.ControllerComments{
			Method: "DishEditPost",
			Router: `/dish/edit/?:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
