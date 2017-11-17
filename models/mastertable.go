package models

import (
	"github.com/astaxie/beego/orm"
)

//Kiem tra mat khau cho login
func CheckMaster(email, password string) bool { // Tra ve  0 - Sai user; 1 - Admin Dang nhap thanh cong; 2 - nhan vien dang nhap thanh cong
	o := orm.NewOrm()

	var user = Master{}

	err := o.QueryTable("master").Filter("Email", email).Filter("Password", password).One(&user)

	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}
