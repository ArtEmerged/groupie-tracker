package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func MapApiTest() {
	resp, err1 := http.Get(apimap + "london" + count + lang + apiformat)
	if err1 != nil {
		log.Println("1", err1)
	}
	var data hadeMarker
	buff := make([]byte, 10000)
	n, err := resp.Body.Read(buff)
	if err != nil {
		log.Println(err)
	}
	err2 := json.Unmarshal(buff[:n], &data)
	if err2 != nil {
		log.Println("2", err2)
	}
	if len(data.Response.GeoObjectCollection.FeatureMember) == 0 {
		fmt.Println("Marker:")
		os.Exit(0)
	}
	geoObgect := data.Response.GeoObjectCollection.FeatureMember[0].GeoObject
	fmt.Println(string(buff[:n]))
	fmt.Println("Marker:", geoObgect)
	resp.Body.Close()
}

func requestMapApi(group *artists) error {
	marker := make([]markers, len(group.Relations))
	index := 0
	for location := range group.Relations {
		pars := (strings.Split(location, "-"))[0]
		resp, err1 := http.Get(apimap + pars + count + lang + apiformat)
		if err1 != nil {
			return err1
		}
		var decodeData hadeMarker
		err2 := json.NewDecoder(resp.Body).Decode(&decodeData)
		if err2 != nil {
			return err2
		}
		geoObgect := decodeData.Response.GeoObjectCollection.FeatureMember[0].GeoObject
		x, y, err3 := convertCoordinates(geoObgect.Point.Pos)
		if err3 != nil {
			return err3
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
