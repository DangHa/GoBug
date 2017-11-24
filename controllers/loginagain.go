package controllers

import (
	"github.com/astaxie/beego"
)

type LoginAgainController struct {
	beego.Controller
}

func (this *LoginAgainController) Get() {
	this.TplName = "loginagain.html"
	this.Render()
}
