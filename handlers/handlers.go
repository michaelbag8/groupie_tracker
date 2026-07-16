package handlers

import (
	"groupie_tracker/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func MakeHomeHandler(fullArtists []models.FullArtist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("templates/index.html")
		if err != nil {
			RenderError(w, "error parsing template file", http.StatusInternalServerError)
			return
		}
		err = temp.Execute(w, fullArtists)
		if err != nil {
			log.Println("error executing template")
		}
	}
}

func MakeArtistHandler(fullArtists []models.FullArtist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path
		idstring := strings.Split(id, "/")
		convID, err := strconv.Atoi(idstring[2])
		if err != nil {
			RenderError(w, "conversion failed", http.StatusBadRequest)
			return
		}
		found := false
		var foundArtist models.FullArtist
		for _, artist := range fullArtists {
			if artist.ID == convID {
				found = true
				foundArtist = artist
				break
			}

		}
		if !found {
			RenderError(w, "Artist does not exist", http.StatusNotFound)
			return
		}
		temp, err := template.ParseFiles("templates/artist.html")
		if err != nil {
			RenderError(w, "error parsing template file", http.StatusInternalServerError)
			return
		}
		err = temp.Execute(w, foundArtist)
		if err != nil {
			log.Println("error executing template")
		}
	}

}
