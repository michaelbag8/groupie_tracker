package main

import (
	"fmt"
	"log"
	"groupie_tracker/services"
	"groupie_tracker/models"
)

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	res, err := services.FetchArtists(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := services.DecodeArtists(res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
