package controllers

import (
	"bugmanage/models"

	"github.com/astaxie/beego"
)

type MasterLoginController struct {
	beego.Controller
}

func (this *MasterLoginController) Get() {
	this.TplName = "master/masterlogin.html"
	this.Render()
}

type Master struct {
	Email    string `form:"Email"`
	Password string `form:"Password"`
}

// Post
func (this *MasterLoginController) Login() {

	u := Master{}

	if err := this.ParseForm(&u); err != nil {
		this.Redirect("/masterlogin/", 302)
		return
	}

	isValidUser := models.CheckMaster(u.Email, u.Password) //Kiem tra mat khau

	if !isValidUser {
		this.Redirect("/masterlogin/", 302)
		return
	}

	this.Redirect("/master/", 302)
}
