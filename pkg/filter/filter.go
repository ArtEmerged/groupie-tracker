package filter

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/models"
)

func FilterFor(r *http.Request, Bands []models.Artists, filterConfig models.Filters) ([]models.Artists, error) {
	var list []models.Artists

	r.ParseForm()

	MinCreationDate, err := strconv.Atoi(r.FormValue("MinCreationDate"))
	if err != nil {
		log.Print("Error on converting string to int:", err)
		return nil, err
	}
	MaxCreationDate, err := strconv.Atoi(r.FormValue("MaxCreationDate"))
	if err != nil {
		log.Print("Error on converting string to int:", err)
		return nil, err
	}
	MinFirstAlbumDate, err := strconv.Atoi(r.FormValue("MinFirstAlbumDate"))
	if err != nil {
		log.Print("Error on converting string to int:", err)
		return nil, err
	}
	MaxFirstAlbumDate, err := strconv.Atoi(r.FormValue("MaxFirstAlbumDate"))
	if err != nil {
		log.Print("Error on converting string to int:", err)
		return nil, err
	}

	Members := []int{}
	for i := 0; i < filterConfig.MaxMembers; i++ {
		if r.FormValue("Member"+strconv.Itoa(i+1)) == "on" {
			Members = append(Members, i+1)
		}
	}
	Locations := []string{}
	Locations = append(Locations, r.Form["locations"]...)

	for _, v := range Bands {
		find := false
		if v.CreationDate >= MinCreationDate && v.CreationDate <= MaxCreationDate {
			FirstAlbum, err := strconv.Atoi(v.FirstAlbum[6:])
			if err != nil {
				log.Print("Error on converting date string to date int")
				return nil, err
			}

			if FirstAlbum >= MinFirstAlbumDate && FirstAlbum <= MaxFirstAlbumDate {
				if len(Members) == 0 {
					if len(Locations) == 0 {
						find = true
					} else {
						count := 0

						for _, loc := range Locations {
							for key := range v.Relations {
								loc = strings.ToLower(strings.ReplaceAll(loc, ",", ""))
								key = strings.ReplaceAll(key, "_", " ")
								key = strings.ReplaceAll(key, "-", " ")
								if strings.Contains(strings.ToLower(key), strings.ToLower(loc)) {
									count++
									break
								}
							}
						}

						if count == len(Locations) {
							find = true
						}
					}
				} else {
					if len(Locations) == 0 {
						for _, num := range Members {
							if num == len(v.Members) {
								find = true
							}
						}
					} else {
						for _, num := range Members {
							if num == len(v.Members) {
								count := 0

								for _, loc := range Locations {
									for key := range v.Relations {
										loc = strings.ToLower(strings.ReplaceAll(loc, ",", ""))
										key = strings.ReplaceAll(key, "_", " ")
										key = strings.ReplaceAll(key, "-", " ")
										if strings.Contains(strings.ToLower(key), strings.ToLower(loc)) {
											count++
											break
										}
									}
								}

								if count == len(Locations) {
									find = true
								}
							}
						}
					}
				}
			}
		}
		if find {
			list = append(list, v)
		}
	}

	return list, nil
}
