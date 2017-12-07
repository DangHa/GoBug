package controllers

import (
	"github.com/astaxie/beego"
)

type LoginAdminController struct {
	beego.Controller
}

func (this *LoginAdminController) Get() {
	this.TplName = "admin/admin.html"
	this.Render()
}

func (this *LoginAdminController) LogOut() {

	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID != nil {
		// UserID is set and can be deleted
		session.Delete("UserID")
	}
}
