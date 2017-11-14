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
}

type Vaitro struct {
	Id        int    `orm:"column(idvaitro);null"`
	Tenvaitro string `orm:"column(tenvaitro)"`
}

type Project struct {
	Id            int    `orm:"column(idProject);null"`
	TenProject    string `orm:"column(tenProject)"`
	MieutaProject string `orm:"column(mieutaProject)"`
}

func init() {
	orm.RegisterModel(new(User), new(Vaitro), new(CongTy), new(Project))
}
