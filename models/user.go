package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//Kiem tra mat khau cho login
func CheckUser(email, password string) int { // Tra ve  0 - Sai user; 1 - Admin Dang nhap thanh cong; 2 - nhan vien dang nhap thanh cong
	o := orm.NewOrm()

	var user = User{}
	err := o.QueryTable("user").Filter("email", email).Filter("password", password).One(&user)

	if err == orm.ErrMultiRows { // Have multiple records
		return wrongUser
	}
	if err == orm.ErrNoRows { // No result
		return wrongUser
	}

	if user.Status == deleteUserStatus { //Tai khoan nay da bi xoa hoac chua duoc hoat dong
		return wrongUser
	}

	if user.IdPosition == adminPosition { // Admin
		return admin
	}
	return members //Nhan Vien
}

//Update (dung cho Admin thay doi trang thai khi duoc chap nhan tao cong ty)
func UpdateUser(email string) {
	o := orm.NewOrm()

	id, err := o.QueryTable("user").Filter("email", email).Update(orm.Params{
		"status": activeUserStatus, // 0 - bi khoa, 1 - hoat dong
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}

//Update information
func UpdateInformation(user User) {
	o := orm.NewOrm()

	id, err := o.QueryTable("user").Filter("idUser", user.Id).Update(orm.Params{
		"userName": user.UserName,
		"password": user.Password,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}

func DeleteUserByUpdateStatus(email string) {
	o := orm.NewOrm()

	id, err := o.QueryTable("user").Filter("email", email).Update(orm.Params{
		"status": deleteUserStatus, // 0 - bi khoa, 1 - hoat dong
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successful update!,", id)
}

// Them User
func AddUser(u User) {
	o := orm.NewOrm()

	qs := o.QueryTable("user")
	i, _ := qs.PrepareInsert()
	id, err := i.Insert(&u)
	if err != nil {
		fmt.Println(err)
		return
	}

	i.Close()

	fmt.Println("Successful add!,", id)
}

func DeleteUser(email string) {
	o := orm.NewOrm()

	_, err := o.QueryTable("user").Filter("Email", email).Delete()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done!")
}

func FindUserWithIdCongTy(idCompany int) string {
	o := orm.NewOrm()

	var u = User{}
	err := o.QueryTable("user").Filter("idCompany", idCompany).One(&u)

	if err == orm.ErrMultiRows { // Have multiple records
		return ""
	}
	if err == orm.ErrNoRows { // No result
		return ""
	}

	return u.Email
}

// Admin Member
func FindCongTyByIdUser(idUser int) int {
	o := orm.NewOrm()

	var u = User{}
	err := o.QueryTable("user").Filter("idUser", idUser).One(&u)

	if err == orm.ErrMultiRows { // Have multiple records
		return notFound
	}
	if err == orm.ErrNoRows { // No result
		return notFound
	}

	return u.IdCompany
}

func FindMemberOfCongTy(idCompany int) []User {
	o := orm.NewOrm()

	var u []User
	_, err := o.QueryTable("user").Filter("idCompany", idCompany).All(&u)

	if err == orm.ErrNoRows { // No result
		return nil
	}

	// Loai bo nhung user co status = 0 - tuc la da bi xoa
	var u2 []User
	for i := 0; i < len(u); i++ {
		if u[i].Status != deleteUserStatus {
			u2 = append(u2, u[i])
		}
	}

	return u2
}

func FindIdUserWithEmail(email string) int {
	o := orm.NewOrm()

	var u = User{}
	err := o.QueryTable("user").Filter("email", email).One(&u)

	if err == orm.ErrNoRows { // No result
		return notFound
	}

	return u.Id
}

func FindUserWithIdUser(iduser int) User {
	o := orm.NewOrm()

	var u = User{}
	err := o.QueryTable("user").Filter("idUser", iduser).One(&u)

	if err == orm.ErrNoRows { // No result
		return u
	}

	return u
}

type RankDevJson struct {
	UserName string `orm:"column(userName)"`
	Project  int    `orm:"column(Project)"`
	Solution int    `orm:"column(Solution)"`
}

func RankDev(idAdmin int) []RankDevJson {
	o := orm.NewOrm()

	var bs []RankDevJson
	num, err := o.Raw("select userName, (select count(idProject) from user_project where idUser = x.idUser) as Project, (select count(idDev) from bug where idDev = x.idUser and solutionDescription not like '') as Solution from user x where idPosition = 1 and idCompany like (select idCompany from user where idUser = ?)", idAdmin).QueryRows(&bs)
	if err != nil {
		fmt.Println("user nums: ", num)
	}

	fmt.Println("rank: ", bs)

	return bs
}
