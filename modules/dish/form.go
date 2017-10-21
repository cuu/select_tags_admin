package dish

import (

	//"fmt"

//	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/cuu/select_tags/models"
	"github.com/cuu/select_tags/utils"
)

type DishForm struct {
	Name string `valid:"Required;MinSize(2)"`
//	Test1 []int `form:"type(select);attr(rel,select2-admin-model);attr(data-model,Nutrition)" valid:""`
//	Test2 int `form:type(select);attr(rel,select2)" valid:""`
	//Nurs []models.Nutrition `form:"type(select);attr(rel,select2);attr(multiple,multiple)" valid:""`

	NursSelect models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Nutrition);attr(multiple,multiple)" valid:""`

	Nurs models.SliceNutritionField `form:"-"`
}



func (form *DishForm) ListNutritions() (int64, error) {
	return models.Nutritions().All(&form.Nurs)
}

func (form *DishForm) Valid( v*validation.Validation) {
	
}


/*
func (form *DishForm) Nurs2SelectData() [][]string {
	data := make([][]string,0,len(form.Nurs))

	for _, n := range form.Nurs {
		data = append(data,[]string{n.Name,utils.ToStr(n.Id)} )
	}
	
	return data
}


func (form *DishForm) NursSelectData() [][]string {
	data := make([][]string,0,len(form.Nurs))

	for _, n := range form.Nurs {
		data = append(data,[]string{n.Name,utils.ToStr(n.Id)} )
	}
	
	return data
}
*/

func (form *DishForm) SetFromDish(m *models.Dish ) {
	utils.SetFormValues(m,form)
}

func (form *DishForm) SaveDish(m *models.Dish) error {

	m.Name = form.Name
	
	err := m.Insert()
	if err == nil {
		/// We need the inserted id for QueryM2M add
		m.SetNutritions(form.Nurs)
	}
	
	return err
}

