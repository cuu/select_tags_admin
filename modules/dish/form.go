package dish

import (

	//"fmt"
//	"time"
	
//	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/cuu/select_tags/models"
	"github.com/cuu/select_tags/utils"
)

type DishForm struct {
//	TheDate  time.Time `form:"type(date);"`   // Specify the date manually
	Name string `valid:"Required;MinSize(2)"`
	
	EstimatePrice int
	
	IngredientsSelect models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Ingredient);attr(multiple,multiple)" valid:""`

	Ingredients models.SliceIngredientField `form:"-"`
}



func (form *DishForm) ListIngredients() (int64, error) {
	return models.Ingredients().Limit(25).All(&form.Ingredients)
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
		m.SetIngredients(form.Ingredients)
	}
	
	return err
}

