package main

import (
	_ "github.com/cuu/select_tags/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"

  . "github.com/cuu/select_tags/controllers"
	
	"fmt"
	"flag"
	"errors"
	"time"
	"runtime"
	"math/rand"
)


func init_database() {
	fmt.Println("create table in database defined")
}

func init() {
	
}

func print_help() {
	fmt.Println("select_tags usage:")
	fmt.Println("select_tags -initdb")
	fmt.Println("select_tags -help ")
	
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
		stack += "SoftRadius "
		ctx.WriteString(stack)
	}
	
}


func before_run_beego() {
	

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
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

	beego.BConfig.RecoverFunc = GuuRecoverPanic
//	beego.SetStaticPath("/AdminLTE", "static/AdminLTE")
//	fmt.Println(beego.AppConfig.DefaultString("DEFAULT::Secret","NULL"))
	//beego.ErrorHandler("404", page_not_found)
	before_run_beego()
	beego.Run()	
}


func main() {

	initdb := flag.Bool("initdb",false, "Init database")
	help := flag.Bool("help",false,"Print Help")

	rand.Seed(time.Now().UTC().UnixNano())

	flag.Parse()

	if *help == true {
		print_help()
		return
	}
	if *initdb == true {
		init_database()
		return
	}
	
	run_beego()
}

