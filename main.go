package main

import (
	"GoBug/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// set default database
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost:3306)/mydb?charset=utf8", 30, 30)
}

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login/", &controllers.LoginController{}, "get:Get;post:Login")
	beego.Router("/signup/", &controllers.CongTyController{}, "get:Get;post:Add;put:Update")
	beego.Run()
}
