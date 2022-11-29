package get

type Everyone struct {
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

func Epicmerge() []Everyone {
	everyone := []Everyone{}
	artists := GetArtists()
	dates := GetDates()
	locations := GetLocations()
	for i := 0; i < len(artists); i++ {
		everyone = append(everyone, Everyone{
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
	return everyone
}
