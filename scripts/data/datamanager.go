package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Universal getData function to grab data from GetRawData and unmarshal it into a struct of your choice, then return the right type of struct
func getData(url string, data interface{}) {
	rawData := getRawData(url)
	err := json.Unmarshal(rawData, &data)
	if err != nil {
		log.Panic("Problem in getData function when unmarshalling data: ", err)
	}
}

func getRawData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Panic("Problem in getRawData function when getting response: ", err)
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic("Problem in getRawData function when reading response body: ", err)
		return []byte{}
	}
	return responseData
}
