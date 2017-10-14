package models

import (
	"fmt"
	"strings"
	
	"github.com/astaxie/beego/orm"
)


type SliceNutritionField []Nutrition


func (e SliceNutritionField) Value() []string {
	var d []string
	for _,p := range e {
		d = append(d,p.Name)
	}
	return d
}


func (e *SliceNutritionField) Set(d []Nutrition ) {
	
	*e = SliceNutritionField(d)
}

func (e *SliceNutritionField) Add(v string ) {
	nur := Nutrition{}
	nur.Name = v
	*e = append(*e,nur)
}

func (e *SliceNutritionField) String() string {
	return  strings.Join(e.Value(),",")
}

func (e *SliceNutritionField) FieldType() int {
	return orm.TypeCharField
}

func (e *SliceNutritionField) SetRaw(value interface{} ) error {
	switch d := value.(type) {
	case []string:
		v := make([]Nutrition,0,len(d))
		for i,p := range d{
			v[i].Name = p
		}
		e.Set(v)
	case string:
		if len(d) > 0 {
			parts := strings.Split(d,",")
			v := make([]Nutrition,0,len(parts))
			for i, p := range parts {
				v[i].Name = strings.TrimSpace(p)
			}
			e.Set(v)
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
