package models

const (
	Apimap    = "https://geocode-maps.yandex.ru/1.x/?&apikey=dfe6e27c-932d-4e77-9f41-1d30f9cdfcf8&geocode="
	Apiformat = "&format=json"
	Count     = "&results=1"
	Lang      = "&lang=en_RU"
)

type HadeMarker struct {
	Response Response `json:"response"`
}
type Response struct {
	GeoObjectCollection GeoObjectCollection `json:"GeoObjectCollection"`
}

type GeoObjectCollection struct {
	FeatureMember []FeatureMember `json:"featureMember"`
}
type FeatureMember struct {
	GeoObject GeoObject `json:"GeoObject"`
}
type GeoObject struct {
	City    string `json:"name"`
	Country string `json:"description"`
	Point   Point  `json:"Point"`
}
type Point struct {
	Pos string `json:"pos"`
}
type Markers struct {
	Country string
	City    string
	X       float64
	Y       float64
}
