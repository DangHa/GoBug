package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func AddUser_Project(up UserProject) {
	o := orm.NewOrm()

	qs := o.QueryTable("user_project")
	i, _ := qs.PrepareInsert()
	id, err := i.Insert(&up)
	if err != nil {
		fmt.Println(err)
		return
	}

	i.Close()

	fmt.Println("Successful add!,", id)
}

//Dang sua
func FindProject(idUser int) []int {
	o := orm.NewOrm()

	var up []*UserProject
	num, err := o.QueryTable("user_project").Filter("idUser", idUser).All(&up)

	if err == orm.ErrNoRows { // No result
		return nil
	}

	idproject := make([]int, num, num)

	for i := 0; i < len(up); i++ {
		idproject[i] = (*up[i]).IdProject
	}

	return idproject
}

func FindUser(idProject int) []int {
	o := orm.NewOrm()

	var up []*UserProject
	num, err := o.QueryTable("user_project").Filter("idProject", idProject).All(&up)

	if err == orm.ErrNoRows { // No result
		return nil
	}

	idusers := make([]int, num, num)

	for i := 0; i < len(up); i++ {
		idusers[i] = (*up[i]).IdUser
	}

	return idusers
}

func FindUserOutSideProject(idProject int) []int {
	o := orm.NewOrm()

	var up []*UserProject
	num, err := o.Raw("SELECT idUser FROM user WHERE status NOT LIKE 0 and idUser NOT IN (SELECT idUser FROM user_project WHERE idProject = ?)", idProject).QueryRows(&up)
	if err != nil {
		fmt.Println("user nums: ", num)
	}

	idusers := make([]int, num, num)

	for i := 0; i < len(up); i++ {
		idusers[i] = (*up[i]).IdUser
	}

	return idusers
}

func DeleteUserInProject(idUser, idProject int) {
	o := orm.NewOrm()

	_, err := o.QueryTable("user_project").Filter("idUser", idUser).Filter("idProject", idProject).Delete()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done!")
}
