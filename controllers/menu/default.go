package menu


import (
//	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	
	. "github.com/cuu/select_tags_admin/controllers"
	"github.com/cuu/select_tags_admin/models"

	"github.com/cuu/select_tags_admin/modules/menu"
//	"github.com/cuu/select_tags_admin/database"
	
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

	for _,p := range menus {
		p.LoadExtras()
		p.LoadBooked()
	}
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

// @router /menu/add [post]
func (this *MenuController) MenuAddPost() {
	this.TplName = "menu/add.tpl"
	form := this.GetForm()

	ids := this.GetStrings("BookedSelect")
	form.Booked.Set(ids)
	ids = this.GetStrings("ExtrasSelect")
	form.Extras.Set(ids)

	if this.ValidFormSets(&form) == false {
		beego.Error("MenuAdd Form Valid Failed: ")
		this.Render()
		return
	}

	menuMd := new(models.Menu)
	if err := form.SaveMenu(menuMd); err == nil {
		this.Redirect("/menu",302)
		return
	} else {
		beego.Error("MenuAdd Failed: ",err)
		this.Render()
	}
	
}
