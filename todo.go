package main

import (
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/tjsage/todo/controllers"
	"github.com/tjsage/todo/models"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
	}))

	m.Get("/", controllers.TaskIndex)
	//m.Post("/tasks", binding.Bind(controllers.TaskSearchData{}), controllers.TaskSearch)
	m.Post("/tasks/new", binding.Bind(models.Task{}), controllers.NewTask)
	m.Run()
}
