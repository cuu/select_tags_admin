package controllers

import (
//		"github.com/astaxie/beego"
//	"net/http/httputil"
	"fmt"
)


//用来测试post数据 
type PtsController struct {
	BaseController
}


// @router /pts [get]
func (c *PtsController) Pts() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "post_test.tpl"
	c.Render()
	
}

// @router /pts [post]
func (c *PtsController) PtsPost() {
	c.Data["Website"] = "beego.me POST"
	c.Data["Email"] = "astaxie@gmail.com"

	fmt.Println(c.RequestDump())

	PostData := fmt.Sprintf("%s",  c.Input())

	c.Data["PostData"] = PostData
	
	c.TplName = "post_test.tpl"
	c.Render()

}
