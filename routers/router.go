package routers

import (
	"github.com/cuu/select_tags_admin/controllers"
	"github.com/cuu/select_tags_admin/controllers/nur"
	"github.com/cuu/select_tags_admin/controllers/dish"
	"github.com/cuu/select_tags_admin/controllers/ingredient"
	"github.com/cuu/select_tags_admin/controllers/menu"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include( &controllers.MainController{})

	beego.Include( &controllers.ModelsGetSearchController{})
	beego.Include( &controllers.ServerPHPController{})
	beego.Include( &controllers.PtsController{})
	
	beego.Include( &nur.NurController{} ) //quote router
	beego.Include( &dish.DishController{} )

	beego.Include( &ingredient.IngredientController{} )

	beego.Include( &menu.MenuController{} )
	
}
