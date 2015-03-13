package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/tjsage/todo/models"
	"github.com/tjsage/todo/services"
)

type Login struct {
	LoginId  int
	UserName string `form:"UserName"`
	Password string `form:"Password"`
}

func LoginPage(r render.Render) {
	r.HTML(200, "login", nil, SIMPLE_LAYOUT)
}

func LoginPost(loginInfo Login, r render.Render, session sessions.Session) {
	login := models.FindLogin(loginInfo.UserName)
	if login.ID != 0 {
		if services.ComparePassword(loginInfo.Password, login.Password) {
			session.Set("user_id", string(login.ID))
			r.Redirect("/")
		}
	}

	r.HTML(200, "login", nil, SIMPLE_LAYOUT)
}

func LoginSetup(r render.Render) {
	models.SetupFirstLogin()
	r.Redirect("/")
}
