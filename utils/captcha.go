package utils

/*
Beego Captcha 
*/
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var (
	Cache cache.Cache
	Captcha *captcha.Captcha
	
)
	


func InitCaptcha() {
	Cache ,err := cache.NewCache("memory",`{"interval":360}`)
	if err != nil {
		beego.Error(err)
	}
	
	Captcha = captcha.NewCaptcha("/captcha/",Cache)
	Captcha.FieldIDName = "CaptchaId"
	Captcha.FieldCaptchaName = "Captcha"
}

