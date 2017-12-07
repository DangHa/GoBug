package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

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
