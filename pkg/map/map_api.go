package mapapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"groupie-tracker/internal/models"
)

func RequestMapApi(group *models.Artists) error {
	marker := make([]models.Markers, len(group.Relations))
	index := 0
	for location := range group.Relations {
		pars := (strings.Split(location, "-"))[0]
		resp, err := http.Get(models.Apimap + pars + models.Count + models.Lang + models.Apiformat)
		if err != nil {
			return err
		}
		var decodeData models.HadeMarker
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
