package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver.
)

var (
	wrongUser = 0
	admin     = 1
	members   = 2

	adminPosition    = 0
	deleteUserStatus = 0
	activeUserStatus = 1
	notFound         = -1
	notFoundPosition = -1
)

type Company struct {
	Id            int    `orm:"column(idCompany);null"`
	CompanyDomain string `orm:"column(companyDomain)"`
	Status        int    `orm:"column(status)"`
}

type User struct {
	Id         int    `orm:"column(idUser);null"`
	UserName   string `orm:"column(userName)"`
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
	BeginDate          string `orm:"column(beginDate)"`
	FinishDate         string `orm:"column(finishDate)"`
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
	Category            string `orm:"column(category)"`
	SolutionDescription string `orm:"column(solutionDescription)"`
	IdDev               int    `orm:"column(idDev)"`
	IdTest              int    `orm:"column(idTest)"`
	IdProject           int    `orm:"column(idProject)"`
	FoundDate           string `orm:"column(foundDate)"`
	UpdateDate          string `orm:"column(updateDate)"`
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
