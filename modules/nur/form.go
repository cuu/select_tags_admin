package nur


import (
//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	
)

type AddNurForm struct {
	NurName string `valid:"Required;MinSize(4)"`
	EverydayDosage int
	Indication string 
}

func (form *AddNurForm ) Valid (v *validation.Validation) {
	return 
}


func (form *AddNurForm) Labels() map[string]string {
	return map[string]string {
		"NurName": "nur name",
		"EverydayDosage": "nur dosage",
		"Indication":  "nur indication",
	}
}

func (form *AddNurForm) Helps() map[string]string {
	return map[string]string {
		"NurName":" Name of Nurition",
		"Indication": "indications ",
	}
}

func (form *AddNurForm) Placeholders() map[string]string {
	return map[string]string{
		"NurName":" Name of Nurition",
		"EverydayDosage": " Everyday dosage in mg",
		"Indication": "Nurition indications",
	}
}
