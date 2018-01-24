package routers

import (
	"bugmanage/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// Log in and sign up danh cho user va admin
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signup/", &controllers.CongTyController{}, "get:Get;post:Add")
	beego.Router("/signup/again/", &controllers.CongTyAgainController{}, "get:Get")
	beego.Router("/login/", &controllers.LoginController{}, "get:Get;post:Login")
	beego.Router("/login/again/", &controllers.LoginAgainController{}, "get:Get")

	//Danh cho User
	beego.Router("/userprojectjson/", &controllers.UserProjectJson{}, "get:Get;post:Post")
	beego.Router("/userbugjson/", &controllers.UserBugJson{}, "get:Get;post:Post;put:Update;delete:Delete")
	beego.Router("/user/findbug/", &controllers.UserFindBug{}, "get:Get;post:Post")
	beego.Router("/user/profile/", &controllers.UserProfile{}, "get:Get;post:Post")

	// Danh cho Admin
	beego.Router("/loginAdmin/", &controllers.LoginAdminController{}, "get:Get;post:LogOut")
	beego.Router("/adminprojectjson/", &controllers.AdminProjectJsonController{}, "get:Get;post:Post;put:Update;delete:Delete")
	beego.Router("/adminmemberprojectjson/", &controllers.AdminMemberProjectJsonController{}, "get:Get;post:Post;put:Update;delete:Delete")
	beego.Router("/adminmember/", &controllers.AdminMemberController{}, "get:Get;post:Post")
	beego.Router("/adminmemberjson/", &controllers.AdminMemberJsonControllers{}, "get:Get;post:Post;delete:Delete")
	beego.Router("/admin/addmember/", &controllers.AdminAddMemberJsonControllers{}, "post:Post")
	beego.Router("/adminstat/", &controllers.AdminStatControllers{})
	beego.Router("/adminstatjson/", &controllers.AdminStatJsonControllers{})
	beego.Router("/admin/bugstat/json/", &controllers.AdminBugStatJsonControllers{})
	beego.Router("/admin/devstat/json/", &controllers.AdminDevStatJsonControllers{})

	// Danh cho master
	beego.Router("/master/login/", &controllers.MasterLoginController{}, "get:Get;post:Login")
	beego.Router("/master/", &controllers.MasterController{}, "get:Get;put:Update;delete:Delete")
	beego.Router("/master/getjson/", &controllers.MasterJsonController{}, "get:Get")
	beego.Router("/master/active/", &controllers.MasterJsonActiveController{}, "get:Get")
	beego.Router("/master/getjsoncongty/", &controllers.MasterJsonCongTyController{}, "get:Get")
}
