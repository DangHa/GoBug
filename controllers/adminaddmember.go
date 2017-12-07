package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

type AdminAddMemberJsonControllers struct {
	beego.Controller
}

func (this *AdminAddMemberJsonControllers) Post() {
	memberJson := MemberProject{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &memberJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(memberJson)

	idusers := models.FindUserOutSideProject(memberJson.IdProject)

	var membersProject []MemberProject

	for i := 0; i < len(idusers); i++ {
		user := models.FindUserWithIdUser(idusers[i])
		if user.IdPosition != 0 {
			member := MemberProject{
				Id:       user.Id,
				Member:   user.Email,
				Position: models.FindPosition(user.IdPosition)}

			membersProject = append(membersProject, member)
		}
	}

	fmt.Println(membersProject)

	resBody, err := json.MarshalIndent(membersProject, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}
