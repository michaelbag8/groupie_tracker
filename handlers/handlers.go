package handlers

import (
	"groupie_tracker/models"
	"html/template"
	"net/http"
)


func MakeHomeHandler(fullArtists []models.FullArtist) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("templates/index.html")
		if err !=nil{
			http.Error(w, "error parsing template file", http.StatusInternalServerError)
			return 
		}

		err = temp.Execute(w, fullArtists)
		if err !=nil{
			http.Error(w, "error executing template", http.StatusInternalServerError)
			return 
		}
		
    }
}
