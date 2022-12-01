package data

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Relations    string
}

func Compile() []Artist {
	allArtists := []Artist{}
	rawArtists := getArtists()
	dates := getDates()
	locations := getLocations()
	for i := 0; i < len(rawArtists); i++ {
		allArtists = append(allArtists, Artist{
			Id:           rawArtists[i].Id,
			Image:        rawArtists[i].Image,
			Name:         rawArtists[i].Name,
			Members:      rawArtists[i].Members,
			CreationDate: rawArtists[i].CreationDate,
			FirstAlbum:   rawArtists[i].FirstAlbum,
			Locations:    locations.Index[i].Locations,
			ConcertDates: dates.Index[i].Dates,
			Relations:    rawArtists[i].Relations,
		})
	}
	return allArtists
}
