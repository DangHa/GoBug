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

}
