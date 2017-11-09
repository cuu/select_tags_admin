package menu


import (
//	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	
	. "github.com/cuu/select_tags_admin/controllers"
	"github.com/cuu/select_tags_admin/models"

	"github.com/cuu/select_tags_admin/modules/menu"
	"github.com/cuu/select_tags_admin/database"
	
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

// @router /menu/edit/?:id [get]
func (this *MenuController) MenuEdit() {
	this.TplName = "menu/edit.tpl"
	id,_ := this.GetInt(":id")
	var menuMd = new(models.Menu)
	if id > 0 {
		qs := models.Menus().Filter("Id",id)
		qs.RelatedSel(1).One(menuMd)
	}

	menuMd.LoadBooked()
	menuMd.LoadExtras()

	if menuMd.Id == 0 {
		this.Abort("404")
		return
	}

	form := menu.MenuForm{MenuMd:menuMd}

	form.SetFromMenu(menuMd)

	this.SetFormSets(&form)
	this.Data["Id"] = id
	this.Render()
	return
	
	
}

// @router /menu/edit/?:id [post]
func (this *MenuController) MenuEditPost() {
	this.TplName = "menu/edit.tpl"
	id,_ := this.GetInt(":id")
	var menuMd = new(models.Menu)
	if id > 0 {
		qs := models.Menus().Filter("Id",id)
		qs.RelatedSel(1).One(menuMd)
	}else {
		beego.Error("Id <= 0,error! Menu Edit Post")
	}

	menuMd.LoadBooked()
	menuMd.LoadExtras()

	if menuMd.Id == 0 {
		this.Abort("404")
		return
	}

	form := menu.MenuForm{MenuMd:menuMd}
	ids := this.GetStrings("BookedSelect")
	form.Booked.Set(ids)
	ids = this.GetStrings("ExtrasSelect")
	form.Extras.Set(ids)

	form.SetFromMenu(menuMd)
	this.Data["Id"] = id
	if !this.ValidFormSets(&form) {
		beego.Error("Upate Menu Post Valid Error")
		this.Render()
		return
	}

	if err := form.UpdateMenu(menuMd); err == nil {
		this.JsStorage("deleteKey","menu/edit")
		this.Redirect("/menu",302)
	}else {
		beego.Error("UpdateMenu falied: ",err)
		this.Render()
	}
	return
	
}

// @router /menu/delete/?:id [get]
func (this *MenuController) MenuDelete() {
	id,_ := this.GetInt(":id")
	var menuMd models.Menu
	if id > 0 {
		qs := models.Menus().Filter("Id",id)
		qs.RelatedSel(1).One(&menuMd)
	}else {
		beego.Error("MenuDelete Id illegal")
		return
	}

	if menuMd.Id == 0 {
		this.Abort("502")
		return
	}

	menuMd.LoadExtras()
	menuMd.LoadBooked()

	database.StartTrans()
	menuMd.RemoveBooked()
	menuMd.RemoveExtras()


	if err := menuMd.Delete(); err == nil {
		database.Commit()
		this.FlashRedirect("/menu",302,"DeleteSuccess")
		return
	} else {
		database.Rollback()
		beego.Error(err)
		this.Data["Error"] = err
		this.Render()
	}
}
