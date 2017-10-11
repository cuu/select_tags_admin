package nur


import (
	"strconv"
//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/cuu/select_tags/models"
	"github.com/cuu/select_tags/utils"
	
)

type NurForm struct {
	Name string `valid:"Required;MinSize(4)"`
	Everyday string `valid:"Numeric"`
	Indication string 
}

func (form *NurForm ) Valid (v *validation.Validation) {

	
	
}


func (form *NurForm) Labels() map[string]string {
	return map[string]string {
		"Name": "nur name",
		"Everyday": "nur dosage",
		"Indication":  "nur indication",
	}
}

func (form *NurForm) Helps() map[string]string {
	return map[string]string {
//		"Name":" Name of Nurition",
//		"Indication": "indications ",
	}
}

func (form *NurForm) Placeholders() map[string]string {
	return map[string]string{
		"Name":" Name of Nurition",
		"Everyday": " Everyday dosage in mg",
		"Indication": "Nurition indications",
	}
}

func (form *NurForm) SetFromNutrition( m *models.Nutrition) {
	utils.SetFormValues(m,form)
	
}

func (form *NurForm) SaveNutrition(m *models.Nutrition) error {

	m.Name = form.Name
	if i,err := strconv.Atoi(form.Everyday);err == nil {
		m.Everyday = i
	}else {
		m.Everyday = -1
	}
	
	m.Indication = form.Indication

	return m.Insert()	
}

func (form *NurForm) UpdateNutrition( m *models.Nutrition) error {
	changes := utils.FormChanges(m,form)
	if len(changes) == 0 {
		return nil
	}

	utils.SetFormValues(form,m)

	changes = append(changes,"Updated")// Updated ==> auto_now, auto update timestamp

	return m.Update(changes...)
	
}
