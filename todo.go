package main

import (
	"data"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"kyClient"
	"log"
	"models"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
	}))

	m.Get("/", func(r render.Render) {
		tasks := data.GetTasks()
		countries := kyClient.GetCountries()
		taskList := models.CreateTaskList(countries, tasks)

		r.HTML(200, "index", taskList)
	})

	m.Get("/countries", func(r render.Render) {
		log.Println("LIST Countries")
		myCountries := kyClient.GetCountries()
		r.JSON(200, myCountries)
	})

	m.Post("/tasks", binding.Bind(models.TaskSearch{}), func(taskSearch models.TaskSearch, r render.Render) {
		log.Println("Searching for: ", taskSearch.SearchTerm)
		tasks := data.SearchTasks(taskSearch.SearchTerm)
		r.JSON(200, tasks)
	})

	m.Post("/Task/New", binding.Bind(data.Task{}), func(task data.Task, r render.Render) {
		log.Println("Task: ", task)
		task.Save()
		r.Redirect("/")
	})

	m.Run()
}
