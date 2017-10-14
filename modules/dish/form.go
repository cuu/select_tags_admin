package dish

import (

	"github.com/astaxie/beego/validation"
	"github.com/cuu/select_tags/models"
//	"github.com/cuu/select_tags/utils"
)

type DishForm struct {
	Name string `valid:"Required;MinSize(2)"`
	Nurs models.SliceNutritionField 
}


func (form *DishForm) Valid( v*validation.Validation) {
	
}

