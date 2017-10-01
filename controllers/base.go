package controllers

import (
	"github.com/astaxie/beego"
//	"github.com/cuu/select_tags/models"

// "fmt"
	
)

type GuuPreparer interface {
	GuuPrepare()
}


type GuuRender interface {
	GuuRender()
}


type BaseController struct {
	beego.Controller
}

func (this *BaseController) GET(key string ) string {
	if this.Ctx.Input.IsGet() {
		return this.GetString(key)
	}

	return ""
}

func (this *BaseController) POST(key string) string {
	if this.Ctx.Input.IsPost(){
		return this.GetString(key)
	}

	return ""
}

func (this *BaseController) Prepare(){

	
	if app, ok := this.AppController.(GuuPreparer); ok {
		app.GuuPrepare()
	}
	
}

func (c *BaseController) Render() error {
	
	if !c.EnableRender {
		return nil
	}
	
	rb, err := c.RenderBytes()
	if err != nil {
		return err
	}
	
	c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	return c.Ctx.Output.Body(rb)
}


