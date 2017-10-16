package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)



type User struct {
	Id int
	UserName    string `orm:"size(30);unique"`
	NickName    string `orm:"size(40);unique"`
	Password    string `orm:"size(128)"`
	WebSiteUrl  string `orm:"size(128)"`
	Company     string `orm:"size(50)"`
  Location    string           `orm:"size(30)"`
  Email       string           `orm:"size(80);unique"`
  GrEmail     string           `orm:"size(32)"`
  Info        string           ``  
  Github      string           `orm:"size(30)"`
  Twitter     string           `orm:"size(30)"`
  Google      string           `orm:"size(30)"`
  Weibo       string           `orm:"size(30)"`
  Linkedin    string           `orm:"size(30)"`
  Facebook    string           `orm:"size(30)"`
  PublicEmail bool             ``  
  Followers   int              ``  
  Following   int              ``  
  FavTopics   int              ``  
  IsAdmin     bool             `orm:"index"`
  IsActive    bool             `orm:"index"`
  IsForbid    bool             `orm:"index"`
  Lang        int              `orm:"index"`
  LangAdds    SliceStringField `orm:"size(50)"`
  Rands       string           `orm:"size(10)"`
  Created     time.Time        `orm:"auto_now_add"`
  Updated     time.Time        `orm:"auto_now"`
}


func Users() orm.QuerySeter {
  return orm.NewOrm().QueryTable("user").OrderBy("-Id")
}

