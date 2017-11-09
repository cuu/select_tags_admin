package models


import (
	"time"
	"fmt"
	"strings"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cuu/select_tags_admin/utils"
)

// a dish like a dish of food
type Dish struct {
	Id int
	Name string `orm:"size(255);unique"`
	//	Nutritions []*Nutrition `orm:"rel(m2m)"`
	//Nutritions  SliceNutritionPointers `orm:"rel(m2m);on_delete(set_null)"`
	FirstClass  string
	SecClass    string
	ThirdClass  string
	
	EstimatePrice int
	
	Ingredients SliceIngredientPointers `orm:"rel(m2m)"`
	Booked SliceMenuPointers `orm:"reverse(many);rel_table(menu_booked)"`
	Extras SliceMenuPointers `orm:"reverse(many);rel_table(menu_extras)"`
	Images SliceImagePointers `orm:"rel(m2m);rel_table(dish_img)"`

	Image1 string
	Image2 string
	Image3 string
	Image4 string
	Image5 string
	Image6 string
	
	Created  time.Time  `orm:"auto_now_add"`
	Updated  time.Time  `orm:"auto_now"`
}


func (m *Dish) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err != nil {
		return err
	}
	return nil
}

func (m *Dish) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil {
		return err
	}
	return nil
}

func (m *Dish) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...); err != nil {
		return err
	}
	return nil
}

func (m *Dish) Delete() error {
	if _,err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Dish) RemoveIngredients() (int64,error) {
	m2m := orm.NewOrm().QueryM2M(m,"Ingredients")

	if len(m.Ingredients) > 0 {
		num,err := m2m.Remove(m.Ingredients)
		return num,err
	}else {
		return 0,nil
	}
	
}

func (m *Dish) LoadIngredients() (int64,error) {
	num,err := orm.NewOrm().LoadRelated(m,"Ingredients")
	fmt.Println("LoadIngredients: ",num)
	return num,err
}

func (m *Dish) SetIngredients( arr []Ingredient ) {
	if m.Id == 0 {
		beego.Error("We should Insert before QueryM2M Add")
		return
	}
	m2m := orm.NewOrm().QueryM2M(m,"Ingredients")

	for i,_ := range arr {
		m2m.Add(arr[i])
	}
}


func Dishes() orm.QuerySeter {
	return orm.NewOrm().QueryTable("dish").OrderBy("-Id")
}


type SliceDishPointers []*Dish

func (e SliceDishPointers) Label() []string {
	var d []string
	for _,p := range e {
		d = append(d,p.Name)
	}
	return d
}

func (e SliceDishPointers) Ids() []string {
	var d []string
	for _,p := range e {
		d = append(d,utils.ToStr(p.Id))
		
	}

	return d
}

func (e *SliceDishPointers) String() string {
	return strings.Join(e.Label(),",")
}


func init(){
	orm.RegisterModel(new(Dish))
}
