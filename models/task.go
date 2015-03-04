package models

import (
	"log"
)

type Task struct {
	ID      int
	Country string `sql:"type:varchar(100);" form:"Country"`
	Name    string `form:"Name" binding:"Required" maxlength:"100" sql:"type:varchar(100)"`
}

func GetTasks() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func SearchTasks(term string) []Task {
	var tasks []Task
	//var adustedSearch string = "%" + term + "%"
	//_, err := dbMap.Select(&tasks, "SELECT * FROM Tasks where Name like ?", adustedSearch)

	db.Where("SOUNDEX(`Name`) = Soundex(?)", term).Find("&tasks")
	return tasks
}

func GetTask(id int) *Task {
	task := &Task{}
	db.Find(task, id)

	return task
}

func DeleteTask(id int) {
	db.Delete(&Task{ID: id})
}

func (task *Task) Save() bool {
	log.Println("Saving task")

	if !task.validate() {
		log.Println("Failed to validate task")
		return false
	}

	if task.ID == 0 {
		db.Create(task)
	} else {
		db.Save(task)
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
