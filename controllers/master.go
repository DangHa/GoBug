package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
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

	//Tao them 1 User la admin cua cong ty
	models.UpdateUser(ctForm.Email)

	//Update status cong ty thanh 1 -- Bat dau hoat dong
	models.UpdateCompany(ctForm.Domain, 1)

	//Gui email cho Admin ket qua tra ve
	from := "tj.hadv@hblab.vn"
	subject := "Yeu cau them cong ty"
	to := ctForm.Email
	htmlContent := "<strong>" + "Domain: " + ctForm.Domain + " Created!" + "</strong>" + "<br>" + "Welcome to GoBug!"

	SendMail(from, to, subject, htmlContent)

}

//Delete
func (this *MasterController) Delete() {

	// JSON chuyen ve tu master html
	ctForm := CongTyForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ctForm)
	if err != nil {
		fmt.Println(err)
	}

	models.DeleteUser(ctForm.Email)
	models.DeleteCompany(ctForm.Domain)

	// Gui email cho Admin ket qua tra ve
	from := "tj.hadv@hblab.vn"
	subject := "Yeu cau them cong ty"
	to := ctForm.Email
	htmlContent := "<strong>" + "Domain: " + ctForm.Domain + " didn't create!" + "</strong>" + "<br>" + "Sorry! Ahihi"

	SendMail(from, to, subject, htmlContent)

}
