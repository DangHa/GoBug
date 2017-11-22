package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

type AdminMemberJsonControllers struct {
	beego.Controller
}

type MemberInformation struct {
	Id     int
	Email  string
	Vaitro string
	Number int
}

var idAdmin = 25 //Id Cua chu cong ty phai sua lai thanh sessionID

func (this *AdminMemberJsonControllers) Get() {

	idCongTy := models.FindCongTyByidUser(idAdmin)
	u := models.FindMemberOfCongTy(idCongTy)

	var members []MemberInformation
	for i := 0; i < len(u); i++ {
		if u[i].Id != idAdmin {
			member := MemberInformation{
				Id:     u[i].Id,
				Email:  u[i].Email,
				Vaitro: models.FindVaitro(u[i].Idvaitro),
				Number: len(models.FindProject(u[i].Id))}
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
	member := MemberInformation{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &member)
	if err != nil {
		fmt.Println(err)
	}

	user := models.User{
		Email:    member.Email,
		Password: "1",
		IdCongTy: models.FindCongTyByidUser(idAdmin),
		Idvaitro: models.FindVaitroWithName(member.Vaitro),
		Status:   1}

	models.AddUser(user)
}

func (this *AdminMemberJsonControllers) Delete() {
	member := MemberInformation{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &member)
	if err != nil {
		fmt.Println(err)
	}

	models.DeleteUserThayStatus(member.Email)
}
