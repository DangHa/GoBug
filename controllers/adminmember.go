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
	userandproject := AddUserProject{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userandproject)
	if err != nil {
		fmt.Println(err)
	}

	iduser := models.FindIdUserWithEmail(userandproject.Email)
	if iduser == -1 {
		return
	}

	up := models.UserProject{
		IdUser:    iduser,
		IdProject: userandproject.Idproject}

	models.AddUser_Project(up)
}
