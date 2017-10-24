package dish

import (
//	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	
	. "github.com/cuu/select_tags/controllers"
	"github.com/cuu/select_tags/models"

	"github.com/cuu/select_tags/modules/dish"
	"github.com/cuu/select_tags/database"
	
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

	form.ListIngredients() 

	return form
}

// @router /dish [get]
func (this *DishController) Dish() {

	var dishes []*models.Dish

	this.Data["Title"] = "Dish"

	qs := models.Dishes().OrderBy("-Created").Limit(25).RelatedSel()

	models.ListObjects(qs,&dishes)

	for _,p := range dishes {
		p.LoadIngredients()
	}
		
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

// @router /dish/add [post]
func (this *DishController) DishAddPost() {
	this.TplName = "dish/add.tpl"
	form := dish.DishForm{}
	
	ids := this.GetStrings("IngredientsSelect")
	form.Ingredients.Set(ids)
	
	if this.ValidFormSets(&form) == false {
		beego.Error("DishAdd Form valid failed: ")
		this.Render()
		return
	}
	
	dishMd := new(models.Dish)
	if err := form.SaveDish(dishMd); err == nil {
		this.Redirect("/dish",302)
		return
	}else {
		beego.Error("Dish Add Failed: ",err)
	}
	
}


// @router /dish/delete/?:id [get]
func (this *DishController) DishDelete() {
	id, _ := this.GetInt(":id")
	var dishMd models.Dish
	if id > 0 {
		qs := models.Dishes().Filter("Id",id)
		qs.RelatedSel(1).One(&dishMd)
	}


	if dishMd.Id == 0 {
		this.Abort("502")
		return
	}
	dishMd.LoadIngredients()

	
	database.StartTrans()
	
	dishMd.RemoveIngredients()
	
	if err := dishMd.Delete(); err == nil {

		database.Commit()
		this.FlashRedirect("/dish",302,"DeleteSuccess")
		return
	} else {
		database.Rollback()
		beego.Error(err)
		this.Data["Error"] = err
		this.Render()
	}
}
