package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func fetchArtists(url string) (string, error) {
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

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	res, err := fetchArtists(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
