package models

import (
	"fmt"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/cuu/select_tags/utils"
	
)


type SliceIngredientField []Ingredient


func (e SliceIngredientField) Value() []string {
	var d []string
	for _,p := range e {
		idstr := utils.StrTo(p.Id).String()
		d = append(d, idstr)
	}
	return d
}

func (e SliceIngredientField) Label() []string {
	var d []string
	for _,p := range e {
		d = append(d,p.Name)
	}
	return d
}

func (e *SliceIngredientField) Set(d []string ) {
		
	v := make([]Ingredient,len(d))
	for i,p := range d{
		// read Ingredient from mysql and assign to
		id,err := utils.StrTo(p).Int()
		if err == nil {
			nur := Ingredient{Id:id}
			err = nur.Read()
			if err == nil {
				fmt.Println( v , " ", i, " " , len(d))
				v[i] = nur
			}else {
				beego.Error("<SliceIngredientField.SetRaw>: ",err)
			}
		}else {
			beego.Error("<SliceIngredientField.SetRaw> illegl id string ",err)
		}
	}

	*e = v
}

func (e *SliceIngredientField) Add(v string ) {
	id,err := utils.StrTo(v).Int()
	if err == nil {
		nur := Ingredient{Id:id}
		err = nur.Read()
		if err == nil {
			*e = append(*e,nur)
		}else {
			beego.Error("<SliceIngredientField.Add>: ",err, v)
		}
	}
}

func (e *SliceIngredientField) String() string {
	
	return  strings.Join(e.Value(),",")
}

func (e *SliceIngredientField) FieldType() int {
		return orm.TypeCharField
	//return orm.TypeIntegerField
}

func (e *SliceIngredientField) SetRaw(value interface{} ) error {
	switch d := value.(type) {
	case []string:
		e.Set(d)
	case string:
		if len(d) > 0 {
			parts := strings.Split(d,",")
			e.Set(parts)
		}
	default:
		return fmt.Errorf("<SliceIngredientField.SetRaw> unknown value `%v`",value)
	}
	return nil
}

func (e *SliceIngredientField) RawValue() interface{} {
	return e.String()
}

func (e *SliceIngredientField) Clean() error {
	return nil
}

var _ orm.Fielder = new(SliceIngredientField)
