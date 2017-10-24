package menu

import (
	
	//"fmt"
	"time"
	
//	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/cuu/select_tags/models"
//	"github.com/cuu/select_tags/utils"
)


type MenuForm struct {
	SpecifyDate  time.Time `form:"type(date);"`   // Specify the date manually
	BookedSelect models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Dish);attr(multiple,multiple)" valid:""`
  ExtraSelect  models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Dish);attr(multiple,multiple)" valid:""`
	Dishes models.SliceDishField `form:"-"`

}

func (form *MenuForm) Valid (v *validation.Validation) {
	
}

