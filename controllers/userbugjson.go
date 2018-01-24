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

func (this *UserBugJson) Post() {

	bugJson := models.Bug{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	bugsjson := models.FindBugWithIdProject(bugJson.IdProject)

	bugs := make([]FindBug, len(bugsjson), len(bugsjson))
	for i := 0; i < len(bugs); i++ {
		bug := FindBug{
			Id:                  bugsjson[i].Id,
			BugName:             bugsjson[i].BugName,
			BugDescription:      bugsjson[i].BugDescription,
			SolutionDescription: bugsjson[i].SolutionDescription,
			Category:            bugsjson[i].Category,
			Tester:              models.FindUserWithIdUser(bugsjson[i].IdTest).UserName,
			Developer:           models.FindUserWithIdUser(bugsjson[i].IdDev).UserName,
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

	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Body(resBody)
	this.ServeJSONP()
}

func (this *UserBugJson) Update() {

	session := this.StartSession()
	User := session.Get("UserID")

	if User == nil {
		return
	}

	idUser := User.(int)

	bugJson := models.Bug{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	// kiem tra xem project nay user co duoc tham gia khong
	projects := models.FindProject(idUser)
	for i := 0; i < len(projects); i++ {
		if projects[i] == bugJson.IdProject {
			break
		}
		if i == len(projects)-1 {
			return
		}
	}

	//Kiem tra la Tester hay Dev xem do co phai la bug cua minh ko
	bug := models.FindBugWithIdBug(bugJson.Id)
	user := models.FindUserWithIdUser(idUser)
	if user.IdPosition == tester && user.Id == bug.IdTest {
		models.UpdateBugByTester(bugJson)
	}
	if user.IdPosition == developer && (user.Id == bug.IdDev || bug.IdDev == 0) {
		if bugJson.SolutionDescription == "" {
			bugJson.IdDev = 0
		} else {
			bugJson.IdDev = user.Id
		}
		fmt.Println(bugJson)
		models.UpdateBugByDev(bugJson)
	}

}

func (this *UserBugJson) Delete() {

	session := this.StartSession()
	User := session.Get("UserID")

	if User == nil {
		return
	}

	idUser := User.(int)

	bugJson := models.Bug{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &bugJson)
	if err != nil {
		fmt.Println(err)
	}

	// kiem tra xem project nay user co duoc tham gia khong
	projects := models.FindProject(idUser)
	bug := models.FindBugWithIdBug(bugJson.Id)

	for i := 0; i < len(projects); i++ {
		if projects[i] == bug.IdProject {
			break
		}
		if i == len(projects)-1 {
			return
		}
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
	Category            string
	Tester              string
	Developer           string
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

	fmt.Println("bugJson:", bugJson)

	bugsjson := models.FindBugWithNameBug(bugJson.BugName)

	bugs := make([]FindBug, len(bugsjson), len(bugsjson))
	for i := 0; i < len(bugs); i++ {
		bug := FindBug{
			Id:                  bugsjson[i].Id,
			BugName:             bugsjson[i].BugName,
			BugDescription:      bugsjson[i].BugDescription,
			SolutionDescription: bugsjson[i].SolutionDescription,
			Tester:              models.FindUserWithIdUser(bugsjson[i].IdTest).UserName,
			Developer:           models.FindUserWithIdUser(bugsjson[i].IdDev).UserName,
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

	if bugJson.BugName == "" {
		return
	}

	fmt.Println(bugJson)

	// kiem tra xem project nay user co duoc tham gia khong
	projects := models.FindProject(idUser)
	for i := 0; i < len(projects); i++ {
		if projects[i] == bugJson.IdProject {
			break
		}
		if i == len(projects)-1 {
			return
		}
	}

	bugJson.FoundDate = time.Now().String()

	bugJson.IdTest = idUser

	models.AddBug(bugJson)
}

// hien thi cac project cua user nay
type UserProfile struct {
	beego.Controller
}

type userProfile struct {
	Name                 string
	Password             string
	PasswordConfirmation string
}

func (this *UserProfile) Get() {
	this.TplName = "profile.html"
	this.Render()
}

func (this *UserProfile) Post() {
	// Check if user is logged in
	session := this.StartSession()
	User := session.Get("UserID")

	if User == nil {
		return
	}

	idUser := User.(int)

	// Lay du lieu tu json
	profile := userProfile{}

	if err := this.ParseForm(&profile); err != nil {
		this.Redirect("/", redirectStatus)
		return
	}

	user := models.FindUserWithIdUser(idUser)
	if profile.Name != "" {
		user.UserName = profile.Name
	}

	if profile.Password != "d41d8cd98f00b204e9800998ecf8427e" && profile.Password == profile.PasswordConfirmation { //d41d8cd98f00b204e9800998ecf8427e == ""
		user.Password = profile.Password
	}

	models.UpdateInformation(user)

	// Xóa sessions
	if User != nil {
		// UserID is set and can be deleted
		session.Delete("UserID")
	}
	this.Redirect("/", redirectStatus)
}
