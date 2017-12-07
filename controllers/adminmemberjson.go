package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type AdminMemberJsonControllers struct {
	beego.Controller
}

type MemberInformation struct {
	Id       string
	Email    string
	Position string
	Number   int
}

func (this *AdminMemberJsonControllers) Get() {

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", 302)
	}

	idAdmin := userId.(int)

	idCongTy := models.FindCongTyByIdUser(idAdmin)
	u := models.FindMemberOfCongTy(idCongTy)

	var members []MemberInformation
	for i := 0; i < len(u); i++ {
		if u[i].Id != idAdmin {
			member := MemberInformation{
				Id:       strconv.Itoa(u[i].Id),
				Email:    u[i].Email,
				Position: models.FindPosition(u[i].IdPosition),
				Number:   len(models.FindProject(u[i].Id))}
			members = append(members, member)
		}
	}

	resBody, err := json.MarshalIndent(members, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

func (this *AdminMemberJsonControllers) Post() {

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", 302)
	}

	idAdmin := userId.(int)

	member := MemberInformation{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &member)
	if err != nil {
		fmt.Println(err)
	}

	user := models.User{
		Email:      member.Email,
		Password:   "1",
		IdCompany:  models.FindCongTyByIdUser(idAdmin),
		IdPosition: models.FindPositionWithName(member.Position),
		Status:     1}

	models.AddUser(user)
}

func (this *AdminMemberJsonControllers) Delete() {

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", 302)
	}

	member := MemberInformation{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &member)
	if err != nil {
		fmt.Println(err)
	}

	models.DeleteUserThayStatus(member.Email)
}
