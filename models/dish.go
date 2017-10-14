package models


import (
	"time"
	"github.com/astaxie/beego/orm"
)


// a dish like a dish of food
type Dish struct {
	Id int
	Name string `orm:"size(255);unique"`
	Nurs []*Nutrition `orm:"rel(m2m)"`
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

		
func Dishes() orm.QuerySeter {
	return orm.NewOrm().QueryTable("dish").OrderBy("-Id")
}


func init(){
	orm.RegisterModel(new(Dish))
}
