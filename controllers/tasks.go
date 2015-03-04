package controllers

import (
	"log"
	"strconv"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
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

func DeleteTask(args martini.Params, r render.Render) {
	i, err := strconv.Atoi(args["id"])
	if err != nil {
		log.Println("Error during DeleteTask", err)
		return
	}

	models.DeleteTask(i)
	r.Redirect("/")
}

func NewTask(task models.Task, r render.Render) {
	log.Println("Task: ", task)
	task.Save()
	r.Redirect("/")
}

/*Private helper functions */

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
