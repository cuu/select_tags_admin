package database

import (
    "fmt"

		"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
	   _ "github.com/go-sql-driver/mysql"
)


var OB orm.Ormer

func Connect() {
  orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	
	default_sql := beego.AppConfig.DefaultString("DEFAULT::db_string","NULL")
	default_sql_type := beego.AppConfig.DefaultString("DEFAULT::db_type","NULL")

	err := orm.RegisterDataBase("default", default_sql_type, default_sql)
	if err != nil {
		 beego.Error(err)
	}
	orm.RunCommand()
  err = orm.RunSyncdb("default", false, false)
	if err != nil {
		beego.Error(err)
	}

	fmt.Println( default_sql_type, " ## " , default_sql)
	
	
}


func NewOrm() orm.Ormer {
	o:= orm.NewOrm()
	o.Using("default")
	return o
}


func StartTrans(){
	o := NewOrm()
	OB = o
	OB.Begin()
}


func Rollback() {
	OB.Rollback()
}


func Commit() {
	OB.Commit()
}

