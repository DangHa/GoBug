package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type AdminProjectJsonController struct {
	beego.Controller
}

type ProjectJSON struct {
	Id      string
	Project string
	Mieuta  string
}

func (this *AdminProjectJsonController) Get() {

	jso := models.FindProjectWithidAdmin(25) // Luu de xac dinh duoc admin nao dang nhap vao he thong

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
	// JSON chuyen ve tu master html
	project := ProjectJSON{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &project)
	if err != nil {
		fmt.Println(err)
	}

	newProject := models.Project{TenProject: project.Project, MieutaProject: project.Mieuta}

	models.AddProject(newProject, 25) // Can co IDAdmin o sessionID
}

func (this *AdminProjectJsonController) Update() {
	// JSON chuyen ve tu master html
	project := ProjectJSON{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &project)
	if err != nil {
		fmt.Println(err)
	}

	idint, _ := strconv.Atoi(project.Id)
	pj := models.Project{Id: idint, TenProject: project.Project, MieutaProject: project.Mieuta}

	models.UpdateProject(pj)
}

func (this *AdminProjectJsonController) Delete() {
	// JSON chuyen ve tu master html
	project := ProjectJSON{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &project)
	if err != nil {
		fmt.Println(err)
	}

	idpro, _ := strconv.Atoi(project.Id)
	models.DeleteProject(idpro)

}
