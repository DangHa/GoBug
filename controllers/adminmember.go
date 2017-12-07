package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type AdminMemberController struct {
	beego.Controller
}

//View
func (this *AdminMemberController) Get() {
	this.TplName = "admin/adminmember.html"
	this.Render()
}

type AddUserProject struct {
	Idproject int
	Email     string
}

func (this *AdminMemberController) Post() {
	userandproject := models.UserProject{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userandproject)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userandproject)

	up := models.UserProject{
		IdUser:    userandproject.IdUser,
		IdProject: userandproject.IdProject}

	models.AddUser_Project(up)
}

// Tim member chua add vao project
type AdminAddMemberJsonControllers struct {
	beego.Controller
}

func (this *AdminAddMemberJsonControllers) Post() {
	memberJson := MemberProject{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &memberJson)
	if err != nil {
		fmt.Println(err)
	}

	idusers := models.FindUserOutSideProject(memberJson.IdProject)

	var membersProject []MemberProject

	for i := 0; i < len(idusers); i++ {
		user := models.FindUserWithIdUser(idusers[i])
		if user.IdPosition != adminPosition {
			member := MemberProject{
				Id:       user.Id,
				Member:   user.Email,
				Position: models.FindPosition(user.IdPosition)}

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

// Tra ve member cua cong ty them member vao cong ty
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

	models.DeleteUserByUpdateStatus(member.Email)
}
