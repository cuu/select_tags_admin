package menu

import (
	
	//"fmt"
	"time"
	
//	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/cuu/select_tags_admin/models"
//	"github.com/cuu/select_tags_admin/utils"
)


type MenuForm struct {
	SpecifyDate  time.Time `form:"type(date);"`   // Specify the date manually
	BookedSelect models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Dish);attr(multiple,multiple)" valid:""`
  ExtrasSelect  models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Dish);attr(multiple,multiple)" valid:""`
	Booked models.SliceDishField `form:"-"`
	Extras models.SliceDishField `form:"-"`

}

func (form *MenuForm) Valid (v *validation.Validation) {
	
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

