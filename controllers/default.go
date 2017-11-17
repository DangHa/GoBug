package controllers

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "login.html"
	this.Render()
}

var keySengrid = "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"

func SendMail(from1, to1, subject1, htmlContent1 string) {
	from := mail.NewEmail("BugManage", from1)
	subject := subject1
	to := mail.NewEmail("Member", to1)
	plainTextContent := "Bug Manage"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent1)
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
