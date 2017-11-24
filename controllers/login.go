package controllers

import (
	"bugmanage/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type User struct {
	Email    string `form:"Email"`
	Password string `form:"Password"`
}

func (this *LoginController) Get() {
	this.TplName = "index.html"
	this.Render()
}

//Post
func (this *LoginController) Login() {
	u := User{}

	if err := this.ParseForm(&u); err != nil {
		this.Redirect("/loginagain/", 302)
		return
	}

	isValidUser := models.CheckUser(u.Email, u.Password) //Kiem tra mat khau

	if isValidUser == 0 {
		this.Redirect("/loginagain/", 302)
		return
	}

	if isValidUser == 1 {
		this.Redirect("/loginAdmin/", 302)
	}

	this.Redirect("/login/", 302)
}
