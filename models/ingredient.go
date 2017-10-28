package models

import (
	"fmt"
	"time"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/cuu/select_tags_admin/utils"
)

type Ingredient struct {
	Id int
	Name string `orm:"size(255);unique"`
	Nutritions SliceNutritionPointers `orm:"rel(m2m)"`
	Dishes SliceDishPointers `orm:"reverse(many)"`
	NatureProperty int
	FiveProperty int
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}


func (m *Ingredient) Insert() error {
	if _,err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Ingredient) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m,fields...); err != nil {
		return err
	}
	return nil
}

func (m *Ingredient) Update(fields ...string) error {
	if _,err := orm.NewOrm().Update(m,fields...); err != nil {
		return err
	}
	return nil
}

func (m *Ingredient) Delete() error {
	_,err := orm.NewOrm().Delete(m)
	return err
}

func (m *Ingredient) String() string {
	return m.Name
}


func (m *Ingredient) RemoveNutritions() (int64,error) {
	m2m := orm.NewOrm().QueryM2M(m,"Nutritions")

	if len(m.Nutritions) > 0 {
		num,err := m2m.Remove(m.Nutritions)
		return num,err
	}else {
		return 0,nil
	}
}

func (m *Ingredient) LoadNutritions() (int64,error) {
	num,err := orm.NewOrm().LoadRelated(m,"Nutritions")
	fmt.Println("LoadNutritions: ",num)
	return num,err
}

func (m *Ingredient) SetNutritions( arr []Nutrition ) {
	if m.Id == 0 {
		beego.Error("We should Insert before QueryM2M Add")
		return
	}
	m2m := orm.NewOrm().QueryM2M(m,"Nutritions")

	for i,_ := range arr {
		m2m.Add(arr[i])
	}
}


func Ingredients() orm.QuerySeter {
	return orm.NewOrm().QueryTable("ingredient").OrderBy("-Id")
}


type SliceIngredientPointers []*Ingredient

func (e SliceIngredientPointers) Label() []string {
	var d []string
	for _,p := range e {
		d = append(d,p.Name)
	}
	return d

}

func (e SliceIngredientPointers) Ids() []string {
	var d []string
	for _,p := range e {
		d = append(d,utils.ToStr(p.Id) )
	}
	return d

}

func (e SliceIngredientPointers) String() string {
	return strings.Join(e.Label(),",")
}

func (e SliceIngredientPointers) IdString() string {
	return strings.Join(e.Ids(),",")
}


func init() {
	orm.RegisterModel(new(Ingredient))
}

