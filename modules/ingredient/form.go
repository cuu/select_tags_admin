package ingredient

import (

	//"fmt"

//	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/cuu/select_tags/models"
	"github.com/cuu/select_tags/utils"
)

	type IngredientForm struct {
	Name string `valid:"Required;MinSize(1)"`
//	Test1 []int `form:"type(select);attr(rel,select2-admin-model);attr(data-model,Nutrition)" valid:""`
//	Test2 int `form:type(select);attr(rel,select2)" valid:""`
	//Nurs []models.Nutrition `form:"type(select);attr(rel,select2);attr(multiple,multiple)" valid:""`

//	EstimatePrice int
	
	NutritionsSelect models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Nutrition);attr(multiple,multiple)" valid:""`

	Nutritions models.SliceNutritionField `form:"-"`
}



func (form *IngredientForm) ListNutritions() (int64, error) {
	return models.Nutritions().Limit(25).All(&form.Nutritions)
}

func (form *IngredientForm) Valid( v*validation.Validation) {
	
}


/*
func (form *IngredientForm) Nurs2SelectData() [][]string {
	data := make([][]string,0,len(form.Nurs))

	for _, n := range form.Nurs {
		data = append(data,[]string{n.Name,utils.ToStr(n.Id)} )
	}
	
	return data
}


func (form *IngredientForm) NursSelectData() [][]string {
	data := make([][]string,0,len(form.Nurs))

	for _, n := range form.Nurs {
		data = append(data,[]string{n.Name,utils.ToStr(n.Id)} )
	}
	
	return data
}
*/

func (form *IngredientForm) SetFromIngredient(m *models.Ingredient ) {
	utils.SetFormValues(m,form)
}

func (form *IngredientForm) SaveIngredient(m *models.Ingredient) error {

	m.Name = form.Name
	
	err := m.Insert()
	if err == nil {
		/// We need the inserted id for QueryM2M add
		m.SetNutritions(form.Nutritions)
	}
	
	return err
}

