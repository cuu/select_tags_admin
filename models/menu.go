package models

import (
	"time"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


type Menu struct {
	Id int
	SpecifyDate time.Time 
	Booked SliceDishPointers `orm:"rel(m2m);rel_table(menu_booked)"` /// booked dishes 
	Extras SliceDishPointers `orm:"rel(m2m);rel_table(menu_extras)"` /// Added other dishes later 
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}



func (m *Menu) Insert() error {
	if _,err := orm.NewOrm().Insert(m);err != nil {
		return err
	}
	return nil
}

func (m *Menu) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil {
		return err
	}
	return nil
}

func (m *Menu) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...); err != nil {
		return err
	}
	return nil
}

func (m *Menu) Delete() error {
	if _,err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

/*
func (m *Menu) RemoveDishes() (int64,error) {
	m2m := orm.NewOrm().QueryM2M(m,"Dishes")

	num,err := m2m.Remove(m.Dishes)
	return num,err
}

func (m *Dish) LoadDishes() (int64,error) {
	num,err := orm.NewOrm().LoadRelated(m,"Dishes")
	fmt.Println("LoadDishes: ",num)
	return num,err
}

func (m *Menu) SetDishes( arr []Dish ) {
	if m.Id == 0 {
		beego.Error("We should Insert before QueryM2M Add")
		return
	}
	m2m := orm.NewOrm().QueryM2M(m,"Dishes")

	for i,_ := range arr {
		m2m.Add(arr[i])
	}
}
*/

func (m *Menu) SetExtras( arr []Dish) {
	if m.Id == 0 {
		beego.Error("We Should Insert before QueryM2M Add")
		return
	}

	m2m := orm.NewOrm().QueryM2M(m,"Extras")
	for i,_ := range arr {
		m2m.Add(arr[i])
	}
	
}

func (m *Menu) RemoveExtras() (int64,error) {
	m2m := orm.NewOrm().QueryM2M(m,"Extras")
	if len(m.Extras) > 0 {
		num,err :=m2m.Remove(m.Extras)
		return num,err
	}else {
		return 0,nil
	}
}

func (m *Menu) LoadExtras() (int64,error) {
	num,err := orm.NewOrm().LoadRelated(m,"Extras") 
	return num,err
}

func (m *Menu) SetBooked( arr []Dish) {
	if m.Id == 0 {
		beego.Error("We Should Insert before QueryM2M Add")
		return
	}

	m2m := orm.NewOrm().QueryM2M(m,"Booked")
	for i,_ := range arr {
		m2m.Add(arr[i])
	}
	
}

func (m *Menu) RemoveBooked() (int64,error) {
	m2m := orm.NewOrm().QueryM2M(m,"Booked")
	if len(m.Booked) > 0 {
		num,err := m2m.Remove(m.Booked)
		return num,err
	} else {
		return 0,nil
	}
}


func (m *Menu) LoadBooked() (int64,error) {
	num,err := orm.NewOrm().LoadRelated(m,"Booked")
	return num,err
}

func Menus() orm.QuerySeter {
	return orm.NewOrm().QueryTable("menu").OrderBy("-Id")
}
	
	
type SliceMenuPointers []*Menu

func (e SliceMenuPointers) Label() []string {
	var d []string
	/*
	for _,_:= range e {
		d = append(d,"Not done")
	}
*/
	return d
}

func (e *SliceMenuPointers) String() string {
	return strings.Join(e.Label(),",")
}


func init(){
	orm.RegisterModel(new(Menu))
}
