package models

import (
	"fmt"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/cuu/select_tags/utils"
	
)


type SliceDishField []Dish


func (e SliceDishField) Value() []string {
	var d []string
	for _,p := range e {
		idstr := utils.StrTo(p.Id).String()
		d = append(d, idstr)
	}
	return d
}

func (e SliceDishField) Label() []string {
	var d []string
	for _,p := range e {
		d = append(d,p.Name)
	}
	return d
}

func (e *SliceDishField) Set(d []string ) {
		
	v := make([]Dish,len(d))
	for i,p := range d{
		// read Dish from mysql and assign to
		id,err := utils.StrTo(p).Int()
		if err == nil {
			nur := Dish{Id:id}
			err = nur.Read()
			if err == nil {
				fmt.Println( v , " ", i, " " , len(d))
				v[i] = nur
			}else {
				beego.Error("<SliceDishField.SetRaw>: ",err)
			}
		}else {
			beego.Error("<SliceDishField.SetRaw> illegl id string ",err)
		}
	}

	*e = v
}

func (e *SliceDishField) Add(v string ) {
	id,err := utils.StrTo(v).Int()
	if err == nil {
		nur := Dish{Id:id}
		err = nur.Read()
		if err == nil {
			*e = append(*e,nur)
		}else {
			beego.Error("<SliceDishField.Add>: ",err, v)
		}
	}
}

func (e *SliceDishField) String() string {
	
	return  strings.Join(e.Value(),",")
}

func (e *SliceDishField) FieldType() int {
		return orm.TypeCharField
	//return orm.TypeIntegerField
}

func (e *SliceDishField) SetRaw(value interface{} ) error {
	switch d := value.(type) {
	case []string:
		e.Set(d)
	case string:
		if len(d) > 0 {
			parts := strings.Split(d,",")
			e.Set(parts)
		}
	default:
		return fmt.Errorf("<SliceDishField.SetRaw> unknown value `%v`",value)
	}
	return nil
}

func (e *SliceDishField) RawValue() interface{} {
	return e.String()
}

func (e *SliceDishField) Clean() error {
	return nil
}

var _ orm.Fielder = new(SliceDishField)
