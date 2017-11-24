package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type UserProjectJson struct {
	beego.Controller
}

var idUser = 29 // Xac nhan seesionid de xac nhan user

type userProject struct {
	Id      string
	Project string
	Mieuta  string
	Number  int
}

func (this *UserProjectJson) Get() {
	idProject := models.FindProject(idUser)

	var up []userProject
	for i := 0; i < len(idProject); i++ {
		project := models.FindProject_Project(idProject[i])
		up = append(up, userProject{Id: strconv.Itoa(idProject[i]),
			Project: project.TenProject,
			Mieuta:  project.MieutaProject,
			Number:  len(models.FindBugWithIdProject(idProject[i]))})
	}

	resBody, err := json.MarshalIndent(up, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

func (this *UserProjectJson) Post() {
	bugJson := models.Bug{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	bugJson.IdUser = idUser

	fmt.Println(bugJson)

	models.AddBug(bugJson)
}
