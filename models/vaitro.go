package models

import (
	"github.com/astaxie/beego/orm"
)

func CheckAdmin(id int) bool {
	o := orm.NewOrm()

	var vt = VaiTro{}
	err := o.QueryTable("vai_tro").Filter("Id", id).One(&vt)

	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}

func FindVaitro(id int) string {
	o := orm.NewOrm()

	var vt = VaiTro{}
	err := o.QueryTable("vai_tro").Filter("Id", id).One(&vt)

	if err == orm.ErrMultiRows { // Have multiple records
		return ""
	}
	if err == orm.ErrNoRows { // No result
		return ""
	}

	return vt.Tenvaitro
}

func FindVaitroWithName(tenvaitro string) int {
	o := orm.NewOrm()

	var vt = VaiTro{}
	err := o.QueryTable("vai_tro").Filter("tenvaitro", tenvaitro).One(&vt)

	if err == orm.ErrMultiRows { // Have multiple records
		return 0
	}
	if err == orm.ErrNoRows { // No result
		return 0
	}

	return vt.Id
}
