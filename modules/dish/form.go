package dish

import (

	"fmt"
//	"strings"
//	"time"
	
//	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/cuu/select_tags_admin/models"
	"github.com/cuu/select_tags_admin/utils"
	"github.com/cuu/select_tags_admin/database"
	
)

type DishForm struct {
//	TheDate  time.Time `form:"type(date);"`   // Specify the date manually
	Name string `valid:"Required;MinSize(2)"`
	
	EstimatePrice int

	FirstClass string `form:"type(select)"`
	SecClass string   `form:"type(select)"`
	ThirdClass string `form:"type(select)"`
	
	IngredientsSelect models.SliceStringField `form:"type(select);class(select2-admin-model);attr(data-model,Ingredient);attr(multiple,multiple)" valid:""`

	Ingredients models.SliceIngredientField `form:"-"`

	Image1 string  `form:"attr(readonly,readonly)" valid:""`
	Image2 string  `form:"attr(readonly,readonly)" valid:""`
	Image3 string  `form:"attr(readonly,readonly)" valid:""`
	Image4 string  `form:"attr(readonly,readonly)" valid:""`
	Image5 string  `form:"attr(readonly,readonly)" valid:""`
	Image6 string  `form:"attr(readonly,readonly)" valid:""`
	
	DishMd  *models.Dish `form:"-"`
}


func (form *DishForm) ListIngredients() (int64, error) {
	return models.Ingredients().Limit(25).All(&form.Ingredients)
}

func (form *DishForm) Valid( v*validation.Validation) {
	
}


func (form *DishForm) FirstClassSelectData() [][]string {
	data := make([][]string,0)
	data = append(data,[]string{"未分类","0"})
	data = append(data,[]string{"凉菜", "1"})
	data = append(data,[]string{"荤菜", "2"})
	data = append(data,[]string{"素菜", "3"})
	data = append(data,[]string{"汤",   "4"})
	return data

}

func (form *DishForm) IngredientsSelectSelectData() [][]string {
	data := make([][]string,0)
	if form.DishMd != nil {
		for _,n := range form.DishMd.Ingredients {
			data = append(data,[]string{n.Name,utils.ToStr(n.Id)} )
		}
	}
	return data
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
	form.IngredientsSelect = m.Ingredients.Ids() // to make them selected="selected"
}

func (form *DishForm) SaveDish(m *models.Dish) error {

	m.Name = form.Name
	m.FirstClass = form.FirstClass
	m.SecClass   = form.SecClass
	m.ThirdClass = form.ThirdClass
	
	err := m.Insert()
	if err == nil {
		/// We need the inserted id for QueryM2M add
		m.SetIngredients(form.Ingredients)
	}
	
	return err
}


func (form *DishForm) UpdateDish(m *models.Dish) error {
	changes := utils.FormChanges(m,form)

	utils.SetFormValues(form,m)

	database.StartTrans()
	m.RemoveIngredients()
	m.SetIngredients(form.Ingredients)
	
	changes = append(changes,"Updated")
	err := m.Update(changes...)
	if err == nil {
		fmt.Println("now commit")
		database.Commit()
	}else{
		database.Rollback()
	}
	return err
	
}
