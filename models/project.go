package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func CheckProject(projectName string) bool {
	o := orm.NewOrm()

	var pj = Project{}

	err := o.QueryTable("project").Filter("projectName", projectName).One(&pj)
	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}

func AddProject(pj Project, idadmin int) {
	o := orm.NewOrm()

	qs := o.QueryTable("project")
	i, _ := qs.PrepareInsert()
	id, err := i.Insert(&pj)
	if err != nil {
		fmt.Println(err)
		return
	}

	i.Close()

	// Tim id cua project vua them
	err = o.QueryTable("project").Filter("projectName", pj.ProjectName).One(&pj)
	if err == orm.ErrMultiRows { // Have multiple records
		return
	}
	if err == orm.ErrNoRows { // No result
		return
	}

	up := UserProject{IdUser: idadmin, IdProject: pj.Id}
	AddUser_Project(up)

	fmt.Println("Successful add!,", id)
}

func UpdateProject(pj Project) {
	o := orm.NewOrm()

	id, err := o.QueryTable("project").Filter("idProject", pj.Id).Update(orm.Params{
		"projectname":        pj.ProjectName,
		"projectDescription": pj.ProjectDescription, //Trang thai master gui 1 - chap nhan; 2 - tu choi
		"beginDate":          pj.BeginDate,
		"finishDate":         pj.FinishDate,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}

func DeleteProject(idproject int) {
	o := orm.NewOrm()

	u := FindUser(idproject)
	for i := 0; i < len(u); i++ {
		DeleteUserInProject(u[i], idproject)
	}

	_, err := o.QueryTable("project").Filter("idProject", idproject).Delete()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done!")
}

func FindProjectWithIdProject(id int) Project {
	o := orm.NewOrm()

	var pj = Project{}
	err := o.QueryTable("project").Filter("idProject", id).One(&pj)

	if err == orm.ErrMultiRows { // Have multiple records
		return pj
	}
	if err == orm.ErrNoRows { // No result
		return pj
	}

	return pj
}

//Liet ke ra bang idAdmin so du an cua cong ty
func FindProjectWithIdAdmin(idAdmin int) []Project {

	idProject := FindProject(idAdmin)

	var pj []Project
	for i := 0; i < len(idProject); i++ {
		pj = append(pj, FindProjectWithIdProject(idProject[i]))
	}

	return pj
}
