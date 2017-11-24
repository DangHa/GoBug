package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver.
)

type CongTy struct {
	Id            int    `orm:"column(idCongTy);null"`
	TenmienCongTy string `orm:"column(tenmienCongTy)"`
	Status        int    `orm:"column(status)"`
}

type User struct {
	Id       int    `orm:"column(idUser);null"`
	Email    string `orm:"column(email)"`
	Password string `orm:"column(password)"`
	IdCongTy int    `orm:"column(idCongTy)"`
	Idvaitro int    `orm:"column(idvaitro)"`
	Status   int    `orm:"column(status)"`
}

type VaiTro struct {
	Id        int    `orm:"column(idvaitro);null"`
	Tenvaitro string `orm:"column(tenvaitro)"`
}

type Project struct {
	Id            int    `orm:"column(idProject);null"`
	TenProject    string `orm:"column(tenProject)"`
	MieutaProject string `orm:"column(mieutaProject)"`
}

type User_project struct {
	Id        int `orm:"column(id);null"`
	IdUser    int `orm:"column(idUser)"`
	IdProject int `orm:"column(idProject)"`
}

type Bug struct {
	Id             int    `orm:"column(idBug);null"`
	TenBug         string `orm:"column(tenBug)"`
	MieutaBug      string `orm:"column(mieutaBug)"`
	MieutaSolution string `orm:"column(mieutaSolution)"`
	IdUser         int    `orm:"column(idUser)"`
	IdProject      int    `orm:"column(idProject)"`
}

type Master struct {
	Id       int    `orm:"column(idmaster);null"`
	Email    string `orm:"column(email)"`
	Password string `orm:"column(password)"`
}

func init() {
	orm.RegisterModel(
		new(User),
		new(VaiTro),
		new(CongTy),
		new(Project),
		new(User_project),
		new(Master),
		new(Bug))
}
