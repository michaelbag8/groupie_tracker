package handlers

import (
	"groupie_tracker/models"
	"html/template"
	"log"
	"net/http"
)

func RenderError(w http.ResponseWriter, message string, statusCode int){
	temp, err := template.ParseFiles("templates/error.html")
	if err!=nil{
		http.Error(w, "error parsing html file", http.StatusInternalServerError)
		return
	}
	errorpage := models.Error{
		Message: message,
		Code: statusCode,
	}
	w.WriteHeader(statusCode)
	err = temp.Execute(w, errorpage)
	if err!=nil{
		log.Println("error executing file")
	}
}
