package server

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func filterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filters/" {
		errPage(w, http.StatusNotFound) // 404
		return
	}

	if r.Method != http.MethodGet {
		errPage(w, http.StatusMethodNotAllowed) // 405
		return
	}

	artistsPage, relations, err := artist()
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}

	confing, err := FilterConfig(artistsPage)
	if err != nil {
		errPage(w, http.StatusBadRequest) // 400
		return
	}

	artistsPage = createDateForSearch(artistsPage, relations)
	list, err := filterFor(r, artistsPage, confing)
	if err != nil {
		errPage(w, http.StatusBadRequest) // 400
		return
	}
	err = tpl.ExecuteTemplate(w, "index.html", list)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
}

func filterFor(r *http.Request, Bands []artists, filterConfig Filters) ([]artists, error) {
	var list []artists

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

type Filters struct {
	MaxCreationDate   int
	MinCreationDate   int
	MaxFirstAlbumDate int
	MinFirstAlbumDate int
	MinMembers        int
	MaxMembers        int
}

func FilterConfig(Bands []artists) (Filters, error) {
	min_creation_date := 2100
	max_creation_date := 0

	min_album_date := 2100
	max_album_date := 0

	min_members := 100
	max_members := 0

	for _, v := range Bands {
		if v.CreationDate < min_creation_date {
			min_creation_date = v.CreationDate
		}

		if v.CreationDate > max_creation_date {
			max_creation_date = v.CreationDate
		}

		year, err := strconv.Atoi(v.FirstAlbum[6:])
		if err != nil {
			log.Fatal("Error on converting date string to date int")
			return Filters{}, err
		}

		if year < min_album_date {
			min_album_date = year
		}

		if year > max_album_date {
			max_album_date = year
		}

		if len(v.Members) < min_members {
			min_members = len(v.Members)
		}

		if len(v.Members) > max_members {
			max_members = len(v.Members)
		}
	}

	return Filters{
		MinCreationDate:   min_creation_date,
		MaxCreationDate:   max_creation_date,
		MinFirstAlbumDate: min_album_date,
		MaxFirstAlbumDate: max_album_date,
		MinMembers:        min_members,
		MaxMembers:        max_members,
	}, nil
}
