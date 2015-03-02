package data

import (
	"log"
)
type Task struct {
	TaskId int
	Name string `form:"Name"`
}


func GetTasks() []Task {
	var tasks []Task
	_, err := dbMap.Select(&tasks, "Select * from Tasks")
	if (err != nil) {
		log.Println("Error - ", err)
	}

	log.Println("Tasks ", tasks)
	return tasks
}

func GetTask(id int) *Task {
	dbMap := GetDbMap()
	task := &Task {}
	dbMap.Get(task, id)

	return task
}

func (task *Task) Save() {
	log.Println("Save")
	if (task.TaskId == 0) {
		dbMap.Insert(task)
	} else {
		dbMap.Update(task)
	}
}

