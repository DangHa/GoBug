package models

import (
	"fmt"
	"time"

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

func FindBugWithNameBug(bugName string) []Bug {
	o := orm.NewOrm()

	var bugs []Bug
	_, err := o.QueryTable("bug").Filter("bugName__icontains", bugName).All(&bugs)

	if err == orm.ErrNoRows {
		return nil
	}

	return bugs
}

func FindBugWithIdBug(idBug int) Bug {
	o := orm.NewOrm()

	var bug Bug
	err := o.QueryTable("bug").Filter("idBug", idBug).One(&bug)

	if err == orm.ErrNoRows {
		return bug
	}

	return bug
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

func UpdateBugByTester(bug Bug) {
	o := orm.NewOrm()

	id, err := o.QueryTable("bug").Filter("idBug", bug.Id).Update(orm.Params{
		"bugName":        bug.BugName,
		"bugDescription": bug.BugDescription,
		"updateDate":     time.Now(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}

func UpdateBugByDev(bug Bug) {
	o := orm.NewOrm()

	id, err := o.QueryTable("bug").Filter("idBug", bug.Id).Update(orm.Params{
		"bugName":             bug.BugName,
		"bugDescription":      bug.BugDescription,
		"solutionDescription": bug.SolutionDescription,
		"IdDev":               bug.IdDev,
		"updateDate":          time.Now(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}

type BugStatJson struct {
	Category string // Category + nguoi giai duoc nhieu nhat
	Bug      int    `orm:"column(Bug)"`
	Project  int    `orm:"column(Project)"`
}

func FindCategoryOfCompany(idAdmin int) []BugStatJson {
	o := orm.NewOrm()

	var bs []BugStatJson
	num, err := o.Raw("Select category, count(idBug) as Bug, count(idProject) as Project from bug where idProject In (select idProject from user_project where idUser = ?) group by category", idAdmin).QueryRows(&bs)
	if err != nil {
		fmt.Println("user nums: ", num)
	}

	return bs
}
