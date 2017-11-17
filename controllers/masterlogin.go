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
		this.Redirect("/", 302)
		return
	}

	isValidUser := models.CheckUser(u.Email, u.Password) //Kiem tra mat khau

	if isValidUser == 0 {
		this.Redirect("/", 302)
		return
	}

	if isValidUser == 1 {
		this.Redirect("/login/", 302)
	}

	this.Redirect("/loginAdmin/", 302)
}
