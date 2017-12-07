package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type AdminMemberController struct {
	beego.Controller
}

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

	up := models.UserProject{
		IdUser:    userandproject.IdUser,
		IdProject: userandproject.IdProject}

	models.AddUser_Project(up)
}
