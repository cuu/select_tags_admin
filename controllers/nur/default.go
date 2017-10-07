package nur

import (
//	"github.com/astaxie/beego"
//  "github.com/astaxie/beego/orm"
	. "github.com/cuu/select_tags/controllers"
	rdb "github.com/cuu/select_tags/database"
	
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
	o := rdb.NewOrm()
	nur := new(models.Nutrition)
	
	this.Data["Title"] = "Nur"
	cnt,err := o.QueryTable(nur).Count()
	
	fmt.Println("Count Nurition: %s, %s",cnt,err)

	this.Data["Count"] = cnt
	
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

	form := nur.AddNurForm{}
	this.SetFormSets(&form)
	
	this.Render()
}

// @router /nur/add [post]
func (this *NurController) AddNurPost() {

	this.TplName = "nur/add.tpl"
	
	this.Redirect("/",302)
	return
}
