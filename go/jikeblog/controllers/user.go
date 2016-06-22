package controllers

import "jikeblog/models/class"

type UserController struct {
	//	beego.Controller
	BaseController
}

func (c *UserController) Profile() {

	maps := c.Ctx.Input.Params()
	u := &class.User{Id: maps[":id"]}
	u.ReadDB()

	c.Data["u"] = u

	a := &class.Article{Author: u}
	as := a.Gets()

	c.Data["articles"] = as

	c.TplName = "user/profile.html"
}

func (c *UserController) PageJoin() {
	c.TplName = "user/join.html"
}

func (c *UserController) PageSetting() {
	c.CheckLogin()
	c.TplName = "user/setting.html"
}
