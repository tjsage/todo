package data

type Task struct {
	TaskId int
	Name string
}


func GetTasks() []Task {
	tasks = []Task{}
	tasks = dbMap.Select(tasks, "Select * from Tasks")
}

func GetTask(id int) *Task {
	dbMap := GetDbMap()
	task := &Task {}
	task = dbMap.Get(id, "TaskId")

	return task
}

func (task *Task) Save() {
	
}

