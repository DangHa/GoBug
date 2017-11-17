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
var MasterEmail = "tj.hadv@hblab.com"

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

	//Gui email cho Master (tj.hadv@hblab.vn) cho ket qua
	from := mail.NewEmail("BugManage", "daxua997@gmail.com")
	subject := "Yeu cau them cong ty"
	to := mail.NewEmail("Master", "tj.hadv@hblab.vn")
	plainTextContent := "Email: " + ctForm.Email + "Ten Mien Cong Ty: " + ctForm.Domain
	htmlContent := "<strong>Email: </strong>" + ctForm.Email + "<br>" + "<strong>Ten Mien Cong Ty: </strong>" + ctForm.Domain +
		"<br>" +
		"<a herf=\"http://localhost:8080/master/\">Master Form<a>"

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

	this.Redirect("/signup/", 302)
}
