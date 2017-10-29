package models

import (
	"time"
	"strings"
	"github.com/astaxie/beego/orm"
)

type Image struct{
	Id int
	FileName string
	Hash     string
	Width    int
	Height   int
	Ext      int       // 1=>jpg,2=>png,3=>gif
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
	Dishes  SliceDishPointers `orm:"reverse(many);rel_table(dish_img)"`
}



func (m *Image) GetExt() string{
	var ext string
	switch m.Ext {
	case 1:
		ext = ".jpg"
	case 2:
		ext = ".png"
	case 3:
		ext = ".gif"
	default:
		ext = ".error"
	}
	return ext
	
}


type SliceImagePointers []*Image


func (e SliceImagePointers) Label() []string {
	var d []string
	for _,p := range e {
		d = append(d,p.FileName)
	}
	return d
}

func (e *SliceImagePointers) String() string {
	return strings.Join(e.Label(),",")
}


func init(){
	orm.RegisterModel(new(Image))
}
