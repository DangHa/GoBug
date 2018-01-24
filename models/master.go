package models

import (
	"github.com/astaxie/beego/orm"
)

//Kiem tra mat khau cho login
func CheckMaster(email, password string) bool {
	o := orm.NewOrm()

	var user = Master{}

	err := o.QueryTable("master").Filter("email", email).Filter("password", password).One(&user)

	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}
