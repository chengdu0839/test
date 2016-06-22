package controllers

import (
	"jikeblog/models/class"

	"github.com/astaxie/beego"
)

// 自定义基类
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	if c.IsLogin() {
		c.Data["user"] = c.GetSession("user").(class.User)
	}
}

// 设置session
func (c *BaseController) DoLogin(u class.User) {
	c.SetSession("user", u)
}

// 注销，销毁session
func (c *BaseController) DoLogout() {
	c.DestroySession()
	c.Redirect("/join", 302)
}

// 判断用户是否登陆
func (c *BaseController) IsLogin() bool {
	return c.GetSession("user") != nil
}

// 判断用户是否登陆，如果没有，跳转到登陆页面
func (c *BaseController) CheckLogin() {
	if !c.IsLogin() {
		c.Redirect("/join", 302)
		c.Abort("302")
	}
}
