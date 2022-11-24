package get

import (
	"encoding/json"
	"log"
)

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

func GetLocations() Locations {

	data := GetData("https://groupietrackers.herokuapp.com/api/locations")
	locations := Locations{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		log.Panic("Problem in GetLocations function when unmarshalling data: ", err)
		return Locations{}
	}
	return locations
}
