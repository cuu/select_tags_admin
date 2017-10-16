package dish

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	
	. "github.com/cuu/select_tags/controllers"
	"github.com/cuu/select_tags/models"

	"github.com/cuu/select_tags/modules/dish"
	
)

type DishController struct {
	BaseController
	object models.Dish
}


func (this *DishController) Object() interface{} {
	return &this.object
}

func (this *DishController) ObjectQs() orm.QuerySeter {
	return models.Dishes().RelatedSel()
}

func (this *DishController) GetForm() dish.DishForm {
	form := dish.DishForm{}

	form.ListNutritions() 

	return form
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
	form := this.GetForm()
	

	
	this.SetFormSets(&form)
	this.Render()
	
}
