package models

import (
	"time"
//	"strconv"
	"github.com/astaxie/beego/orm"
)

type Nutrition struct {
  Id   int
	Name string `orm:"size(255);unique"`
	Everyday int ``
	Indication string `orm:"size(1024)"`
	Created  time.Time  `orm:"auto_now_add"`
	Updated  time.Time  `orm:"auto_now"`
}



func (m *Nutrition) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Nutrition) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil {
		return err
	}
	return nil
}

func (m *Nutrition) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...); err != nil {
		return err
	}
	return nil
}

func (m *Nutrition) Delete() error {
	if _,err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}


func (m *Nutrition) String() string {
	return m.Name
}

func Nutritions() orm.QuerySeter {
	return orm.NewOrm().QueryTable("nutrition").OrderBy("-Id")
}


func init(){
    orm.RegisterModel(new(Nutrition))
}
