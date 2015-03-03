package kyClient

import (
	"encoding/json"
	"log"
)

type Country struct {
	Id          int
	DisplayName string `json:"display"`
}

type countryReturnData struct {
	Countries []Country
}

func GetCountries() []Country {
	payload := Get("/list/countries")
	var countryReturnData = &countryReturnData{}

	err := json.Unmarshal(payload, countryReturnData)
	if err != nil {
		log.Println("Error unmarshaling countries", err)
	}

	log.Println("Country Return Data", countryReturnData)
	return countryReturnData.Countries
}
