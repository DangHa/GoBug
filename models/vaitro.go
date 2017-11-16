package models

import (
	"github.com/astaxie/beego/orm"
)

func CheckAdmin(id int) bool {
	o := orm.NewOrm()

	var vt = Vaitro{}
	err := o.QueryTable("vai_tro").Filter("Id", id).One(&vt)

	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}
