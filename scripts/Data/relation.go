package get

type Relations struct {
	Id             uint64                 `json:"id"`
	DatesLocations map[string]interface{} `json:"datesLocations"`
}
