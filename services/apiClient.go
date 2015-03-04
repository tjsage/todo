package services

import (
	"io/ioutil"
	"log"
	"net/http"
)

//Makes a GET call to the API Server.
func ApiGet(url string) []byte {
	fullUrl := GetConfig().APIServer + url
	resp, err := http.Get(fullUrl)

	log.Println("Making an API Call to ", fullUrl)
	if err != nil {
		log.Println("Error making api api call to ", url, err)
	} else {
		content, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			return content
		} else {
			log.Println("Error during Reading Response Body ", err)
		}
	}

	return nil
}
