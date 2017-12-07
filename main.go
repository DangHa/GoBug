package main

import (
	_ "bugmanage/routers"

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

	beego.Run()
}
