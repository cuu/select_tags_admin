package models

import (
	"fmt"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/cuu/select_tags_admin/utils"
	
)


type SliceNutritionField []Nutrition


func (e SliceNutritionField) Value() []string {
	var d []string
	for _,p := range e {
		idstr := utils.StrTo(p.Id).String()
		d = append(d, idstr)
	}
	return d
}

func (e SliceNutritionField) Label() []string {
	var d []string
	for _,p := range e {
		d = append(d,p.Name)
	}
	return d
}

func (e *SliceNutritionField) Set(d []string ) {
		
	v := make([]Nutrition,len(d))
	for i,p := range d{
		// read Nutrition from mysql and assign to
		id,err := utils.StrTo(p).Int()
		if err == nil {
			nur := Nutrition{Id:id}
			err = nur.Read()
			if err == nil {
				fmt.Println( v , " ", i, " " , len(d))
				v[i] = nur
			}else {
				beego.Error("<SliceNutritionField.SetRaw>: ",err)
			}
		}else {
			beego.Error("<SliceNutritionField.SetRaw> illegl id string ",err)
		}
	}

	*e = v
}

func (e *SliceNutritionField) Add(v string ) {
	id,err := utils.StrTo(v).Int()
	if err == nil {
		nur := Nutrition{Id:id}
		err = nur.Read()
		if err == nil {
			*e = append(*e,nur)
		}else {
			beego.Error("<SliceNutritionField.Add>: ",err, v)
		}
	}
}

func (e *SliceNutritionField) String() string {
	
	return  strings.Join(e.Value(),",")
}

func (e *SliceNutritionField) FieldType() int {
		return orm.TypeCharField
	//return orm.TypeIntegerField
}

func (e *SliceNutritionField) SetRaw(value interface{} ) error {
	switch d := value.(type) {
	case []string:
		e.Set(d)
	case string:
		if len(d) > 0 {
			parts := strings.Split(d,",")
			e.Set(parts)
		}
	default:
		return fmt.Errorf("<SliceNutritionField.SetRaw> unknown value `%v`",value)
	}
	return nil
}

func (e *SliceNutritionField) RawValue() interface{} {
	return e.String()
}

func (e *SliceNutritionField) Clean() error {
	return nil
}

var _ orm.Fielder = new(SliceNutritionField)
