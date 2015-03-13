package main

import (
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/tjsage/todo/controllers"
	"github.com/tjsage/todo/models"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
	}))

	store := sessions.NewCookieStore([]byte("NkUaEUCJZ4T6Yb8aaWrZByzK"))
	m.Use(sessions.Sessions("my_session", store))

	mywhiteList := []string{"/login", "/login/setup"}
	m.Use(SetupLoginHandler(mywhiteList))

	m.Get("/", controllers.TaskIndex)
	m.Get("/tasks/delete/:id", controllers.DeleteTask)
	m.Get("/login", controllers.LoginPage)
	m.Get("/login/setup", controllers.LoginSetup)

	//m.Post("/tasks", binding.Bind(controllers.TaskSearchData{}), controllers.TaskSearch)
	m.Post("/tasks/new", binding.Bind(models.Task{}), controllers.NewTask)
	m.Post("/login", binding.Bind(controllers.Login{}), controllers.LoginPost)
	m.Run()
}
