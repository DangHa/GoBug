package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

// View project cua cong ty
type AdminProjectJsonController struct {
	beego.Controller
}

type ProjectJSON struct {
	Id          string
	Project     string
	Description string
	BeginDate   string
	FinishDate  string
}

func (this *AdminProjectJsonController) Get() {

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", 302)
		return
	}

	idAdmin := userId.(int)
	fmt.Println("Id: ", idAdmin)

	jso := models.FindProjectWithIdAdmin(idAdmin) // Luu de xac dinh duoc admin nao dang nhap vao he thong

	resBody, err := json.MarshalIndent(jso, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

func (this *AdminProjectJsonController) Post() {

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", 302)
		return
	}

	idAdmin := userId.(int)

	// JSON chuyen ve tu master html
	project := ProjectJSON{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &project)
	if err != nil {
		fmt.Println(err)
	}

	newProject := models.Project{
		ProjectName:        project.Project,
		ProjectDescription: project.Description,
		BeginDate:          ConvertDate(project.BeginDate),
		FinishDate:         ConvertDate(project.FinishDate)}

	models.AddProject(newProject, idAdmin) // Can co IDAdmin o sessionID
}

func (this *AdminProjectJsonController) Update() {

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", 302)
	}

	// JSON chuyen ve tu master html
	project := ProjectJSON{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &project)
	if err != nil {
		fmt.Println(err)
	}

	idint, _ := strconv.Atoi(project.Id)
	pj := models.Project{
		Id:                 idint,
		ProjectName:        project.Project,
		ProjectDescription: project.Description,
		BeginDate:          project.BeginDate,
		FinishDate:         project.FinishDate}

	models.UpdateProject(pj)
}

func (this *AdminProjectJsonController) Delete() {

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", 302)
	}

	// JSON chuyen ve tu master html
	project := ProjectJSON{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &project)
	if err != nil {
		fmt.Println(err)
	}

	idpro, _ := strconv.Atoi(project.Id)
	models.DeleteProject(idpro)

}

// Cac member trong tung du an
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
		if user.IdPosition != 0 {
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

func (this *AdminMemberProjectJsonController) Update() {

}

func (this *AdminMemberProjectJsonController) Delete() {
	userandproject := models.UserProject{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userandproject)
	if err != nil {
		fmt.Println(err)
	}

	models.DeleteUserInProject(userandproject.IdUser, userandproject.IdProject)
}
