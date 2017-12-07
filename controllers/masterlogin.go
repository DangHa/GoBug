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

var idMaster = 1

// Post
func (this *MasterLoginController) Login() {

	u := Master{}

	if err := this.ParseForm(&u); err != nil {
		this.Redirect("/master/login/", 302)
		return
	}

	isValidUser := models.CheckMaster(u.Email, u.Password) //Kiem tra mat khau

	if !isValidUser {
		this.Redirect("/master/login/", 302)
		return
	}

	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		session.Set("UserID", idMaster)
	}

	this.Redirect("/master/", 302)
}
