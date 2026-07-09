package main

import (
	"fmt"
	"log"
	"groupie_tracker/services"
)

func main() {
	urlArtist := "https://groupietrackers.herokuapp.com/api/artists"
	urlLocation := "https://groupietrackers.herokuapp.com/api/locations"
	urlDates := "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation := "https://groupietrackers.herokuapp.com/api/relation"

	res, err := services.FetchData(urlArtist)
	if err != nil {
		log.Fatal(err)
	}

	data, err := services.DecodeArtists(res)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data{
		fmt.Println(v.Name)
	}

}
