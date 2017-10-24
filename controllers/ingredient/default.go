package ingredient

import (
//	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	
	. "github.com/cuu/select_tags/controllers"
	"github.com/cuu/select_tags/models"

	"github.com/cuu/select_tags/modules/ingredient"
	"github.com/cuu/select_tags/database"
	
)

type IngredientController struct {
	BaseController
	object models.Ingredient
}


func (this *IngredientController) Object() interface{} {
	return &this.object
}

func (this *IngredientController) ObjectQs() orm.QuerySeter {
	return models.Ingredients().RelatedSel()
}

func (this *IngredientController) GetForm() ingredient.IngredientForm {
	form := ingredient.IngredientForm{}

	form.ListNutritions() 

	return form
}

// @router /ingredient [get]
func (this *IngredientController) Ingredient() {

	var ingredients []*models.Ingredient

	this.Data["Title"] = "Ingredient"

	qs := models.Ingredients().OrderBy("-Created").Limit(25).RelatedSel()

	models.ListObjects(qs,&ingredients)

	for _,p := range ingredients {
		p.LoadNutritions()
	}
		
	this.Data["Count"] = len(ingredients)
	this.Data["Ingredients"] = ingredients

	this.TplName ="ingredient/index.tpl"
	this.Render()
}

// @router /ingredient/add [get]
func (this *IngredientController) IngredientAdd() {


	this.TplName = "ingredient/add.tpl"
	form := this.GetForm()
	

	
	this.SetFormSets(&form)
	this.Render()
	
}

// @router /ingredient/add [post]
func (this *IngredientController) IngredientAddPost() {
	this.TplName = "ingredient/add.tpl"
	form := ingredient.IngredientForm{}
	
	ids := this.GetStrings("NutritionsSelect")
	form.Nutritions.Set(ids)
	
	if this.ValidFormSets(&form) == false {
		beego.Error("IngredientAdd Form valid failed: ")
		this.Render()
		return
	}
	
	ingredientMd := new(models.Ingredient)
	if err := form.SaveIngredient(ingredientMd); err == nil {
		this.Redirect("/ingredient",302)
		return
	}else {
		beego.Error("Ingredient Add Failed: ",err)
	}
	
}


// @router /ingredient/delete/?:id [get]
func (this *IngredientController) IngredientDelete() {
	id, _ := this.GetInt(":id")
	var ingredientMd models.Ingredient
	if id > 0 {
		qs := models.Ingredients().Filter("Id",id)
		qs.RelatedSel(1).One(&ingredientMd)
	}


	if ingredientMd.Id == 0 {
		this.Abort("502")
		return
	}
	ingredientMd.LoadNutritions()

	
	database.StartTrans()
	
	ingredientMd.RemoveNutritions()
	
	if err := ingredientMd.Delete(); err == nil {

		database.Commit()
		this.FlashRedirect("/ingredient",302,"DeleteSuccess")
		return
	} else {
		database.Rollback()
		beego.Error(err)
		this.Data["Error"] = err
		this.Render()
	}
}
