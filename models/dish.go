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



func Dishes() orm.QuerySeter {
	return orm.NewOrm().QueryTable("dish").OrderBy("-Id")
}


func init(){
	orm.RegisterModel(new(Dish))
}
