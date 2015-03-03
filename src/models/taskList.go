package models

import (
	"data"
	"kyClient"
	"log"
)

type TaskList struct {
	Countries []kyClient.Country
	Tasks     []data.Task
}

func CreateTaskList(countries []kyClient.Country, tasks []data.Task) *TaskList {
	var countryMap = make(map[int]kyClient.Country)

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
