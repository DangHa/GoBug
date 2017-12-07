package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"

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

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", 302)
	}

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

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", 302)
	}

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

// Log in master
type MasterLoginController struct {
	beego.Controller
}

func (this *MasterLoginController) Get() {
	this.TplName = "master/masterlogin.html"
	this.Render()
}

type Master struct {
	Email    string `form:"Email"`
	Password string `form:"Password"`
}

// Post
func (this *MasterLoginController) Login() {

	u := Master{}

	if err := this.ParseForm(&u); err != nil {
		this.Redirect("/master/login/", redirectStatus)
		return
	}

	isValidUser := models.CheckMaster(u.Email, u.Password) //Kiem tra mat khau

	if !isValidUser {
		this.Redirect("/master/login/", redirectStatus)
		return
	}

	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")

	if userID == nil {
		session.Set("UserID", idMaster)
	}

	this.Redirect("/master/", redirectStatus)
}

// Hien cac cong ty cho duoc thanh lap
type MasterJsonController struct {
	beego.Controller
}

func (this *MasterJsonController) Get() {

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", redirectStatus)
	}

	jso := models.FindCompanyWithStatus(waitStatus)

	resBody, err := json.MarshalIndent(jso, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

// Lay cac cong ty dang hoat dong
type MasterJsonCongTyController struct {
	beego.Controller
}

type MasterJsonActiveController struct {
	beego.Controller
}

// Get
func (this *MasterJsonActiveController) Get() {
	this.TplName = "master/mastercongty.html"
	this.Render()
}

func (this *MasterJsonCongTyController) Get() {

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", redirectStatus)
	}

	jso := models.FindCompanyWithStatus(activeStatus)

	resBody, err := json.MarshalIndent(jso, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}
