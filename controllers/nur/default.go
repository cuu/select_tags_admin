package nur

import (
	"fmt"

	
	"github.com/astaxie/beego"
//  "github.com/astaxie/beego/orm"
	. "github.com/cuu/select_tags_admin/controllers"
//rdb "github.com/cuu/select_tags_admin/database"
	
	"github.com/cuu/select_tags_admin/models"
	"github.com/cuu/select_tags_admin/modules/nur"

	
	
)

type NurController struct {
	BaseController
}



// @router /nur [get]
func (this *NurController) Nur() {

	var nurs []models.Nutrition
	
	this.Data["Title"] = "Nur"
	
	qs := models.Nutritions().OrderBy("-Created").Limit(25).RelatedSel()

	models.ListObjects(qs,&nurs)

	fmt.Println(nurs)
	
	this.Data["Count"] = len(nurs)
	this.Data["Nurs"]  = nurs
	
	this.TplName = "nur/index.tpl"
	
	this.Render()
}

// @router /nur/add [get]
func (this *NurController) NurAdd() {
	/*
	o := rdb.NewOrm()
	nur := new(models.Nutrition)
	*/

	
	this.TplName = "nur/add.tpl"

	form := nur.NurForm{}
	this.SetFormSets(&form)
	
	this.Render()
}

// @router /nur/add [post]
func (this *NurController) NurAddPost() {

	this.TplName = "nur/add.tpl"

	form := nur.NurForm{}
	if this.ValidFormSets(&form) == false {
		beego.Error("NurAddForm valid failed:")
		this.Render()
		return
	}

	nurMd := new(models.Nutrition)
	if err := form.SaveNutrition(nurMd); err == nil {
		this.Redirect("/nur",302)
		return
	}else {
		beego.Error("Nur Add Failed:",err)
	}

	this.Render()
	return
}


// @router /nur/edit/?:id [get]
func (this *NurController) NurEdit() {

	this.TplName = "nur/edit.tpl"
	id , _ := this.GetInt(":id")

	var  nurMd models.Nutrition

	// If I had bind data to user, I'll filter Nutritions with User.Id
	if id > 0 {
		qs := models.Nutritions().Filter("Id",id)

		qs.RelatedSel(1).One(&nurMd)
		
	}
	
	if nurMd.Id == 0 {
		this.Abort("404")
		return
	}

	form := nur.NurForm{}
	form.SetFromNutrition(&nurMd)
	
	
	this.SetFormSets(&form)// generate html code
	this.Data["Id"] = id  // edit.tpl needs it 
	this.Render()
	return
}

// @router /nur/edit/?:id [post]
func (this *NurController) NurEditPost() {
	id , _ := this.GetInt(":id")
	
	var  nurMd models.Nutrition

		// If I had bind data to user, I'll filter Nutritions with User.Id
	if id > 0 {
		qs := models.Nutritions().Filter("Id",id)

		qs.RelatedSel(1).One(&nurMd)
		
	}
	
	if nurMd.Id == 0 {
		this.Abort("404")
		return
	}

	
	form := nur.NurForm{}
	form.SetFromNutrition(&nurMd)
	if !this.ValidFormSets(&form) {
		return
	}

	if err := form.UpdateNutrition(&nurMd); err == nil {
		this.JsStorage("deleteKey","nur/edit") // Set in Cookie
		this.Redirect("/nur",302)
	}else {
		beego.Error("UpdateNutrition failed: ",err)
	}
	return
}

// @router /nur/delete/?:id [get]
func (this *NurController) NurDelete() {
	id , _ := this.GetInt(":id")
	
	var  nurMd models.Nutrition

		// If I had bind data to user, I'll filter Nutritions with User.Id
	if id > 0 {
		qs := models.Nutritions().Filter("Id",id)

		qs.RelatedSel(1).One(&nurMd)
		
	}
	
	if nurMd.Id == 0 {
		this.Abort("502")
		return
	}

	if err := nurMd.Delete(); err == nil {
		this.FlashRedirect("/nur",302,"DeleteSuccess")
		return
	}else {
		beego.Error(err)
		this.Data["Error"] = err
	}
}

