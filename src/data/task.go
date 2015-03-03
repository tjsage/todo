package data

import (
	"log"
)

type Task struct {
	TaskId      int
	NewId       int
	SomeLand    string
	Name        string `form:"Name" binding:"Required" maxlength:"100"`
	CountryCode int    `form:"countryCode" binding:"Required"`
}

func GetTasks() []Task {
	var tasks []Task
	_, err := dbMap.Select(&tasks, "Select * from Tasks")
	if err != nil {
		log.Println("Error - ", err)
	}

	log.Println("Tasks ", tasks)
	return tasks
}

func SearchTasks(term string) []Task {
	var tasks []Task
	//var adustedSearch string = "%" + term + "%"
	//_, err := dbMap.Select(&tasks, "SELECT * FROM Tasks where Name like ?", adustedSearch)

	_, err := dbMap.Select(&tasks, "SELECT * FROM Tasks Where SOUNDEX(`Name`) = SOUNDEX(?)", term)
	if err != nil {
		log.Println(err)
	}

	return tasks
}

func GetTask(id int) *Task {
	dbMap := GetDbMap()
	task := &Task{}
	dbMap.Get(task, id)

	return task
}

func (task *Task) Save() bool {
	log.Println("Saving task")

	if !task.validate() {
		log.Println("Failed to validate task")
		return false
	}

	if task.TaskId == 0 {
		dbMap.Insert(task)
	} else {
		dbMap.Update(task)
	}

	return true
}

func (task *Task) validate() bool {
	var dataValid = true

	if len(task.Name) > 100 {
		dataValid = false
		log.Println("Task length must be less than 100")
	}

	return dataValid
}
