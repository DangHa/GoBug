package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// Kiem tra co ton tai cong ty khong
func CheckCompanyWithDomain(domain string) bool {
	o := orm.NewOrm()

	var ct = Company{}
	err := o.QueryTable("company").Filter("companyDomain", domain).One(&ct)

	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}

// Tim cong ty -1 la ko thay
func FindCompanyWithDomain(domain string) int {
	o := orm.NewOrm()

	var ct = Company{}
	err := o.QueryTable("company").Filter("companyDomain", domain).One(&ct)

	if err == orm.ErrMultiRows { // Have multiple records
		return notFound
	}
	if err == orm.ErrNoRows { // No result
		return notFound
	}

	return ct.Id
}

// Them Cong ty
func AddCompany(ct Company) { //Tra ve
	o := orm.NewOrm()

	qs := o.QueryTable("company")
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
func UpdateCompany(domain string, status int) {
	o := orm.NewOrm()

	id, err := o.QueryTable("company").Filter("companyDomain", domain).Update(orm.Params{
		"status": status, //Trang thai master gui 1 - chap nhan; 2 - tu choi
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}

func DeleteCompany(domain string) {
	o := orm.NewOrm()

	_, err := o.QueryTable("company").Filter("companyDomain", domain).Delete()
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
func FindCompanyWithStatus(s int) []Domain_Email {
	o := orm.NewOrm()

	var ct []*Company
	num, err := o.QueryTable("company").Filter("status", s).All(&ct)
	if err == orm.ErrNoRows { // No result
		return nil
	}

	domains := make([]Domain_Email, 0, num)
	for i := 0; i < len(ct); i++ {
		email := FindUserWithIdCongTy((*ct[i]).Id)
		domain := Domain_Email{Domain: (*ct[i]).CompanyDomain, Email: email}
		domains = append(domains, domain)
	}

	return domains
}
