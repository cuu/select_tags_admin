package menu


import (
//	"fmt"
//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	
	. "github.com/cuu/select_tags/controllers"
	"github.com/cuu/select_tags/models"

	"github.com/cuu/select_tags/modules/menu"
//	"github.com/cuu/select_tags/database"
	
)


type MenuController struct {
	BaseController
	object models.Menu
}


func (this *MenuController) Object() interface{} {
	return &this.object
}

func (this *MenuController) ObjectQs() orm.QuerySeter {
	return models.Dishes().RelatedSel()
}

func (this *MenuController) GetForm() menu.MenuForm {
	form := menu.MenuForm{}
	return form
}

// @router /menu [get]
func (this *MenuController) Menu() {
	var menus []*models.Menu
	this.Data["Title"]  = "Menus"

	qs := models.Menus().OrderBy("-SpecifyDate").Limit(25).RelatedSel()

	models.ListObjects(qs,&menus)

	this.Data["Count"] = len(menus)
	this.Data["Menus"] = menus

	this.TplName ="menu/index.tpl"
	this.Render()
}

// @router /menu/add [get]
func (this *MenuController) MenuAdd() {
	this.TplName = "menu/add.tpl"
	form := this.GetForm()

	this.SetFormSets(&form)
	this.Render()

}

