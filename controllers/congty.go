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

	//Gui email cho nguoi vua dang ki xac nhan la da dang ki
	from := "tj.hadv@hblab.vn"
	to := ctForm.Email
	subject := "Yeu cau them cong ty"
	htmlContent := "<strong> Đã gửi yêu cầu thêm công ty của bạn!</strong><br>Vui lòng chờ xác nhận cho phép tạo công ty"

	a := SendMail(from, to, subject, htmlContent)

	if !a { //Neu ko co email nay thi dung
		this.Redirect("/signupagain/", 302)
		return
	}

	//Gui email cho Master (tj.hadv@hblab.vn) cho ket qua
	from = "daxua997@gmail.com"
	to = "tj.hadv@hblab.vn"
	subject = "Yeu cau them cong ty"
	htmlContent = "<strong>Email: </strong>" + ctForm.Email + "<br>" + "<strong>Ten Mien Cong Ty: </strong>" + ctForm.Domain +
		"<br>" +
		"<a href=\"http://localhost:8080/masterlogin/\">Master Form<a>"

	SendMail(from, to, subject, htmlContent)

	var ct = models.CongTy{TenmienCongTy: ctForm.Domain, Status: 0} // Add 1 cong ty o trang thai  0-Cho chap nhan cua master
	models.AddCongTy(ct)

	//Tao them 1 User la admin cua cong ty
	idCongty := models.FindCongTy(ctForm.Domain)                                                         // Tim idCongty vua roi
	admin := models.User{Email: ctForm.Email, Password: "1", IdCongTy: idCongty, Idvaitro: 0, Status: 0} // Password: 1 , vaitro: 0-admin
	models.AddUser(admin)

	this.Redirect("/signup/", 302)
}

// Dang ki lai
type CongTyAgainController struct {
	beego.Controller
}

// Get
func (this *CongTyAgainController) Get() {
	this.TplName = "signupagain.html"
	this.Render()
}
