package controllers

import (
	"bugmanage/models"
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

	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil { // kiem tra co session id chua
		this.TplName = "login.html"
		this.Render()
		return
	}

	idUser, ok := userId.(int)
	if !ok {
		return
	}

	checkAdmin := models.FindUserWithIdUser(idUser)
	//Kiem tra xem co la member hay admin
	if checkAdmin.IdPosition == 0 {
		this.Redirect("/loginAdmin/", 302)
		return
	}

	this.Redirect("/login/", 302)

}

var keySengrid = "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"

func SendMail(from1, to1, subject1, htmlContent1 string) bool {
	from := mail.NewEmail("BugManage", from1)
	subject := subject1
	to := mail.NewEmail("Member", to1)
	plainTextContent := "Bug Manage"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent1)
	client := sendgrid.NewSendClient(keySengrid)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return false
	} else {
		fmt.Println(response.StatusCode)
		if response.StatusCode == 400 {
			return false
		}
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		return true
	}
}

//Chuyen lai theo chuan de cho vao database
func ConvertDate(date string) string {
	return date[6:] + "-" + date[:2] + "-" + date[3:5]
}
