package models

import (
	"github.com/tjsage/todo/services"
)

type Login struct {
	ID       int
	UserName string `sql:"type:varchar(100)"`
	Password string `sql:"type:char(60)"`
}

func GetLogin(id int) *Login {
	var login = &Login{}
	db.Find(login, id)
	return login
}

func FindLogin(user string) *Login {
	var login = &Login{}
	db.Where("user_name = ?", user).Find(login)
	return login
}

func SetupFirstLogin() {
	adminLogin := FindLogin("admin")

	if adminLogin.ID == 0 {
		adminLogin.UserName = "admin"
		adminLogin.Password = services.EncryptPassword("password")
		db.Save(adminLogin)
	}
}
