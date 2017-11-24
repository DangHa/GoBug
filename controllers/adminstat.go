package controllers

import (
	"bugmanage/models"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
)

type AdminStatControllers struct {
	beego.Controller
}

func (this *AdminStatControllers) Get() {
	this.TplName = "admin/statistic.html"
	this.Render()
}

type AdminStatJsonControllers struct {
	beego.Controller
}

type statAdmin struct {
	Project  string
	Member   int
	Bug      int
	Solution int
}

func (this *AdminStatJsonControllers) Get() { // lay idAdmin de tim

	projects := models.FindProjectWithidAdmin(idAdmin)

	var stats []statAdmin
	for i := 0; i < len(projects); i++ {
		//Dem so solution
		solu := 0
		bugs := models.FindBugWithIdProject(projects[i].Id)
		for j := 0; j < len(bugs); j++ {
			if bugs[j].MieutaSolution != "" {
				solu++
			}
		}

		stat := statAdmin{
			Project:  projects[i].TenProject,
			Member:   len(models.FindUser(projects[i].Id)),
			Bug:      len(bugs),
			Solution: solu}

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

func (this *AdminStatJsonControllers) Post() {

}
