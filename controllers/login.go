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
		this.Redirect("/login/again/", redirectStatus)
		return
	}

	isValidUser := models.CheckUser(u.Email, u.Password) //Kiem tra mat khau

	if isValidUser == 0 {
		this.Redirect("/login/again/", redirectStatus)
		return
	}

	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		session.Set("UserID", models.FindIdUserWithEmail(u.Email))
	}

	idUser, ok := userID.(int)
	if !ok {
		this.Redirect("/", redirectStatus)
		return
	}

	checkAdmin := models.FindUserWithIdUser(idUser)
	//Kiem tra xem co la member hay admin
	if checkAdmin.IdPosition == 0 {
		this.Redirect("/loginAdmin/", redirectStatus)
		return
	}

	this.Redirect("/login/", redirectStatus)
}

// login again
type LoginAgainController struct {
	beego.Controller
}

func (this *LoginAgainController) Get() {
	this.TplName = "loginagain.html"
	this.Render()
}
