package services

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"groupie_tracker/models"
)

func FetchArtists(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error getting data from API %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "",fmt.Errorf("error getting data from API %d", resp.StatusCode)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body %w", err)
	}

	return string(content), nil
}

func DecodeArtists(jsonStr string) ([]models.Artist, error){
	var result []models.Artist
	if err :=json.Unmarshal([]byte(jsonStr), &result);err!=nil{
		return nil, fmt.Errorf("invalid json %w", err)
	}
	return result, nil
}

/*
func DecodeArtists(jsonStr string) ([]models.Artist, error) {
	var result []models.Artist

	decoder := json.NewDecoder(strings.NewReader(jsonStr))
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("invalid json: %w", err)
	}

	return result, nil
}
*/	