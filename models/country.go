package models

import (
	"encoding/json"
	"log"

	"github.com/kyani-inc/road-rally/service/cacher"
	"github.com/tjsage/todo/services"
)

type Country struct {
	Id          int
	DisplayName string `json:"display"`
}

type countryReturnData struct {
	Countries []Country
}

func init() {
	cacher.InitCache(cacher.Options{})
}

func GetCountries() []Country {
	var countryReturnData = &countryReturnData{}
	var payload []byte
	payload = cacher.Get("countries")

	if len(payload) == 0 {
		payload = services.ApiGet("/list/countries")
		cacher.Set("countries", payload)
	}

	err := json.Unmarshal(payload, countryReturnData)
	if err != nil {
		log.Println("Error unmarshaling countries", err)
	}

	log.Println("Country Return Data", countryReturnData)

	return countryReturnData.Countries
}
