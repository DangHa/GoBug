package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

type AdminProjectJsonController struct {
	beego.Controller
}

func (this *AdminProjectJsonController) Get() {

	jso := models.FindProjectWithidAdmin(1)

	resBody, err := json.MarshalIndent(jso, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

type ProjectJSON struct {
	Project string
	Mieuta  string
}

func (this *AdminProjectJsonController) Post() {
	// JSON chuyen ve tu master html
	project := ProjectJSON{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &project)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(project.Project)
}

func (this *AdminProjectJsonController) Update() {
	// JSON chuyen ve tu master html
	project := ProjectJSON{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &project)
	if err != nil {
		fmt.Println(err)
	}
}

func (this *AdminProjectJsonController) Delete() {
	// JSON chuyen ve tu master html
	project := ProjectJSON{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &project)
	if err != nil {
		fmt.Println(err)
	}
}
