package controllers

import (
	"html/template"
	"reflect"
	"time"
	"net/http/httputil"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	
	"github.com/cuu/select_tags/utils"

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



func (this *BaseController) InPost() bool {
  return this.Ctx.Input.IsPost()
}

func (this *BaseController) InGet() bool {
  return this.Ctx.Input.IsGet()
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

func (this *BaseController) RequestDump() string {
	requestDump, err := httputil.DumpRequest(this.Ctx.Request, true)
	if err != nil {
		beego.Error(err)
	}
	return string(requestDump)

/*
something like this:
POST / HTTP/1.1
Host: 127.0.0.1:8080
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*\/*;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4
Cache-Control: max-age=0
Connection: keep-alive
Content-Length: 40
Content-Type: application/x-www-form-urlencoded
	Cookie: GUID_8888=Jkrbr17qaFrnIZOzvpT6; CSRF-Token-I4SDB=uTwEyY9NK4DWJPgxhPojHkxigRq6rEwY; sessionid-I4SDB=q7wtmvQqQeLMEjqNcgA2iKCAKtDXX7gv; io=VE_DCWmfBzLlMiKeAAAE; wetalk_sess=59888bcd659d1...
		
*/
	
}


func (this *BaseController) Prepare(){



  // page start time
  this.Data["PageStartTime"] = time.Now()

  // start session
  this.StartSession()


	  // check flash redirect, if match url then end, else for redirect return
  if match, redir := this.CheckFlashRedirect(this.Ctx.Request.RequestURI); redir {
    return
  } else if match {
    this.EndFlashRedirect()
  }



  // read flash message
  beego.ReadFromRequest(&this.Controller)

  // pass xsrf helper to template context
  xsrfToken := this.Controller.XSRFToken()
  this.Data["xsrf_token"] = xsrfToken
  this.Data["xsrf_html"] = template.HTML(this.Controller.XSRFFormHTML())
	
	this.Data["AppUrl"] = beego.AppConfig.String("appurl")


	this.Data["Get"] = this.GET
	this.Data["Post"] = this.POST
	this.Data["InPost"]= this.InPost
	this.Data["InGet"]  = this.InGet
	
	
  // if method is GET then auto create a form once token
  if this.Ctx.Request.Method == "GET" {
    this.FormOnceCreate()
  }


	
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





// read beego flash message
func (this *BaseController) FlashRead(key string) (string, bool) {
  if data, ok := this.Data["flash"].(map[string]string); ok {
    value, ok := data[key]
    return value, ok
  }
  return "", false
}

// write beego flash message
func (this *BaseController) FlashWrite(key string, value string) {
  flash := beego.NewFlash()
  flash.Data[key] = value
  flash.Store(&this.Controller)
}

// check flash redirect, ensure browser redirect to uri and display flash message.
func (this *BaseController) CheckFlashRedirect(value string) (match bool, redirect bool) {
  v := this.GetSession("on_redirect")
  if params, ok := v.([]interface{}); ok {
    if len(params) != 5 {
      this.EndFlashRedirect()
      goto end
    }
    uri := utils.ToStr(params[0])
    code := 302
    if c, ok := params[1].(int); ok {
      if c/100 == 3 {
        code = c
      }
    }
    flag := utils.ToStr(params[2])
    flagVal := utils.ToStr(params[3])
    times := 0
    if v, ok := params[4].(int); ok {
      times = v
    }

    times += 1
    if times > 3 {
      // if max retry times reached then end
      this.EndFlashRedirect()
      goto end
    }

    // match uri or flash flag
    if uri == value || flag == value {
      match = true
    } else {
      // if no match then continue redirect
      this.FlashRedirect(uri, code, flag, flagVal, times)
      redirect = true
    }
  }
end:
  return match, redirect
}



// create form once html
func (this *BaseController) FormOnceCreate(args ...bool) {
  var value string
  var creat bool
  creat = len(args) > 0 && args[0]
  if !creat {
    if v, ok := this.GetSession("form_once").(string); ok && v != "" {
      value = v
    } else {
      creat = true
    }
  }
  if creat {
    value = utils.GetRandomString(10)
    this.SetSession("form_once", value)
  }
  this.Data["once_token"] = value
  this.Data["once_html"] = template.HTML(`<input type="hidden" name="_once" value="` + value + `">`)
}


// set flash redirect
func (this *BaseController) FlashRedirect(uri string, code int, flag string, args ...interface{}) {
  flagVal := "true"
  times := 0
  for _, arg := range args {
    switch v := arg.(type) {
    case string:
      flagVal = v
    case int:
      times = v
    }
  }

  if len(uri) == 0 || uri[0] != '/' {
    panic("flash reirect only support same host redirect")
  }

  params := []interface{}{uri, code, flag, flagVal, times}
  this.SetSession("on_redirect", params)

  this.FlashWrite(flag, flagVal)
  this.Redirect(uri, code)
}

// clear flash redirect
func (this *BaseController) EndFlashRedirect() {
  this.DelSession("on_redirect")
}


// check form once, void re-submit
func (this *BaseController) FormOnceNotMatch() bool {
  notMatch := false
  recreat := false

  // get token from request param / header
  var value string
  if vus, ok := this.Input()["_once"]; ok && len(vus) > 0 {
    value = vus[0]
  } else {
    value = this.Ctx.Input.Header("X-Form-Once")
  }

  // exist in session
  if v, ok := this.GetSession("form_once").(string); ok && v != "" {
    // not match
    if value != v {
      notMatch = true
    } else {
      // if matched then re-creat once
      recreat = true
    }
  }

  this.FormOnceCreate(recreat)
  return notMatch
}

func (this *BaseController) validForm(form interface{}, names ...string) (bool, map[string]*validation.Error) {
  // parse request params to form ptr struct
  utils.ParseForm(form, this.Input())

  // Put data back in case users input invalid data for any section.
  name := reflect.ValueOf(form).Elem().Type().Name()
  if len(names) > 0 {
    name = names[0]
  }
  this.Data[name] = form

  errName := name + "Error"

  // check form once
  if this.FormOnceNotMatch() {
		beego.Error("BaseController validForm FormOnceNotMatch,Maybe miss .once_html in tpl")
    return false, nil
  }

  // Verify basic input.
  valid := validation.Validation{}
  if ok, _ := valid.Valid(form); !ok {
    errs := valid.ErrorMap()
    this.Data[errName] = &valid
		beego.Error(errName,errs)
    return false, errs
  }
  return true, nil
}

// valid form and put errors to tempalte context
func (this *BaseController) ValidForm(form interface{}, names ...string) bool {
  valid, _ := this.validForm(form, names...)
  return valid
}

// valid form and put errors to tempalte context
func (this *BaseController) ValidFormSets(form interface{}, names ...string) bool {
  valid, errs := this.validForm(form, names...)
  this.setFormSets(form, errs, names...)
  return valid
}

func (this *BaseController) SetFormSets(form interface{}, names ...string) *utils.FormSets {
  return this.setFormSets(form, nil, names...)
}

func (this *BaseController) setFormSets(form interface{}, errs map[string]*validation.Error, names ...string) *utils.FormSets {
  formSets := utils.NewFormSets(form, errs, nil) // utils/forms.go
  name := reflect.ValueOf(form).Elem().Type().Name()
  if len(names) > 0 {
    name = names[0]
  }
  name += "Sets"
  this.Data[name] = formSets

  return formSets
}

// add valid error to FormError
func (this *BaseController) SetFormError(form interface{}, fieldName, errMsg string, names ...string) {
  name := reflect.ValueOf(form).Elem().Type().Name()
  if len(names) > 0 {
    name = names[0]
  }
  errName := name + "Error"
  setsName := name + "Sets"

  if valid, ok := this.Data[errName].(*validation.Validation); ok {
    //valid.SetError(fieldName, this.Tr(errMsg))
		valid.SetError(fieldName, errMsg)
  }

  if fSets, ok := this.Data[setsName].(*utils.FormSets); ok {
    fSets.SetError(fieldName, errMsg)
  }
}

// check xsrf and show a friendly page
func (this *BaseController) CheckXsrfCookie() bool {
  return this.Controller.CheckXSRFCookie()
}

func (this *BaseController) SystemException() {

}

func (this *BaseController) IsAjax() bool {
  return this.Ctx.Input.Header("X-Requested-With") == "XMLHttpRequest"
}

func (this *BaseController) SetPaginator(per int, nums int64) *utils.Paginator {
  p := utils.NewPaginator(this.Ctx.Request, per, nums)
  this.Data["paginator"] = p
  return p
}

func (this *BaseController) JsStorage(action, key string, values ...string) {
  value := action + ":::" + key
  if len(values) > 0 {
    value += ":::" + values[0]
  }
  this.Ctx.SetCookie("JsStorage", value, 1<<31-1, "/", nil, nil, false)
}

