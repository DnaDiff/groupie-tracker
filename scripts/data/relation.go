package data

import (
	"encoding/json"
	"log"
	"strings"
)

type Relation struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// get the relations from the API
func GetRelations() Relation {
	data := GetData("https://groupietrackers.herokuapp.com/api/relations")
	relations := Relation{}
	// handle error if data is nil
	err := json.Unmarshal(data, &relations)
	if err != nil {
		log.Panic("Problem in GetRelations function when unmarshalling data: ", err)
		return Relation{}
	}
	test, _ := json.Marshal(relations.Index)
	test = []byte(strings.Title(strings.ReplaceAll(string(test), "_", " ")))
	json.Unmarshal(test, &relations.Index)
	return relations
}
