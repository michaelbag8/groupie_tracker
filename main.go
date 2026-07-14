package main

import (
	"fmt"
	"groupie_tracker/handlers"
	"groupie_tracker/services"
	"groupie_tracker/utils"
	"log"
	"net/http"
)

func main() {
	urlArtist := "https://groupietrackers.herokuapp.com/api/artists"
	urlLocation := "https://groupietrackers.herokuapp.com/api/locations"
	urlDates := "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation := "https://groupietrackers.herokuapp.com/api/relation"

	location, err := services.FetchData(urlLocation)
	if err != nil {
		log.Fatal(err)
	}
	artist, err := services.FetchData(urlArtist)
	if err != nil {
		log.Fatal(err)
	}
	dates, err := services.FetchData(urlDates)
	if err != nil {
		log.Fatal(err)
	}
	relation, err := services.FetchData(urlRelation)
	if err != nil {
		log.Fatal(err)
	}

	art, err := services.DecodeArtists(artist)
	if err != nil {
		log.Fatal(err)
	}

	loc, err := services.DecodeLocations(location)
	if err != nil {
		log.Fatal(err)
	}

	rel, err := services.DecodeRelations(relation)
	if err != nil {
		log.Fatal(err)
	}

	date, err := services.DecodeDates(dates)
	if err != nil {
		log.Fatal(err)
	}

	d := utils.BuildDatesMap(date)
	l := utils.BuildLocationsMap(loc)
	r := utils.BuildRelationsMap(rel)
	final := services.MergeArtists(art, l, d, r)

	fmt.Println("server running at http://localhost:8080/")

	http.HandleFunc("/", handlers.MakeHomeHandler(final))
	http.HandleFunc("/artist/", handlers.MakeArtistHandler(final))
	log.Fatal((http.ListenAndServe(":8080", nil)))

}
