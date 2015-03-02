
package main

import (
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
	//"github.com/codegangsta/martini-contrib/binding"
	"data"
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

	m.Run()
}