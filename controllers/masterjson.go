package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
)

type MasterJsonController struct {
	beego.Controller
}

func (this *MasterJsonController) Get() {

	ct := models.FindCongTyTheoStatus(0)

	type jsonDomain struct {
		Domain []string
	}

	var jso = jsonDomain{Domain: ct}

	resBody, err := json.MarshalIndent(jso, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}
