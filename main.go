package main

import (
	"bugmanage/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
)

func init() {
	// set default database
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost:3306)/mydb?charset=utf8", 30, 30)
}

func main() {

	// SessionID
	sessionconf := &session.ManagerConfig{
		CookieName: "bugmanageID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	// Log in and sign up danh cho user va admin
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signup/", &controllers.CongTyController{}, "get:Get;post:Add")
	beego.Router("/signupagain/", &controllers.CongTyAgainController{}, "get:Get")
	beego.Router("/login/", &controllers.LoginController{}, "get:Get;post:Login")
	beego.Router("/loginagain/", &controllers.LoginAgainController{}, "get:Get")

	//Danh cho User
	beego.Router("/userprojectjson/", &controllers.UserProjectJson{}, "get:Get;post:Post")
	beego.Router("/userbugjson/", &controllers.UserBugJson{}, "get:Get;post:Post;put:Update;delete:Delete")

	// Danh cho Admin
	beego.Router("/loginAdmin/", &controllers.LoginAdminController{})
	beego.Router("/adminprojectjson/", &controllers.AdminProjectJsonController{}, "get:Get;post:Post;put:Update;delete:Delete")
	beego.Router("/adminmemberprojectjson/", &controllers.AdminMemberProjectJsonController{}, "get:Get;post:Post;put:Update;delete:Delete")
	beego.Router("/adminmember/", &controllers.AdminMemberController{}, "get:Get;post:Post")
	beego.Router("/adminmemberjson/", &controllers.AdminMemberJsonControllers{}, "get:Get;post:Post;delete:Delete")
	beego.Router("/adminstat/", &controllers.AdminStatControllers{})
	beego.Router("/adminstatjson/", &controllers.AdminStatJsonControllers{})

	// Danh cho master
	beego.Router("/masterlogin/", &controllers.MasterLoginController{}, "get:Get;post:Login")
	beego.Router("/master/", &controllers.MasterController{}, "get:Get;put:Update;delete:Delete")
	beego.Router("/mastergetjson/", &controllers.MasterJsonController{}, "get:Get")
	beego.Router("/masteractive/", &controllers.MasterJsonActiveController{}, "get:Get")
	beego.Router("/mastergetjsoncongty/", &controllers.MasterJsonCongTyController{}, "get:Get")
	beego.Run()
}
