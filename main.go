package main

import (
	"fmt"
	"log"
	"groupie_tracker/services"
)

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	res, err := services.FetchArtists(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
