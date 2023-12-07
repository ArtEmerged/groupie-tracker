package models

type Artists struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relations    map[string][]string `json:"-"`
	Markers      []Markers
	CountMark    int
}

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}


const (
	ArtistsAPI  = "https://groupietrackers.herokuapp.com/api/artists"
	RelationAPI = "https://groupietrackers.herokuapp.com/api/relation"
)