package dish

import (
	"fmt"

	. "github.com/cuu/select_tags/controllers"
	"github.com/cuu/select_tags/models"

	"github.com/cuu/select_tags/modules/dish"
	
)

type DishController struct {
	BaseController
}

// @router /dish [get]
func (this *DishController) Dish() {

	var dishes []models.Dish

	this.Data["Title"] = "Dish"

	qs := models.Dishes().OrderBy("-Created").Limit(25).RelatedSel()

	models.ListObjects(qs,&dishes)
	fmt.Println(dishes)

	this.Data["Count"] = len(dishes)
	this.Data["Dishes"] = dishes

	this.TplName ="dish/index.tpl"
	this.Render()
}

// @router /dish/add [get]
func (this *DishController) DishAdd() {


	this.TplName = "dish/add.tpl"
	form := dish.DishForm{}
	this.SetFormSets(&form)
	this.Render()
	
}
