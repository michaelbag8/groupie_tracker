package utils

import (
	"groupie_tracker/models"
)

func BuildRelationsMap(items []models.RelationEntry) map[int]models.RelationEntry {
	result := make(map[int]models.RelationEntry)
	for _, item := range items {
		result[item.ID] = item
	}
	return result
}
func BuildDatesMap(items []models.DateEntry) map[int]models.DateEntry {
	result := make(map[int]models.DateEntry)
	for _, item := range items {
		result[item.ID] = item
	}
	return result
}

func BuildLocationsMap(items []models.LocationEntry) map[int]models.LocationEntry {
	result := make(map[int]models.LocationEntry)
	for _, item := range items {
		result[item.ID] = item
	}
	return result
}