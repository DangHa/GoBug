package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MasterController struct {
	beego.Controller
}

// Get
func (this *MasterController) Get() {
	this.TplName = "master/master.html"
	this.Render()
}

// Put
func (this *MasterController) Update() {

	// JSON chuyen ve tu master html
	ctForm := CongTyForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ctForm)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("12321: ", ctForm)

	//Update status cong ty thanh 1 -- Bat dau hoat dong
	models.UpdateCongTy(ctForm.Domain, 1)

	//Tao them 1 User la admin cua cong ty
	idCongty := models.FindCongTy(ctForm.Domain)                                              // Tim idCongty vua roi
	admin := models.User{Email: ctForm.Email, Password: "1", IdCongTy: idCongty, Idvaitro: 0} // Password: 1 , vaitro: 0-admin
	models.AddUser(admin)

	//Gui email cho Admin ket qua tra ve
	from := mail.NewEmail("Master", "tj.hadv@hblab.vn")
	subject := "Yeu cau them cong ty"
	to := mail.NewEmail("Admin", ctForm.Email)
	plainTextContent := "Domain: " + ctForm.Domain + " Created!"
	htmlContent := "<strong>" + plainTextContent + "</strong>" + "<br>" + "Welcome to GoBug!"

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

	this.Redirect("/master/", 302)
}

//Delete
func (this *MasterController) Delete() {

	// JSON chuyen ve tu master html
	ctForm := CongTyForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ctForm)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("delete: ", ctForm)

	//Gui email cho Admin ket qua tra ve
	from := mail.NewEmail("Master", "tj.hadv@hblab.vn")
	subject := "Yeu cau them cong ty"
	to := mail.NewEmail("Admin", ctForm.Email)
	plainTextContent := "Domain: " + ctForm.Domain + " didn't created!"
	htmlContent := "<strong>" + plainTextContent + "</strong>" + "<br>" + "Sorry! Ahihi"

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

	models.DeleteCongTy(ctForm.Domain)

	this.Redirect("/master/", 302)
}
