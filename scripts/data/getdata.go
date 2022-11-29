package get

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Panic("Problem in GetData function when getting response: ", err)
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic("Problem in GetData function when reading response body: ", err)
		return []byte{}
	}
	return responseData
}
