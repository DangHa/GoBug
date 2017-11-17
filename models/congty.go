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
	err := o.QueryTable("cong_ty").Filter("tenmienCongTy", domain).One(&ct)

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

	qs := o.QueryTable("cong_ty")
	i, _ := qs.PrepareInsert()
	id, err := i.Insert(&ct)
	if err != nil {
		fmt.Println(err)
		return
	}

	i.Close()

	fmt.Println("Successful add!,", id)
}

// Update
func UpdateCongTy(tenmienCongTy string, status int) {
	o := orm.NewOrm()

	id, err := o.QueryTable("cong_ty").Filter("tenmienCongTy", tenmienCongTy).Update(orm.Params{
		"status": status, //Trang thai master gui 1 - chap nhan; 2 - tu choi
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}

func DeleteCongTy(tenmienCongTy string) {
	o := orm.NewOrm()

	_, err := o.QueryTable("cong_ty").Filter("tenmienCongTy", tenmienCongTy).Delete()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done!")
}

type Domain_Email struct {
	Email  string
	Domain string
}

// Dung cho Master
func FindCongTyTheoStatus(s int) []Domain_Email { // 0 - la chua hoat dong, 1 - la da hoat dong
	o := orm.NewOrm()

	var ct []*CongTy
	num, err := o.QueryTable("cong_ty").Filter("status", s).All(&ct)
	if err == orm.ErrNoRows { // No result
		return nil
	}

	domains := make([]Domain_Email, 0, num)
	for i := 0; i < len(ct); i++ {
		email := FindUserWithIdCongTy((*ct[i]).Id)
		domain := Domain_Email{Domain: (*ct[i]).TenmienCongTy, Email: email}
		domains = append(domains, domain)
	}

	return domains
}
