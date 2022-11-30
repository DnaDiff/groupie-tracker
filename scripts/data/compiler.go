package data

type Artists struct {
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

func Compile() []Artists {
	rawArtists := []Artists{}
	artists := GetArtists()
	dates := GetDates()
	locations := GetLocations()
	for i := 0; i < len(artists); i++ {
		rawArtists = append(rawArtists, Artists{
			Id:           artists[i].Id,
			Image:        artists[i].Image,
			Name:         artists[i].Name,
			Members:      artists[i].Members,
			CreationDate: artists[i].CreationDate,
			FirstAlbum:   artists[i].FirstAlbum,
			Locations:    locations.Index[i].Locations,
			ConcertDates: dates.Index[i].Dates,
			Relations:    artists[i].Relations,
		})
	}
	return rawArtists
}
