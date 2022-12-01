package data

// Artist type
type rawArtist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func getArtists() []rawArtist {
	artists := []rawArtist{}
	getData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	return artists
}

// Dates type
type rawDates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

func getDates() rawDates {
	dates := rawDates{}
	getData("https://groupietrackers.herokuapp.com/api/dates", &dates)
	return dates
}

// Locations type
type rawLocations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

func getLocations() rawLocations {
	locations := rawLocations{}
	getData("https://groupietrackers.herokuapp.com/api/locations", &locations)
	return locations
}

// Relation type
type rawRelation struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func getRelations() rawRelation {
	relations := rawRelation{}
	getData("https://groupietrackers.herokuapp.com/api/relation", &relations)
	return relations
}
