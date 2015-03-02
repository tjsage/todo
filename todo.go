
package main

import (
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/binding"
	"data"
	"log"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options {
		Directory: "templates",
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		tasks := data.GetTasks()
		r.HTML(200, "index", tasks)
	})

	m.Post("/Task/New", binding.Bind(data.Task{}), func(task data.Task, r render.Render) {
		log.Println("Task: ", task)
		task.Save()
		r.Redirect("/")
	})

	m.Run()
}