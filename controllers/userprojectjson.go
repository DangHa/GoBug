package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type UserProjectJson struct {
	beego.Controller
}

type userProject struct {
	Id          string
	Project     string
	Description string
	Number      int
	BeginDate   string
	FinishDate  string
	IdPosition  int
}

func (this *UserProjectJson) Get() {

	// Check if user is logged in
	session := this.StartSession()
	User := session.Get("UserID")

	if User == nil {
		return
	}

	idUser := User.(int)

	user := models.FindUserWithIdUser(idUser)

	idProject := models.FindProject(idUser)

	var up []userProject
	for i := 0; i < len(idProject); i++ {
		project := models.FindProjectWithIdProject(idProject[i])
		up = append(up, userProject{Id: strconv.Itoa(idProject[i]),
			Project:     project.ProjectName,
			Description: project.ProjectDescription,
			Number:      len(models.FindBugWithIdProject(idProject[i])),
			BeginDate:   project.BeginDate,
			FinishDate:  project.FinishDate,
			IdPosition:  user.IdPosition})
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

// Ke de post bug
func (this *UserProjectJson) Post() {

	// Check if user is logged in
	session := this.StartSession()
	User := session.Get("UserID")

	if User == nil {
		return
	}

	idUser := User.(int)

	// Lay du lieu tu json
	bugJson := models.Bug{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	bugJson.FoundDate = time.Now().String()

	bugJson.IdUser = idUser

	models.AddBug(bugJson)
}
