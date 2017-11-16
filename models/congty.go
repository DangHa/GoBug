package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// Kiem tra co ton tai cong ty khong
func CheckCongTy(domain string) bool {
	o := orm.NewOrm()

	var ct = CongTy{}
	err := o.QueryTable("congty").Filter("tenmienCongTy", domain).One(&ct)

	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}

// Tim cong ty -1 la ko thay
func FindCongTy(domain string) int {
	o := orm.NewOrm()

	var ct = CongTy{}
	err := o.QueryTable("congty").Filter("tenmienCongTy", domain).One(&ct)

	if err == orm.ErrMultiRows { // Have multiple records
		return -1
	}
	if err == orm.ErrNoRows { // No result
		return -1
	}

	return ct.Id
}

// Them Cong ty
func AddCongTy(ct CongTy) { //Tra ve
	o := orm.NewOrm()

	//qs := o.QueryTable("congty")
	id, err := o.Insert(&ct)
	if err != nil {
		fmt.Println(err)
		return
	}

	//i.Close()

	fmt.Println("Successful add!,", id)
}

// Update
func UpdateCongTy(tenmienCongTy string, status int) {
	o := orm.NewOrm()

	id, err := o.QueryTable("user").Filter("tenmienCongTy", tenmienCongTy).Update(orm.Params{
		"status": status, //Trang thai master gui 1 - chap nhan; 2 - tu choi
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}
