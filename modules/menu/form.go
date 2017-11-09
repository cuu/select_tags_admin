package menu

import (
	
	"fmt"
	"time"
	
//	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/cuu/select_tags_admin/models"
	"github.com/cuu/select_tags_admin/utils"
	"github.com/cuu/select_tags_admin/database"
)


type MenuForm struct {
	SpecifyDate   time.Time `form:"type(date);"`   // Specify the date manually
	BookedSelect  models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Dish);attr(multiple,multiple)" valid:""`
  ExtrasSelect  models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Dish);attr(multiple,multiple)" valid:""`
	Booked models.SliceDishField `form:"-"`
	Extras models.SliceDishField `form:"-"`

	MenuMd *models.Menu `form:"-"`
}

func (form *MenuForm) Valid (v *validation.Validation) {
	
}

func (form *MenuForm) BookedSelectSelectData() [][]string {
	data := make([][]string,0)
	if form.MenuMd != nil {
		for _,n := range form.MenuMd.Booked {
			data = append(data,[]string{n.Name,utils.ToStr(n.Id)})
		}
	}
	return data
}


func (form *MenuForm) ExtrasSelectSelectData() [][]string {
	data := make([][]string,0)
	if form.MenuMd != nil {
		for _,n := range form.MenuMd.Extras {
			data = append(data,[]string{n.Name,utils.ToStr(n.Id)})
		}
	}
	return data
}


func (form *MenuForm) SaveMenu(m *models.Menu ) error {

	m.SpecifyDate = form.SpecifyDate
	err := m.Insert()

	if err == nil {
		m.SetBooked(form.Booked)
		m.SetExtras(form.Extras)
	}

	return err
}


func (form *MenuForm) SetFromMenu(m *models.Menu)  {
	utils.SetFormValues(m,form)
	form.BookedSelect = m.Booked.Ids()
	form.ExtrasSelect = m.Extras.Ids()
	
}

func (form *MenuForm) UpdateMenu(m *models.Menu) error {
	changes := utils.FormChanges(m,form)
	utils.SetFormValues(form,m)

	database.StartTrans()
	m.RemoveBooked()
	m.SetBooked(form.Booked)

	m.RemoveExtras()
	m.SetExtras(form.Extras)

	changes = append(changes,"Updated")

	err := m.Update(changes...)
	if err == nil {
		fmt.Println("UpdateMenu start commit")
		database.Commit()
	}else {
		database.Rollback()
	}
	return err
}
