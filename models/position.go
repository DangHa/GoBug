package models

import (
	"github.com/astaxie/beego/orm"
)

func CheckAdmin(id int) bool {
	o := orm.NewOrm()

	var vt = Position{}
	err := o.QueryTable("position").Filter("Id", id).One(&vt)

	if err == orm.ErrMultiRows { // Have multiple records
		return false
	}
	if err == orm.ErrNoRows { // No result
		return false
	}

	return true
}

func FindPosition(idPosition int) string {
	o := orm.NewOrm()

	var vt = Position{}
	err := o.QueryTable("position").Filter("Id", idPosition).One(&vt)

	if err == orm.ErrMultiRows { // Have multiple records
		return ""
	}
	if err == orm.ErrNoRows { // No result
		return ""
	}

	return vt.PositionName
}

func FindPositionWithName(positionName string) int {
	o := orm.NewOrm()

	var vt = Position{}
	err := o.QueryTable("position").Filter("positionName", positionName).One(&vt)

	if err == orm.ErrMultiRows { // Have multiple records
		return 0
	}
	if err == orm.ErrNoRows { // No result
		return 0
	}

	return vt.Id
}
