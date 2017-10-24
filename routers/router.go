package routers

import (
	"github.com/cuu/select_tags/controllers"
	"github.com/cuu/select_tags/controllers/nur"
	"github.com/cuu/select_tags/controllers/dish"
	"github.com/cuu/select_tags/controllers/ingredient"
	"github.com/cuu/select_tags/controllers/menu"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include( &controllers.MainController{})

	beego.Include( &controllers.ModelsGetSearchController{})
	
	beego.Include( &nur.NurController{} ) //quote router
	beego.Include( &dish.DishController{} )

	beego.Include( &ingredient.IngredientController{} )

	beego.Include( &menu.MenuController{} )
	
}
