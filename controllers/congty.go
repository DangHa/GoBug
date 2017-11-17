package controllers

import (
	"bugmanage/models"

	"github.com/astaxie/beego"
)

type CongTyController struct {
	beego.Controller
}

type CongTyForm struct {
	Email  string `form: "Email"`
	Domain string `form: "Domain"`
}

// Get
func (this *CongTyController) Get() {
	this.TplName = "signup.html"
	this.Render()
}

// Post
func (this *CongTyController) Add() {

	ctForm := CongTyForm{}

	if err := this.ParseForm(&ctForm); err != nil {
		this.Redirect("/signup/", 302)
		return
	}

	var ct = models.CongTy{TenmienCongTy: ctForm.Domain, Status: 0} // Add 1 cong ty o trang thai  0-Cho chap nhan cua master
	models.AddCongTy(ct)

	//Tao them 1 User la admin cua cong ty
	idCongty := models.FindCongTy(ctForm.Domain)                                                         // Tim idCongty vua roi
	admin := models.User{Email: ctForm.Email, Password: "1", IdCongTy: idCongty, Idvaitro: 0, Status: 0} // Password: 1 , vaitro: 0-admin
	models.AddUser(admin)

	//Gui email cho Master (tj.hadv@hblab.vn) cho ket qua
	from := "daxua997@gmail.com"
	to := "tj.hadv@hblab.vn"
	subject := "Yeu cau them cong ty"
	htmlContent := "<strong>Email: </strong>" + ctForm.Email + "<br>" + "<strong>Ten Mien Cong Ty: </strong>" + ctForm.Domain +
		"<br>" +
		"<a href=\"http://localhost:8080/masterlogin/\">Master Form<a>"

	SendMail(from, to, subject, htmlContent)

	this.Redirect("/signup/", 302)
}
