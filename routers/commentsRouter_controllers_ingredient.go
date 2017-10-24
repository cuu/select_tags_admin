package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/ingredient:IngredientController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/ingredient:IngredientController"],
		beego.ControllerComments{
			Method: "Ingredient",
			Router: `/ingredient`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/ingredient:IngredientController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/ingredient:IngredientController"],
		beego.ControllerComments{
			Method: "IngredientAdd",
			Router: `/ingredient/add`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/ingredient:IngredientController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/ingredient:IngredientController"],
		beego.ControllerComments{
			Method: "IngredientAddPost",
			Router: `/ingredient/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/ingredient:IngredientController"] = append(beego.GlobalControllerRouter["github.com/cuu/select_tags/controllers/ingredient:IngredientController"],
		beego.ControllerComments{
			Method: "IngredientDelete",
			Router: `/ingredient/delete/?:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
