package controllers

import "github.com/astaxie/beego"

type LoginAdminController struct {
	beego.Controller
}

func (this *LoginAdminController) Get() {
	this.TplName = "admin/admin.html"
	this.Render()
}
