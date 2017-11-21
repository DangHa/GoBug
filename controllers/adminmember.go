package controllers

import "github.com/astaxie/beego"

type AdminMemberController struct {
	beego.Controller
}

func (this *AdminMemberController) Get() {
	this.TplName = "admin/adminmember.html"
	this.Render()
}
