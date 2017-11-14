package controllers

import (
	"GoBug/models"

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
	this.TplName = "login.html"
	this.Render()
}

//Post
func (this *LoginController) Login() {
	u := User{}

	if err := this.ParseForm(&u); err != nil {
		this.Redirect("/", 302)
		return
	}

	isValidUser := models.CheckUser(u.Email, u.Password) //Kiem tra mat khau

	if !isValidUser {
		this.Redirect("/", 302)
		return
	}

	this.Redirect("/login/", 302)
}
