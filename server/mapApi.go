package server

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	apimap    = "https://geocode-maps.yandex.ru/1.x/?&apikey=dfe6e27c-932d-4e77-9f41-1d30f9cdfcf8&geocode="
	apiformat = "&format=json"
	count     = "&results=1"
	lang      = "&lang=en_RU"
)

type hadeMarker struct {
	Response response `json:"response"`
}
type response struct {
	GeoObjectCollection geoObjectCollection `json:"GeoObjectCollection"`
}

type geoObjectCollection struct {
	FeatureMember []featureMember `json:"featureMember"`
}
type featureMember struct {
	GeoObject geoObject `json:"GeoObject"`
}
type geoObject struct {
	City    string `json:"name"`
	Country string `json:"description"`
	Point   point  `json:"Point"`
}
type point struct {
	Pos string `json:"pos"`
}
type markers struct {
	Country string
	City    string
	X       float64
	Y       float64
}

func requestMapApi(group *artists) error {
	marker := make([]markers, len(group.Relations))
	index := 0
	for location := range group.Relations {
		pars := (strings.Split(location, "-"))[0]
		resp, err := http.Get(apimap + pars + count + lang + apiformat)
		if err != nil {
			return err
		}
		var decodeData hadeMarker
		err = json.NewDecoder(resp.Body).Decode(&decodeData)
		if err != nil {
			return err
		}
		geoObgect := decodeData.Response.GeoObjectCollection.FeatureMember[0].GeoObject
		x, y, err := convertCoordinates(geoObgect.Point.Pos)
		if err != nil {
			return err
		}
		marker[index].City = geoObgect.City
		marker[index].Country = geoObgect.Country
		marker[index].X = x
		marker[index].Y = y
		index++
		resp.Body.Close()
	}
	group.Markers = marker
	group.CountMark = len(marker)
	return nil
}
