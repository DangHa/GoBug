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

// Bug cua cac project
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

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	models.UpdateBug(bugJson)
}

func (this *UserBugJson) Delete() {
	bugJson := models.Bug{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	models.DeleteBugWithIdBug(bugJson.Id)
}

// Find bug view
type UserFindBug struct {
	beego.Controller
}

func (this *UserFindBug) Get() {
	this.TplName = "finduser.html"
	this.Render()
}

type FindBug struct {
	Id                  int
	BugName             string
	BugDescription      string
	SolutionDescription string
	User                string
	Project             string
	FoundDate           string
	UpdateDate          string
}

func (this *UserFindBug) Post() {
	bugJson := models.Bug{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	bugsjson := models.FindBugWithNameBug(bugJson.BugName)

	bugs := make([]FindBug, len(bugsjson), len(bugsjson))
	for i := 0; i < len(bugs); i++ {
		bug := FindBug{
			Id:                  bugsjson[i].Id,
			BugName:             bugsjson[i].BugName,
			BugDescription:      bugsjson[i].BugDescription,
			SolutionDescription: bugsjson[i].SolutionDescription,
			User:                models.FindUserWithIdUser(bugsjson[i].IdUser).Email,
			Project:             models.FindProjectWithIdProject(bugsjson[i].IdProject).ProjectName,
			FoundDate:           bugsjson[i].FoundDate,
			UpdateDate:          bugsjson[i].UpdateDate,
		}
		bugs[i] = bug
	}

	resBody, err := json.MarshalIndent(bugs, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bugs)

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

// hien thi cac project cua user nay
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
