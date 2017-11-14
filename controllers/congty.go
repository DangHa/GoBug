package controllers

import (
	"GoBug/models"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type CongTyController struct {
	beego.Controller
}

type CongTyForm struct {
	Eamil         string `form: "Email"`
	Password      string `form: "Password"`
	TenmienCongTy string `form: "TenMien"`
}

var keySengrid = "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"

func (this *CongTyController) Get() {
	this.TplName = "signup.html"
	this.Render()
}

func (this *CongTyController) Add() {
	ctForm := CongTyForm{}

	if err := this.ParseForm(&ctForm); err != nil {
		this.Redirect("/signup/", 302)
		return
	}

	fmt.Println("2", ctForm.Eamil, ctForm.TenmienCongTy)
	var ct = models.CongTy{TenmienCongTy: ctForm.TenmienCongTy, Status: 0} // doi nhan ket qua cua master
	models.AddCongTy(ct)

	//Gui email cho Master (tj.hadv@hblab.vn) cho ket qua
	from := mail.NewEmail("ha1", "daxua997@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("ha2", "tj.hadv@hblab.vn")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(keySengrid)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println("1")
		fmt.Println(response.Body)
		fmt.Println("2")
		fmt.Println(response.Headers)
	}
}

func (this *CongTyController) Update() {

}
