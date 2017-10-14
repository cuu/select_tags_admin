package main

import (
	_ "github.com/cuu/select_tags/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/beego/compress"
	
  . "github.com/cuu/select_tags/controllers"
	 "github.com/cuu/select_tags/database"
	 "github.com/cuu/select_tags/utils"
	"fmt"
	"flag"
	"errors"
	"time"
	"runtime"
	"math/rand"
)

var (
	BinaryName ="select_tags"
	IsProMode bool
	AppUrl string
	CompressConfPath = "conf/compress.json"
)


func print_help() {
	fmt.Println(BinaryName, " usage:")
	fmt.Println(BinaryName, " -help ")
	
}


func GuuRecoverPanic(ctx *context.Context) {
	ErrAbort := errors.New("User stop run")
	if err := recover(); err != nil {
		if err == ErrAbort {
			return
		}

		var stack string
		logs.Critical("the request url is ", ctx.Input.URL())
		logs.Critical("Handler crashed with error", err)
		stack += fmt.Sprintln("the request url is ", ctx.Input.URL() )
		stack += fmt.Sprintln("Handler crashed with error: ", err)
		
		for i := 1; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			logs.Critical(fmt.Sprintf("%s:%d", file, line))
			stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
		} 
		
		if ctx.Output.Status != 0 {
			ctx.ResponseWriter.WriteHeader(ctx.Output.Status)
		} else {
			ctx.ResponseWriter.WriteHeader(500)
		}
		stack += "SelectTags "
		ctx.WriteString(stack)
	}
	
}


func settingCompress() {
	  setting, err := compress.LoadJsonConf(CompressConfPath, IsProMode, AppUrl)
  if err != nil {
    beego.Error(err)
    return
  }

  setting.RunCommand()

  if IsProMode {
    setting.RunCompress(true, false, true)
  }

  beego.AddFuncMap("compress_js", setting.Js.CompressJs)
  beego.AddFuncMap("compress_css", setting.Css.CompressCss)
}


func before_run_beego() {
	
	IsProMode = beego.BConfig.RunMode == "pro"
  if IsProMode {
    beego.SetLevel(beego.LevelInformational)
  }

	AppUrl = beego.AppConfig.String("appurl")
	
	settingCompress()

	beego.ErrorController(&ErrorController{})

	
}

func run_beego(){

	
	beego.BConfig.WebConfig.DirectoryIndex = true
	/*
	beego.BConfig.Listen.AdminEnable = true
	beego.BConfig.Listen.AdminAddr = "localhost"
	beego.BConfig.Listen.AdminPort = 8088
	*/
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "guusessionID"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 0
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 86400
	beego.BConfig.WebConfig.Session.SessionProvider = "file"


	beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFExpire = 86400 * 365
	
	beego.BConfig.WebConfig.FlashName ="GUU_SLASH"
	//beego.BConfig.WebConfig.FlashSeperator ="GUUSLASH"
	
	
	beego.BConfig.RecoverFunc = GuuRecoverPanic
	beego.BConfig.ServerName = "Apache2.0"
	
	beego.SetStaticPath("/alte", "static/thirdpart/adminlte")
	//beego.ErrorHandler("404", page_not_found)
	before_run_beego()

	utils.InitCaptcha()
	utils.InitForms()

	database.Connect()
	
	beego.Run()	
}


func main() {

	help := flag.Bool("help",false,"Print Help")

	flag.Parse()

	if *help == true {
		print_help()
		return
	}
	
	rand.Seed(time.Now().UTC().UnixNano())
	run_beego()
}

