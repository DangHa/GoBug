package controllers

import (
	"bugmanage/models"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type CongTyController struct {
	beego.Controller
}

var keySengrid = "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"

// Get
func (this *CongTyController) Get() {
	this.TplName = "signup.html"
	this.Render()
}

// Post
func (this *CongTyController) Add() {
	type CongTyForm struct {
		Email  string `form: "Email"`
		Domain string `form: "Domain"`
	}

	ctForm := CongTyForm{}

	if err := this.ParseForm(&ctForm); err != nil {
		this.Redirect("/signup/", 302)
		return
	}

	var ct = models.CongTy{TenmienCongTy: ctForm.Domain, Status: 0} // Add 1 cong ty o trang thai  0-Cho chap nhan cua master
	models.AddCongTy(ct)

	//Gui email cho Master (tj.hadv@hblab.vn) cho ket qua
	from := mail.NewEmail("BugManage", "daxua997@gmail.com")
	subject := "Them cong ty va admin"
	to := mail.NewEmail("Master", "tj.hadv@hblab.vn")
	plainTextContent := "Email: " + ctForm.Email + "Ten Mien Cong Ty: " + ctForm.Domain
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(keySengrid)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// Put
func (this *CongTyController) Update() {
	type CongTyForm struct {
		Email  string `form: "Email"`
		Domain string `form: "Domain"`
		Status int    `form: "Status"`
	}

	// Form chuyen ve tu email
	ctForm := CongTyForm{}

	if err := this.ParseForm(&ctForm); err != nil {
		this.Redirect("/signup/", 302)
		return
	}

	models.UpdateCongTy(ctForm.Domain, ctForm.Status)
	result := "Tu choi!"

	if ctForm.Status == 1 { // Neu chap nhan
		//Tao them 1 User la admin cua cong ty
		idCongty := models.FindCongTy(ctForm.Domain)
		admin := models.User{Email: ctForm.Email, Password: "1", IdCongTy: idCongty, Idvaitro: 0} // Password: 1 , vaitro: 0-admin
		models.AddUser(admin)

		result = "Chap Nhan!"
	}

	//Gui email cho Admin ket qua tra ve
	from := mail.NewEmail("Master", "tj.hadv@hblab.vn")
	subject := "Yeu cau them cong ty"
	to := mail.NewEmail("Admin", ctForm.Email)
	plainTextContent := "Email: " + ctForm.Email + "Ten Mien Cong Ty: " + ctForm.Domain + " " + result
	htmlContent := "<strong>" + plainTextContent + "</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(keySengrid)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
