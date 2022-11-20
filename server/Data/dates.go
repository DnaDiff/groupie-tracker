package get

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

func GetDates() Dates {
	
	data := GetData("https://groupietrackers.herokuapp.com/api/dates")
	dates := Dates{}
	err := json.Unmarshal(data, &dates)
	if err != nil {
		log.Panic("Problem in GetDates function when unmarshalling data: ", err)
		return Dates{}
	}
	
	fmt.Println(dates.Index[0].ID)
	fmt.Println(dates.Index[0].Dates)

	return dates
}
