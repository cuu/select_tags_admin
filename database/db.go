package database

import (
    "fmt"
    "github.com/astaxie/beego/orm"
	   _ "github.com/go-sql-driver/mysql"
		"github.com/astaxie/beego"
)


func init() {
  orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	
	orm.RegisterDataBase("default", "mysql", "root:newpass@/dishtag?charset=utf8")

	fmt.Println(beego.AppConfig.DefaultString("DEFAULT:db_string","NULL") )
	
}


func NewOrm() orm.Ormer {
	o:= orm.NewOrm()
	o.Using("default")
	return o
}
