package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
)

type AdminMemberJsonControllers struct {
	beego.Controller
}

func (this *AdminMemberJsonControllers) Get() {

	jso := models.FindProjectWithidAdmin(25) // Luu de xac dinh duoc admin nao dang nhap vao he thong

	resBody, err := json.MarshalIndent(jso, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}
