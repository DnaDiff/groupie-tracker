package get

type Locations struct {
	Id        int    `json:"id"`
	Location  string `json:"location"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	ArtistId  int    `json:"artistId"`
}

func getLocations() {
	// data := GetData("https://groupietrackers.herokuapp.com/api/locations")
	// locations := []Location{}

}
