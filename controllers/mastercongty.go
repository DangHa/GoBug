package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
)

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

	jso := models.FindCongTyTheoStatus(1)

	resBody, err := json.MarshalIndent(jso, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}
