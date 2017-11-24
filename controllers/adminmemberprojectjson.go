package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

type AdminMemberProjectJsonController struct {
	beego.Controller
}

//create table
func (this *AdminMemberProjectJsonController) Get() {

}

type MemberProject struct {
	Id        int
	Member    string
	Position  string
	IdProject int
}

func (this *AdminMemberProjectJsonController) Post() {
	memberJson := MemberProject{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &memberJson)
	if err != nil {
		fmt.Println(err)
	}

	idusers := models.FindUser(memberJson.IdProject)

	var membersProject []MemberProject

	for i := 0; i < len(idusers); i++ {
		user := models.FindUserWithIdUser(idusers[i])
		if user.Idvaitro != 0 {
			member := MemberProject{
				Id:       user.Id,
				Member:   user.Email,
				Position: models.FindVaitro(user.Idvaitro)}

			membersProject = append(membersProject, member)
		}
	}

	resBody, err := json.MarshalIndent(membersProject, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

func (this *AdminMemberProjectJsonController) Update() {

}

func (this *AdminMemberProjectJsonController) Delete() {
	userandproject := models.User_project{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userandproject)
	if err != nil {
		fmt.Println(err)
	}

	models.DeleteUserKhoiProject(userandproject.IdUser, userandproject.IdProject)
}
