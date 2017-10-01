package models

import "github.com/astaxie/beego/orm"

type Nutrition struct {
  Id   int
	Name string
	Everyday int
	indication string
}

func init(){
    orm.RegisterModel(new(Nutrition))
}

