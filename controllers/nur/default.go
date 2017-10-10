package nur

import (
	"github.com/astaxie/beego"
//  "github.com/astaxie/beego/orm"
	. "github.com/cuu/select_tags/controllers"
//rdb "github.com/cuu/select_tags/database"
	
	 "github.com/cuu/select_tags/models"
	"github.com/cuu/select_tags/modules/nur"

	
	"fmt"
)

type NurController struct {
	BaseController
}


func (this *NurController) URLMapping() {
    this.Mapping("GetNur", this.GetNur)
}

// @router /nur [get]
func (this *NurController) GetNur() {

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
func (this *NurController) AddNur() {
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
func (this *NurController) AddNurPost() {

	this.TplName = "nur/add.tpl"

	form := nur.NurForm{}
	if this.ValidFormSets(&form) == false {
		beego.Error("AddNurForm valid failed:")
		this.Render()
		return
	}

	nur := new(models.Nutrition)
	if err := models.SaveNur(nur,form.Name,form.Everyday,form.Indication); err == nil {
	
		this.Redirect("/nur",302)
		return
	}else {
		beego.Error("Add Nur Failed:",err)
	}

	this.Render()
	return
}


// @router /nur/edit/?:id [get]
func (this *NurController) EditNur() {

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
func (this *NurController) EditNurPost() {
	id , _ := this.GetInt(":id")

	fmt.Println("EditNurPost id:",id,this.POST(":id"))
	
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
