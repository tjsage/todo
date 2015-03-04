package controllers

import (
	"log"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/tjsage/todo/models"
)

type IndexView struct {
	Countries []models.Country
	Tasks     []models.Task
}

type TaskList struct {
	Countries []models.Country
	Tasks     []models.Task
}

type TaskSearchView struct {
	SearchTerm string `form:"SearchTerm"`
}

func TaskIndex(r render.Render) {
	tasks := models.GetTasks()
	countries := models.GetCountries()
	taskList := createTaskList(countries, tasks)

	r.HTML(200, "index", taskList)
}

/*func TaskSearch(taskSearch models.Task, r render.Render) {
	log.Println("Searching for: ", taskSearch.SearchTerm)
	tasks := data.SearchTasks(taskSearch.SearchTerm)
	r.JSON(200, tasks)
}*/

func NewTask(task models.Task, r render.Render) {
	log.Println("Task: ", task)
	task.Save()
	r.Redirect("/")
}

func createTaskList(countries []models.Country, tasks []models.Task) *TaskList {
	var countryMap = make(map[int]models.Country)

	for _, country := range countries {
		countryMap[country.Id] = country
	}

	for _, task := range tasks {
		log.Printf("%+v", task)
		//task.CountryName = countryMap[task.CountryCode].DisplayName
	}

	return &TaskList{
		Countries: countries,
		Tasks:     tasks,
	}
}
