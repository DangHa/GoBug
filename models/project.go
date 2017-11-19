package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func CheckProject(tenProject string) bool {
	o := orm.NewOrm()

	var pj = Project{}

	err := o.QueryTable("project").Filter("tenProject", tenProject).One(&pj)
	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}

func AddProject(pj Project) {
	o := orm.NewOrm()

	qs := o.QueryTable("project")
	i, _ := qs.PrepareInsert()
	id, err := i.Insert(&pj)
	if err != nil {
		fmt.Println(err)
		return
	}

	i.Close()

	fmt.Println("Successful add!,", id)
}

func UpdateProject(pj Project) {
	o := orm.NewOrm()

	id, err := o.QueryTable("project").Filter("tenProject", pj.TenProject).Update(orm.Params{
		"mieutaProject": pj.MieutaProject, //Trang thai master gui 1 - chap nhan; 2 - tu choi
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}

func DeleteProject(domain string) {
	o := orm.NewOrm()

	_, err := o.QueryTable("project").Filter("tenProject", domain).Delete()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done!")
}

func FindProject_Project(id int) Project {
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
func FindProjectWithidAdmin(idAdmin int) []Project {

	idProject := FindProject(idAdmin)

	var pj []Project
	for i := 0; i < len(idProject); i++ {
		pj = append(pj, FindProject_Project(idProject[i]))
	}

	return pj
}
