package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver.
)

type Company struct {
	Id            int    `orm:"column(idCompany);null"`
	CompanyDomain string `orm:"column(companyDomain)"`
	Status        int    `orm:"column(status)"`
}

type User struct {
	Id         int    `orm:"column(idUser);null"`
	Email      string `orm:"column(email)"`
	Password   string `orm:"column(password)"`
	IdCompany  int    `orm:"column(idCompany)"`
	IdPosition int    `orm:"column(idPosition)"`
	Status     int    `orm:"column(status)"`
}

type Position struct {
	Id           int    `orm:"column(idPosition);null"`
	PositionName string `orm:"column(positionName)"`
}

type Project struct {
	Id                 int    `orm:"column(idProject);null"`
	ProjectName        string `orm:"column(projectName)"`
	ProjectDescription string `orm:"column(projectDescription)"`
}

type UserProject struct {
	Id        int `orm:"column(id);null"`
	IdUser    int `orm:"column(idUser)"`
	IdProject int `orm:"column(idProject)"`
}

type Bug struct {
	Id                  int    `orm:"column(idBug);null"`
	BugName             string `orm:"column(bugName)"`
	BugDescription      string `orm:"column(bugDescription)"`
	SolutionDescription string `orm:"column(solutionDescription)"`
	IdUser              int    `orm:"column(idUser)"`
	IdProject           int    `orm:"column(idProject)"`
}

type Master struct {
	Id       int    `orm:"column(idmaster);null"`
	Email    string `orm:"column(email)"`
	Password string `orm:"column(password)"`
}

func init() {
	orm.RegisterModel(
		new(User),
		new(Position),
		new(Company),
		new(Project),
		new(UserProject),
		new(Master),
		new(Bug))
}
