package services

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"groupie_tracker/models"
)


func FetchData(url string) (string, error) {
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

func DecodeRelations(jsonStr string) ([]models.RelationEntry, error){
	var result models.RelationsResponse
	if err :=json.Unmarshal([]byte(jsonStr), &result);err!=nil{
		return nil, fmt.Errorf("invalid json %w", err)
	}
	return result.Index, nil
}

func DecodeLocations(jsonStr string) ([]models.LocationEntry, error){
	var result models.LocationsResponse
	if err :=json.Unmarshal([]byte(jsonStr), &result);err!=nil{
		return nil, fmt.Errorf("invalid json %w", err)
	}

	return result.Index, nil
}

func DecodeDates(jsonStr string) ([]models.DateEntry, error){
	var result models.DatesResponse
	if err :=json.Unmarshal([]byte(jsonStr), &result);err!=nil{
		return nil, fmt.Errorf("invalid json %w", err)
	}
	return result.Index, nil
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

func MergeArtists(artists []models.Artist, locMap map[int]models.LocationEntry, dateMap map[int]models.DateEntry, relMap map[int]models.RelationEntry) []models.FullArtist {
	var result []models.FullArtist
	for _, artist := range artists {
		value, ok := locMap[artist.ID]
		if !ok {
			fmt.Printf("location not found for artist ID %d", artist.ID)
		}

		values, ok := dateMap[artist.ID]
		if !ok {
			fmt.Printf("date not found for artist ID %d", artist.ID)
		}

		val, ok := relMap[artist.ID]
		if !ok {
			fmt.Printf("relation not found for artist ID %d", artist.ID)
		}

		full := models.FullArtist{
			ID:             artist.ID,
			Name:           artist.Name,
			Image:          artist.Image,
			Members:        artist.Members,
			CreationDate:   artist.CreationDate,
			FirstAlbum:     artist.FirstAlbum,
			Locations:      value.Locations,
			ConcertDates:   values.Dates,
			DatesLocations: val.DatesLocations,
		}
		result = append(result, full)

	}

	return result
}
