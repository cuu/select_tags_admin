package controllers

import (
//		"github.com/astaxie/beego"
//	"net/http/httputil"
	"fmt"
)

type MainController struct {
	BaseController
}


// @router / [get]
func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.Render()
	
}

// @router / [post]
func (c *MainController) IndexPost() {
	c.Data["Website"] = "beego.me POST"
	c.Data["Email"] = "astaxie@gmail.com"


	fmt.Println(c.RequestDump())

	PostData := fmt.Sprintf("%s",  c.Input())

	c.Data["PostData"] = PostData
	
	c.TplName = "index.tpl"
	c.Render()

}
