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

	// Check if user is logged in

	ctForm := CongTyForm{}

	if err := this.ParseForm(&ctForm); err != nil {
		this.Redirect("/signup/", redirectStatus)
		return
	}

	for i := 0; i < len(ctForm.Email); i++ {
		if ctForm.Email[i] == '@' && len(ctForm.Email)-5 > i {
			break
		}
		if i == len(ctForm.Email)-1 {
			this.Redirect("/signup/again/", redirectStatus)
			return
		}
	}

	//Gui email cho nguoi vua dang ki xac nhan la da dang ki
	from := "tj.hadv@hblab.vn"
	to := ctForm.Email
	subject := "Yeu cau them cong ty"
	htmlContent := "<strong> Đã gửi yêu cầu thêm công ty của bạn!</strong><br>Vui lòng chờ xác nhận cho phép tạo công ty"

	checkSend := SendMail(from, to, subject, htmlContent)

	if !checkSend { //Neu ko co email nay thi dung
		this.Redirect("/signup/again/", redirectStatus)
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

	var ct = models.Company{CompanyDomain: ctForm.Domain, Status: 0} // Add 1 cong ty o trang thai  0-Cho chap nhan cua master
	models.AddCompany(ct)

	//Tao them 1 User la admin cua cong ty
	idCongty := models.FindCompanyWithDomain(ctForm.Domain)                                                 // Tim idCongty vua roi
	admin := models.User{Email: ctForm.Email, Password: "1", IdCompany: idCongty, IdPosition: 0, Status: 0} // Password: 1 , vaitro: 0-admin
	models.AddUser(admin)

	this.Redirect("/signup/", redirectStatus)
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
