package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//Kiem tra mat khau cho login
func CheckUser(email, password string) int { // Tra ve  0 - Sai user; 1 - Admin Dang nhap thanh cong; 2 - nhan vien dang nhap thanh cong
	o := orm.NewOrm()

	var user = User{}
	err := o.QueryTable("user").Filter("Email", email).Filter("Password", password).One(&user)

	if err == orm.ErrMultiRows { // Have multiple records
		return 0
	}
	if err == orm.ErrNoRows { // No result
		return 0
	}

	if user.Idvaitro == 0 { // Admin
		return 1
	}
	return 2 //Nhan Vien
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
