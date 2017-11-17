package main

import (
	"bugmanage/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// set default database
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost:3306)/mydb?charset=utf8", 30, 30)
}

func main() {
	// Danh cho master
	beego.Router("/masterlogin/", &controllers.MasterLoginController{}, "get:Get;post:Login")
	beego.Router("/master/", &controllers.MasterController{}, "get:Get;put:Update;delete:Delete")
	beego.Router("/mastergetjson/", &controllers.MasterJsonController{}, "get:Get")

	// Danh cho user va admin
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login/", &controllers.LoginController{}, "get:Get;post:Login")
	beego.Router("/signup/", &controllers.CongTyController{}, "get:Get;post:Add")

	beego.Router("/loginAdmin/", &controllers.LoginAdminController{})
	beego.Run()
}
