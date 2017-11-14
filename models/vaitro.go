package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func AddVaitro() {
	o := orm.NewOrm()

	var vt Vaitro
	vt.Tenvaitro = "3"

	id, err := o.Insert(&vt)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	if err == nil {
		fmt.Println(id)
		return
	}

}
