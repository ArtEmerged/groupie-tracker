package server

import (
	"fmt"
	"log"
	"net/http"
)

const apimap = "https://geocode-maps.yandex.ru/1.x/?&apikey=dfe6e27c-932d-4e77-9f41-1d30f9cdfcf8&geocode="
const apiformat = "&format=json"
const count = "&results=1"
const lang = "&lang=en_RU"

func MapApi() {
	resp, err := http.Get(apimap + "amsterdam-netherlands" + count + lang + apiformat)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	buff := make([]byte, 10000)
	n, err := resp.Body.Read(buff)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(buff[:n]))
}
