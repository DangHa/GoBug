package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

//Log Out
type LoginAdminController struct {
	beego.Controller
}

func (this *LoginAdminController) Get() {
	this.TplName = "admin/admin.html"
	this.Render()
}

func (this *LoginAdminController) LogOut() {

	// Check if user is logged in
	session := this.StartSession()
	userID := session.Get("UserID")
	if userID != nil {
		// UserID is set and can be deleted
		session.Delete("UserID")
	}
}

// Gui du lieu ve de xu li bieu do
type AdminStatControllers struct {
	beego.Controller
}

func (this *AdminStatControllers) Get() {
	this.TplName = "admin/statistic.html"
	this.Render()
}

//Bang project
type AdminStatJsonControllers struct {
	beego.Controller
}

type statAdmin struct {
	Project    string
	Member     int
	Bug        int
	Solution   int
	BeginDate  string
	FinishDate string
}

func (this *AdminStatJsonControllers) Get() { // lay idAdmin de tim

	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", redirectStatus)
		return
	}

	idAdmin := userId.(int)

	projects := models.FindProjectWithIdAdmin(idAdmin)

	var stats []statAdmin
	for i := 0; i < len(projects); i++ {
		//Dem so solution
		solu := 0
		bugs := models.FindBugWithIdProject(projects[i].Id)
		for j := 0; j < len(bugs); j++ {
			if bugs[j].SolutionDescription != "" {
				solu++
			}
		}

		stat := statAdmin{
			Project:    projects[i].ProjectName,
			Member:     len(models.FindUser(projects[i].Id)) - 1,
			Bug:        len(bugs),
			Solution:   solu,
			BeginDate:  projects[i].BeginDate,
			FinishDate: projects[i].FinishDate}

		stats = append(stats, stat)
	}

	resBody, err := json.MarshalIndent(stats, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

//Bieu do Bug
type AdminBugStatJsonControllers struct {
	beego.Controller
}

func (this *AdminBugStatJsonControllers) Get() {
	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", redirectStatus)
		return
	}

	idAdmin := userId.(int)

	category := models.FindCategoryOfCompany(idAdmin)

	resBody, err := json.MarshalIndent(category, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

//Bieu do Dev
type AdminDevStatJsonControllers struct {
	beego.Controller
}

func (this *AdminDevStatJsonControllers) Get() {
	// Check if user is logged in
	session := this.StartSession()
	userId := session.Get("UserID")

	if userId == nil {
		this.Redirect("/", redirectStatus)
		return
	}

	idAdmin := userId.(int)

	rank := models.RankDev(idAdmin)

	fmt.Println(rank)

	resBody, err := json.MarshalIndent(rank, "", "  ") //Get 200
	if err != nil {
		log.Fatal(err)
	}

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}
