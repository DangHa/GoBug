package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func FindBugWithIdProject(idProject int) []Bug {
	o := orm.NewOrm()

	var bugs []Bug
	_, err := o.QueryTable("bug").Filter("idProject", idProject).All(&bugs)

	if err == orm.ErrNoRows { // No result
		return nil
	}

	return bugs
}

func AddBug(bug Bug) {
	o := orm.NewOrm()

	qs := o.QueryTable("bug")
	i, _ := qs.PrepareInsert()
	id, err := i.Insert(&bug)
	if err != nil {
		fmt.Println(err)
		return
	}

	i.Close()
	fmt.Println("Successful add!,", id)
}

func DeleteBugWithIdBug(id int) {
	o := orm.NewOrm()

	_, err := o.QueryTable("bug").Filter("idBug", id).Delete()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done!")
}

func UpdateBug(bug Bug) {
	o := orm.NewOrm()

	id, err := o.QueryTable("bug").Filter("idBug", bug.Id).Update(orm.Params{
		"bugName":             bug.BugName,
		"bugDescription":      bug.BugDescription,
		"solutionDescription": bug.SolutionDescription,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}
