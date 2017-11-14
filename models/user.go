package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//Kiem tra mat khau cho login
func CheckUser(email, password string) bool {
	o := orm.NewOrm()

	var user = User{}
	err := o.QueryTable("user").Filter("Email", email).Filter("Password", password).One(&user)

	u := User{Email: "132", Password: "223", IdCongTy: 1, Idvaitro: 1}
	AddUser(u)
	ct := CongTy{TenmienCongTy: "hdfa", Status: 0}
	AddCongTy(ct)
	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}

// Them User
func AddUser(u User) {
	o := orm.NewOrm()

	qs := o.QueryTable("user")
	i, _ := qs.PrepareInsert()
	id, err := i.Insert(&u)
	if err != nil {
		fmt.Println(err)
		return
	}

	i.Close()

	fmt.Println("Successful add!,", id)
}
