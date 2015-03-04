package models

import (
	"encoding/json"
	"log"

	"github.com/tjsage/todo/services"
)

type Country struct {
	Id          int
	DisplayName string `json:"display"`
}

type countryReturnData struct {
	Countries []Country
}

func GetCountries() []Country {
	payload := services.ApiGet("/list/countries")
	var countryReturnData = &countryReturnData{}

	err := json.Unmarshal(payload, countryReturnData)
	if err != nil {
		log.Println("Error unmarshaling countries", err)
	}

	log.Println("Country Return Data", countryReturnData)
	return countryReturnData.Countries
}
