package models


import (
	"time"
	"fmt"
	"strings"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// a dish like a dish of food
type Dish struct {
	Id int
	Name string `orm:"size(255);unique"`
	//	Nutritions []*Nutrition `orm:"rel(m2m)"`
	Nutritions  SliceNutritionPointers `orm:"rel(m2m);on_delete(set_null)"`
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

func (m *Dish) RemoveNutritions() (int64,error) {
	m2m := orm.NewOrm().QueryM2M(m,"Nutritions")

	num,err := m2m.Remove(m.Nutritions)
	return num,err
}

func (m *Dish) LoadNutritions() (int64,error) {
	num,err := orm.NewOrm().LoadRelated(m,"Nutritions")
	fmt.Println("LoadNutritions: ",num)
	return num,err
}

func (m *Dish) SetNutritions( nurs []Nutrition) {
	if m.Id == 0 {
		beego.Error("We should Insert before QueryM2M Add")
		return
	}
	m2m := orm.NewOrm().QueryM2M(m,"Nutritions")

	for i,_ := range nurs {
		m2m.Add(nurs[i])
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

func (e *SliceDishPointers) String() string {
	return strings.Join(e.Label(),",")
}


func init(){
	orm.RegisterModel(new(Dish))
}
