package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

type UserBugJson struct {
	beego.Controller
}

func (this *UserBugJson) Get() {

}

func (this *UserBugJson) Post() {

	bugJson := models.Bug{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	bugs := models.FindBugWithIdProject(bugJson.IdProject)

	resBody, err := json.MarshalIndent(bugs, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

func (this *UserBugJson) Update() {
	bugJson := models.Bug{}

	fmt.Println(this.Ctx.Input.RequestBody, string(this.Ctx.Input.RequestBody))

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bugJson)
	models.UpdateBug(bugJson)
}

func (this *UserBugJson) Delete() {
	bugJson := models.Bug{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	models.DeleteBug(bugJson.Id)
}
