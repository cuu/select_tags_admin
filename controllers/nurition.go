package controllers

import (
//	"github.com/astaxie/beego"
//  "github.com/astaxie/beego/orm"
	rdb "github.com/cuu/select_tags/database"
	 "github.com/cuu/select_tags/models"

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
	
	this.TplName = "nurition.tpl"
	
	this.Render()
}
